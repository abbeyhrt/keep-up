package config

import (
	"net/url"
	"os"
)

// Config struct for creating server and database connection
type Config struct {
	Host     string
	Port     string
	Addr     string
	Env      string
	Postgres url.URL
}

// New function to create a new instance of a server or database
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

	sslmode := os.Getenv("DB_SSL_MODE")
	if sslmode == "" {
		sslmode = "require"
	}

	pg := url.URL{
		Scheme: "postgres",
		User: url.UserPassword(
			os.Getenv("DB_USERNAME"),
			os.Getenv("DB_PASSWORD"),
		),
		Host: os.Getenv("DB_ADDRESS"),
		Path: os.Getenv("DB_NAME"),
	}

	v := url.Values{}

	v.Set("sslmode", sslmode)

	pg.RawQuery = v.Encode()

	cfg := Config{
		Host:     host,
		Port:     port,
		Addr:     host + ":" + port,
		Env:      env,
		Postgres: pg,
	}

	return cfg
}
