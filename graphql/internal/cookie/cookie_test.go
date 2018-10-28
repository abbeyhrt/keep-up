package cookie_test

import (
	"testing"

	"github.com/abbeyhrt/keep-up/graphql/internal/cookie"
)

func TestCookieCreate(t *testing.T) {
	name := "cookie_name"
	value := "cookie_value"
	c := cookie.Create(name, value)

	// Make sure that we can create a cookie by name and value and that all the
	// return values match what we expect
	if c.Name != name {
		t.Fatalf(
			"Expected cookie name to match: %s, instead recieved: %s",
			name,
			c.Name,
		)
	}

	if c.Value != value {
		t.Fatalf(
			"Expected cookie value to match: %s, instead recieved: %s",
			value,
			c.Value,
		)
	}

	if c.Path != "/" {
		t.Fatalf(
			"Expected cookie path to match: /, instead recieved: %s",
			c.Path,
		)
	}

	if c.Secure == false {
		t.Fatal("Expected cookie to be secure")
	}

	if c.HttpOnly == false {
		t.Fatal("Expected cookie to be httpOnly")
	}
}
