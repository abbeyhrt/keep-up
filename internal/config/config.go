package config

import "os"

type Config struct {
	Host string
	Port string
	Addr string
	Env  string
}

func New() Config {
	host := os.Getenv("HOST")
	if host == "" {
		host = "0.0.0.0"
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	env := os.Getenv("DEPLOY_ENV")
	if env == "" {
		env = "development"
	}

	cfg := Config{
		Host: host,
		Port: port,
		Addr: host + ":" + port,
		Env:  env,
	}

	return cfg
}
