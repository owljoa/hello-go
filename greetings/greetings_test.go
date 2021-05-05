package greetings

import (
	"regexp"
	"testing"
)

// TestHelloName calls greetings.Hello with a name, checking
// for a valid return value.
func TestHelloName(t *testing.T) {
	name := "Owljoa"
	want := regexp.MustCompile(`\b` + name + `\b`)
	message, err := Hello("Owljoa")
	if !want.MatchString(message) || err != nil {
		t.Fatalf(`Hello("Owljoa") = %q, %v, want match for %#q nil`, message, err, want)
	}
}

// TestHelloEmpty calls greetings.Hello with an empty string
// checking for an error.
func TestHelloEmpty(t *testing.T) {
	message, err := Hello("")
	if message != "" || err == nil {
		t.Fatalf(`Hello("") = %q, %v, want "", error`, message, err)
	}
}
