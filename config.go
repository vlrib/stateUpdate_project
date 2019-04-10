package main

type DBConfig struct {
	Host string
	Port int

	User     string
	Password string
	dbName   string
}

type Config struct {
	DB DBConfig
}

func loadConfig() Config {
	config := Config{}
	config.DB.Host = "localhost"
	config.DB.Port = 5432
	config.DB.User = "postgres"
	config.DB.dbName = "teste_db"
	config.DB.Password = "your-password"

	return config
}
