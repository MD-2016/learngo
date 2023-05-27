package di

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	buff := bytes.Buffer{}
	Greet(&buff, "MD")

	got := buff.String()
	want := "Hello, MD"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
