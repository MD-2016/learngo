package blogpost

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"strings"
)

type Post struct {
	Title       string
	Description string
	Tags        []string
	Body        string
}

const (
	titleSep = "Title: "
	desSep   = "Description: "
	tagSep   = "Tags: "
)

func newPost(ptb io.Reader) (Post, error) {
	scan := bufio.NewScanner(ptb)
	readMeta := func(tagName string) string {
		scan.Scan()
		return strings.TrimPrefix(scan.Text(), tagName)
	}
	return Post{
		Title:       readMeta(titleSep),
		Description: readMeta(desSep),
		Tags:        strings.Split(readMeta(tagSep), ", "),
		Body:        readBody(scan),
	}, nil
}

func readBody(scan *bufio.Scanner) string {
	scan.Scan()
	buff := bytes.Buffer{}
	for scan.Scan() {
		fmt.Fprintln(&buff, scan.Text())
	}
	return strings.TrimSuffix(buff.String(), "\n")
}
