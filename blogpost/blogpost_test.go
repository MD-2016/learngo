package blogpost_test

import (
	"blogpost"
	"reflect"
	"testing"
	"testing/fstest"
)

func TestNewBlogPosts(t *testing.T) {
	const (
		firstBody = `Title: Post 1
Description: Description 1
Tags: tdd, go
---
Hello
World`
		secondBody = `Title: Post 2
Description: Description 2
Tags: rust, borrow-checker
---
B
L
M`
	)

	fs := fstest.MapFS{
		"hello world.md":  {Data: []byte(firstBody)},
		"hello-world2.md": {Data: []byte(secondBody)},
	}

	pts, err := blogpost.NewFSPosts(fs)

	assertNoError(t, err)

	assertPostLen(t, pts, fs)

	assertPost(t, pts[0], blogpost.Post{
		Title:       "Post 1",
		Description: "Description 1",
		Tags:        []string{"tdd", "go"},
		Body: `Hello
World`,
	})
}

func assertNoError(t *testing.T, err error) {
	t.Helper()
	if err != nil {
		t.Fatal(err)
	}
}

func assertPostLen(t *testing.T, pts []blogpost.Post, fs fstest.MapFS) {
	t.Helper()
	if len(pts) != len(fs) {
		t.Errorf("got %d posts, wanted %d posts", len(pts), len(fs))
	}
}

func assertPost(t *testing.T, got blogpost.Post, want blogpost.Post) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %+v, want %+v", got, want)
	}
}

//test for update
