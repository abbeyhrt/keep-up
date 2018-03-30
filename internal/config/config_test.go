package config_test

import (
	"os"
	"testing"

	"github.com/abbeyhrt/keep-up-graphql/internal/config"
)

func TestDefaultConfig(t *testing.T) {
	os.Clearenv()

	cfg := config.New()
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

	cfg := config.New()

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
