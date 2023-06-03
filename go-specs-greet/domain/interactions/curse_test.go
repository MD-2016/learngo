package interactions_test

import (
	"testing"

	"learngo/go-specs-greet/domain/interactions"
	"learngo/go-specs-greet/specifications"
)

func TestCurse(t *testing.T) {
	specifications.CurseSpecification(
		t,
		specifications.CurseAdapter(interactions.Curse),
	)
}
