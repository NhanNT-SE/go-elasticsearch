package config

type Base struct {
	HTTPAddress string `mapstructure:"http_address"`
}

type FirebaseConfig struct {
	CredentialsJSON string `mapstructure:"credentials_json"`
	CredentialsFile string `mapstructure:"credentials_file"`
}

type MongoConfig struct {
	URI      string `mapstructure:"uri"`
	Database string `mapstructure:"database"`
}

type RedisConfig struct {
	Addr     string `mapstructure:"addr"`
	Password string `mapstructure:"password"`
	User     string `mapstructure:"user"`
	DB       int    `mapstructure:"db"`
}

type KafkaConfig struct {
	BootstrapServers string `mapstructure:"bootstrap_servers"`
	GroupID          string `mapstructure:"group_id"`
	AutoOffsetReset  string `mapstructure:"auto_offset_reset"`
}

type ServerConfig struct {
	Base          `mapstructure:",squash"`
	Firebase      FirebaseConfig      `mapstructure:"firebase"`
	Mongo         MongoConfig         `mapstructure:"mongo"`
	Redis         RedisConfig         `mapstructure:"redis"`
	Elasticsearch ElasticsearchConfig `mapstructure:"elasticsearch"`
}

type EventListenerConfig struct {
	Mongo MongoConfig `mapstructure:"mongo"`
	Redis RedisConfig `mapstructure:"redis"`
	Kafka KafkaConfig `mapstructure:"kafka"`
}

type ElasticsearchConfig struct {
	Addr     string `mapstructure:"addr"`
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
}
