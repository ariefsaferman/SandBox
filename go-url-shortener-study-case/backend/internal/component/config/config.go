package config

type Config struct {
	Application AppConfig
	Database    DBConfig
	Redis       RedisConfig
	Asynq       AsynqConfig
}

type RedisConfig struct {
	Host string
	Port int
}

type AppConfig struct {
	Environment string
	Port        int
	BaseUrl     string
	SlugLength  int
}

type DBConfig struct {
	Host     string
	Port     int
	Password string
	User     string
	Name     string
}

type AsynqConfig struct {
	ConcurrentWorker int
	QueueCritical    int
	QueueDefault     int
	QueueLow         int
}
