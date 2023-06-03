package interactions_test

import (
	"testing"

	"learngo/go-specs-greet/domain/interactions"
	"learngo/go-specs-greet/specifications"

	"github.com/alecthomas/assert/v2"
)

func TestGreet(t *testing.T) {
	specifications.GreetSpecification(
		t,
		specifications.GreetAdapter(interactions.Greet),
	)

	t.Run("default name to world if it's an empty string", func(t *testing.T) {
		assert.Equal(t, "Hello, World", interactions.Greet(""))
	})
}
