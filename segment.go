// Copyright 2019 Tetrate Labs
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package go2sky

import (
	"sync/atomic"

	"github.com/marlonfan/go2sky/internal/idgen"
	"github.com/marlonfan/go2sky/internal/tool"
	"github.com/marlonfan/go2sky/propagation"
	"github.com/marlonfan/go2sky/reporter/grpc/common"
	v2 "github.com/marlonfan/go2sky/reporter/grpc/language-agent-v2"
)

func newSegmentSpan(defaultSpan *defaultSpan, parentSpan segmentSpan) (s segmentSpan) {
	ssi := &segmentSpanImpl{
		defaultSpan: *defaultSpan,
	}
	ssi.createSegmentContext(parentSpan)
	if parentSpan == nil || !parentSpan.segmentRegister() {
		rs := newSegmentRoot(ssi)
		rs.createRootSegmentContext(parentSpan)
		s = rs
	} else {
		s = ssi
	}
	return
}

// SegmentContext is the context in a segment
type SegmentContext struct {
	TraceID         []int64
	SegmentID       []int64
	SpanID          int32
	ParentSpanID    int32
	ParentSegmentID []int64
	collect         chan<- ReportedSpan
	refNum          *int32
	spanIDGenerator *int32
}

// ReportedSpan is accessed by Reporter to load reported data
type ReportedSpan interface {
	Context() *SegmentContext
	Refs() []*propagation.SpanContext
	StartTime() int64
	EndTime() int64
	OperationName() string
	Peer() string
	SpanType() common.SpanType
	SpanLayer() common.SpanLayer
	IsError() bool
	Tags() []*common.KeyStringValuePair
	Logs() []*v2.Log
	ComponentID() int32
}

type segmentSpan interface {
	Span
	context() SegmentContext
	segmentRegister() bool
}

type segmentSpanImpl struct {
	defaultSpan
	SegmentContext
}

// For Span

func (s *segmentSpanImpl) End() {
	s.defaultSpan.End()
	go func() {
		s.Context().collect <- s
	}()
}

// For Reported Span

func (s *segmentSpanImpl) Context() *SegmentContext {
	return &s.SegmentContext
}

func (s *segmentSpanImpl) Refs() []*propagation.SpanContext {
	return s.defaultSpan.Refs
}

func (s *segmentSpanImpl) StartTime() int64 {
	return tool.Millisecond(s.defaultSpan.StartTime)
}

func (s *segmentSpanImpl) EndTime() int64 {
	return tool.Millisecond(s.defaultSpan.EndTime)
}

func (s *segmentSpanImpl) OperationName() string {
	return s.defaultSpan.OperationName
}

func (s *segmentSpanImpl) Peer() string {
	return s.defaultSpan.Peer
}

func (s *segmentSpanImpl) SpanType() common.SpanType {
	return common.SpanType(s.defaultSpan.SpanType)
}

func (s *segmentSpanImpl) SpanLayer() common.SpanLayer {
	return s.defaultSpan.Layer
}

func (s *segmentSpanImpl) IsError() bool {
	return s.defaultSpan.IsError
}

func (s *segmentSpanImpl) Tags() []*common.KeyStringValuePair {
	return s.defaultSpan.Tags
}

func (s *segmentSpanImpl) Logs() []*v2.Log {
	return s.defaultSpan.Logs
}

func (s *segmentSpanImpl) ComponentID() int32 {
	return s.defaultSpan.ComponentID
}

func (s *segmentSpanImpl) context() SegmentContext {
	return s.SegmentContext
}

func (s *segmentSpanImpl) segmentRegister() bool {
	for {
		o := atomic.LoadInt32(s.Context().refNum)
		if o < 0 {
			return false
		}
		if atomic.CompareAndSwapInt32(s.Context().refNum, o, o+1) {
			return true
		}
	}
}

func (s *segmentSpanImpl) createSegmentContext(parent segmentSpan) {
	if parent == nil {
		s.SegmentContext = SegmentContext{}
		if len(s.defaultSpan.Refs) > 0 {
			s.TraceID = s.defaultSpan.Refs[0].TraceID
		} else {
			s.TraceID = idgen.GenerateGlobalID()
		}
	} else {
		s.SegmentContext = parent.context()
		s.ParentSegmentID = s.SegmentID
		s.ParentSpanID = s.SpanID
		s.SpanID = atomic.AddInt32(s.Context().spanIDGenerator, 1)
	}
}

type rootSegmentSpan struct {
	*segmentSpanImpl
	notify  <-chan ReportedSpan
	segment []ReportedSpan
	doneCh  chan int32
}

func (rs *rootSegmentSpan) End() {
	rs.defaultSpan.End()
	go func() {
		rs.doneCh <- atomic.SwapInt32(rs.Context().refNum, -1)
	}()
}

func (rs *rootSegmentSpan) createRootSegmentContext(parent segmentSpan) {
	rs.SegmentID = idgen.GenerateScopedGlobalID(int64(rs.tracer.instanceID))
	i := int32(0)
	rs.spanIDGenerator = &i
	rs.SpanID = i
	rs.ParentSpanID = -1
}

func newSegmentRoot(segmentSpan *segmentSpanImpl) *rootSegmentSpan {
	s := &rootSegmentSpan{
		segmentSpanImpl: segmentSpan,
	}
	var init int32
	s.refNum = &init
	ch := make(chan ReportedSpan)
	s.collect = ch
	s.notify = ch
	s.segment = make([]ReportedSpan, 0, 10)
	s.doneCh = make(chan int32)
	go func() {
		total := -1
		defer close(ch)
		defer close(s.doneCh)
		for {
			select {
			case span := <-s.notify:
				s.segment = append(s.segment, span)
			case n := <-s.doneCh:
				total = int(n)
			}
			if total == len(s.segment) {
				break
			}
		}
		s.tracer.reporter.Send(append(s.segment, s))
	}()
	return s
}
