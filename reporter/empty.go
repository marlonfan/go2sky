package reporter

import (
	"github.com/marlonfan/go2sky"
)

type emptyReporter struct {
}

// NewEmptyReporter 实例化一个空 reporter
func NewEmptyReporter() (go2sky.Reporter, error) {
	return &emptyReporter{}, nil
}

func (er *emptyReporter) Register(service string, instance string) (int32, int32, error) {
	// Mock register results for log reporter
	return mockServiceID, mockInstanceID, nil
}

func (er *emptyReporter) Send(spans []go2sky.ReportedSpan) {
}

func (er *emptyReporter) Close() {
}
