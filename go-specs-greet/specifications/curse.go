package specifications

import (
	"testing"

	"github.com/alecthomas/assert/v2"
)

type MeanGreeter interface {
	Curse(name string) (string, error)
}

func CurseSpecification(t *testing.T, mean MeanGreeter) {
	got, err := mean.Curse("MD")
	assert.NoError(t, err)
	assert.Equal(t, got, "Go to hell, MD!")
}
