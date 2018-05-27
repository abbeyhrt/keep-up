package config

import (
	"net/url"
	"os"

	"golang.org/x/oauth2"
)

// Provider struct
type Provider struct {
	OAuth    oauth2.Config
	UserInfo url.URL
}

// Config struct for creating server and database connection
type Config struct {
	Host         string
	Port         string
	Addr         string
	Env          string
	Postgres     url.URL
	Google       Provider
	CookieSecret string
}

// New function to create a new instance of a server or database
func New() (Config, error) {
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
	v.Set("connect_timeout", "15")

	pg.RawQuery = v.Encode()

	oauth := oauth2.Config{
		ClientID:     os.Getenv("GOOGLE_CLIENT_ID"),
		ClientSecret: os.Getenv("GOOGLE_CLIENT_SECRET"),
		RedirectURL:  os.Getenv("GOOGLE_REDIRECT_URL"),
		Scopes: []string{
			"https://www.googleapis.com/auth/userinfo.email",
			"https://www.googleapis.com/auth/userinfo.profile",
		},
		Endpoint: oauth2.Endpoint{
			AuthURL:  "https://accounts.google.com/o/oauth2/auth",
			TokenURL: "https://accounts.google.com/o/oauth2/token",
		},
	}

	userinfo, err := url.Parse(os.Getenv("GOOGLE_ACCESS_URL"))
	if err != nil {
		return Config{}, err
	}

	cookie := os.Getenv("COOKIE_SECRET")

	return Config{
		Host:     host,
		Port:     port,
		Addr:     host + ":" + port,
		Env:      env,
		Postgres: pg,
		Google: Provider{
			OAuth:    oauth,
			UserInfo: *userinfo,
		},
		CookieSecret: cookie,
	}, nil
}
