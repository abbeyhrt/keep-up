package config_test

import (
	"net/url"
	"os"
	"testing"

	"github.com/abbeyhrt/keep-up-graphql/internal/config"
)

func TestDefaultConfig(t *testing.T) {
	os.Clearenv()

	cfg, _ := config.New()
	expected := struct {
		port string
		host string
		addr string
	}{"3000", "0.0.0.0", "0.0.0.0:3000"}

	if cfg.Port != expected.port {
		t.Errorf("expected port '%s', instead got '%s'", expected.port, cfg.Port)
	}

	if cfg.Host != expected.host {
		t.Errorf("expected host '%s', instead got '%s'", expected.host, cfg.Host)
	}

	if cfg.Addr != expected.addr {
		t.Errorf("expected addr '%s', instead got '%s'", expected.addr, cfg.Addr)
	}
}

func TestConfigFromEnv(t *testing.T) {
	os.Clearenv()

	expected := struct {
		port string
		host string
		addr string
	}{"3000", "0.0.0.0", "0.0.0.0:3000"}

	os.Setenv("PORT", expected.port)
	os.Setenv("HOST", expected.host)

	cfg, _ := config.New()

	if cfg.Port != expected.port {
		t.Errorf("expected port '%s', instead got '%s'", expected.port, cfg.Port)
	}

	if cfg.Host != expected.host {
		t.Errorf("expected host '%s', instead got '%s'", expected.host, cfg.Host)
	}

	if cfg.Addr != expected.addr {
		t.Errorf("expected addr '%s', instead got '%s'", expected.addr, cfg.Addr)
	}
}

func TestDatabaseConfig(t *testing.T) {
	os.Clearenv()

	expected := struct {
		sslmode  string
		username string
		password string
		name     string
		address  string
	}{
		"require",
		"user",
		"password",
		"postgres",
		"bumblebee",
	}

	os.Setenv("DB_SSL_MODE", expected.sslmode)
	os.Setenv("DB_USERNAME", expected.username)
	os.Setenv("DB_PASSWORD", expected.password)
	os.Setenv("DB_NAME", expected.name)
	os.Setenv("DB_ADDRESS", expected.address)

	cfg, _ := config.New()

	if cfg.Postgres.Host != expected.address {
		t.Errorf(
			"expected host: %s, instead got: %s",
			expected.address,
			cfg.Postgres.Host,
		)
	}

	user := url.UserPassword(expected.username, expected.password)
	if cfg.Postgres.User.String() != user.String() {
		t.Errorf(
			"expected user: %s, instead got: %s",
			user.String(),
			cfg.Postgres.User.String(),
		)
	}

	if cfg.Postgres.Path != expected.name {
		t.Errorf(
			"expected host: %s, instead got: %s",
			expected.name,
			cfg.Postgres.Path,
		)
	}

	v := url.Values{}

	v.Set("sslmode", expected.sslmode)
	v.Set("connect_timeout", "15")

	rawQuery := v.Encode()

	if cfg.Postgres.RawQuery != rawQuery {
		t.Errorf(
			"expected sslmode: %s, instead got: %s",
			rawQuery,
			cfg.Postgres.RawQuery,
		)
	}
}
