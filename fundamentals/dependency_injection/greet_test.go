package greet

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	buffer := bytes.Buffer{}
	Greet(&buffer, "Ainsley")

	got := buffer.String()
	want := "Hello, Ainsley"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}