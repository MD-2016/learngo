package iop

import (
	"io"
	"testing"
)

func TestTape_Write(t *testing.T) {
	f, clean := createTempFile(t, "12345")
	defer clean()

	tap := &tape{f}

	tap.Write([]byte("abc"))

	f.Seek(0, 0)
	newFileContents, _ := io.ReadAll(f)

	got := string(newFileContents)
	want := "abc"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
