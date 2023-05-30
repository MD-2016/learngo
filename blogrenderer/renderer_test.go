package blogrenderer_test

import (
	"blogrenderer"
	"bytes"
	"io"
	"testing"

	approvals "github.com/approvals/go-approval-tests"
)

func TestRender(t *testing.T) {
	var (
		pt = blogrenderer.Post{
			Title: "hello world",
			Body: `# First recipe!
Welcome to my **amazing blog**. I am going to write about my family recipes, and make sure I write a long, irrelevant and boring story about my family before you get to the actual instructions.`,
			Description: "This is a description",
			Tags:        []string{"go", "tdd"},
		}
	)

	postRender, err := blogrenderer.NewPostRend()

	if err != nil {
		t.Fatal(err)
	}
	t.Run("converts post into HTML", func(t *testing.T) {
		buff := bytes.Buffer{}

		if err := postRender.Render(&buff, pt); err != nil {
			t.Fatal(err)
		}

		approvals.VerifyString(t, buff.String())
	})

	t.Run("renders index of posts", func(t *testing.T) {
		buff := bytes.Buffer{}
		pts := []blogrenderer.Post{{Title: "Hello World"}, {Title: "Hello World 2"}}

		if err := postRender.RenderIndex(&buff, pts); err != nil {
			t.Fatal(err)
		}

		approvals.VerifyString(t, buff.String())
	})
}

func BenchmarkRender(b *testing.B) {
	var (
		pt = blogrenderer.Post{
			Title:       "hello world",
			Body:        "This is a post",
			Description: "This is a description",
			Tags:        []string{"go", "tdd"},
		}
	)

	ptRend, err := blogrenderer.NewPostRend()

	if err != nil {
		b.Fatal(err)
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		ptRend.Render(io.Discard, pt)
	}
}
