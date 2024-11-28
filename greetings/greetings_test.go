package greetings

import (
	"testing"
	"regexp"
)

func TestHelloName(t *testing.T) {
	//Given
	name := "Gladys"
	want :=regexp.MustCompile(`\b`+name+`\b`)
	
	//When
	msg, err := Hello("Gladys")

	//Then
	if !want.MatchString(msg) || err != nil {
		t.Fatalf(`Hello("Gladys") = %q, %v, want match for %#q, nil`, msg, err, want)
	}
}

func TestHelloEmpty(t * testing.T) {
	
	//When
	msg, err := Hello("")

	//Then
	if msg != "" || err == nil {
		t.Fatalf(`Hello("") = %q, %v, want "", error`, msg, err)
	}
}
