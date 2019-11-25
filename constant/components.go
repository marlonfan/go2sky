package constant

const (
	Tomcat               = 1
	HttpClient           = 2
	Dubbo                = 3
	H2                   = 4
	Mysql                = 5
	ORACLE               = 6
	Redis                = 7
	Motan                = 8
	MongoDB              = 9
	Resin                = 10
	Feign                = 11
	OKHttp               = 12
	SpringRestTemplate   = 13
	SpringMVC            = 14
	Struts2              = 15
	NutzMVC              = 16
	NutzHttp             = 17
	JettyClient          = 18
	JettyServer          = 19
	Memcached            = 20
	ShardingJDBC         = 21
	PostgreSQL           = 22
	GRPC                 = 23
	ElasticJob           = 24
	RocketMQ             = 25
	httpasyncclient      = 26
	Kafka                = 27
	ServiceComb          = 28
	Hystrix              = 29
	Jedis                = 30
	SQLite               = 31
	h2JdbcDriver         = 32
	MysqlConnectorJava   = 33
	ojdbc                = 34
	Spymemcached         = 35
	Xmemcached           = 36
	PostgresqlJdbcDriver = 37
	RocketMQProducer     = 38
	RocketMQConsumer     = 39
	KafkaProducer        = 40
	KafkaConsumer        = 41
	MongodbDriver        = 42
	SOFARPC              = 43
	ActiveMQ             = 44
	ActivemqProducer     = 45
	ActivemqConsumer     = 46
	Elasticsearch        = 47
	TransportClient      = 48
	Tttp                 = 49
	Rpc                  = 50
	RabbitMQ             = 51
	RabbitmqProducer     = 52
	RabbitmqConsumer     = 53
	Canal                = 54
	Gson                 = 55
	Redisson             = 56
	Lettuce              = 57
	Zookeeper            = 58
	Vertx                = 59
	ShardingSphere       = 60
	SpringCloudGateway   = 61
	RESTEasy             = 62
	SolrJ                = 63
	Solr                 = 64
	SpringAsync          = 65
	JdkHttp              = 66
	SpringWebflux        = 67
	Play                 = 68
	CassandraJavaDriver  = 69
	Cassandra            = 70
	Light4J              = 71
	Pulsar               = 72
	PulsarProducer       = 73
	PulsarConsumer       = 74
	Ehcache              = 75
	SocketIO             = 76
	RestHighLevelClient  = 77
	SpringTx             = 78

	// .NET/.NET Core components
	// [3000, 4000) for C#/.NET only
	AspNetCore                          = 3001
	EntityFrameworkCore                 = 3002
	SqlClient                           = 3003
	CAP                                 = 3004
	StackExchangeRedis                  = 3005
	SqlServer                           = 3006
	Npgsql                              = 3007
	MySqlConnector                      = 3008
	EntityFrameworkCoreInMemory         = 3009
	EntityFrameworkCoreSqlServer        = 3010
	EntityFrameworkCoreSqlite           = 3011
	PomeloEntityFrameworkCoreMySql      = 3012
	NpgsqlEntityFrameworkCorePostgreSQL = 3013
	InMemoryDatabase                    = 3014
	AspNet                              = 3015
	SmartSql                            = 3016

	// NoeJS components
	// [4000, 5000) for Node.js agent
	HttpServer = 4001
	express    = 4002
	Egg        = 4003
	Koa        = 4004

	// Golang components
	// [5000, 6000) for Golang agent
	ServiceCombMesher        = 5001
	ServiceCombServiceCenter = 5002

// Component Server mapping defines the server display names of some components
// e.g.
// Jedis is a client library in Java for Redis server
// Component-Server-Mappings:
//   mongodb-driver: MongoDB
//   rocketMQ-producer: RocketMQ
//   rocketMQ-consumer: RocketMQ
//   kafka-producer: Kafka
//   kafka-consumer: Kafka
//   activemq-producer: ActiveMQ
//   activemq-consumer: ActiveMQ
//   rabbitmq-producer: RabbitMQ
//   rabbitmq-consumer: RabbitMQ
//   postgresql-jdbc-driver: PostgreSQL
//   Xmemcached: Memcached
//   Spymemcached: Memcached
//   h2-jdbc-driver: H2
//   mysql-connector-java: Mysql
//   Jedis: Redis
//   StackExchange.Redis: Redis
//   Redisson: Redis
//   Lettuce: Redis
//   Zookeeper: Zookeeper
//   SqlClient: SqlServer
//   Npgsql: PostgreSQL
//   MySqlConnector: Mysql
//   EntityFrameworkCore.InMemory: InMemoryDatabase
//   EntityFrameworkCore.SqlServer: SqlServer
//   EntityFrameworkCore.Sqlite: SQLite
//   Pomelo.EntityFrameworkCore.MySql: Mysql
//   Npgsql.EntityFrameworkCore.PostgreSQL: PostgreSQL
//   transport-client: Elasticsearch
//   SolrJ: Solr
//   cassandra-java-driver: Cassandra
//   pulsar-producer: Pulsar
//   pulsar-consumer: Pulsar
//   rest-high-level-client: Elasticsearc
)
