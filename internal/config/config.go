package config

type ServerConfig struct {
	Port string
}

type RedisConfig struct {
	PubSubChannel string
	Endpoints     string
	Database      int
}
