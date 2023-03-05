package main

import (
	"bytes"
	"testing"
)

func TestGreeting(t *testing.T) {
	buffer := bytes.Buffer{}
	Greeting(&buffer, "Chiris")

	got := buffer.String()
	want := "Hello, Chiris"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
