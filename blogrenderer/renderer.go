package blogrenderer

import (
	"embed"
	"html/template"
	"io"

	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/parser"
)

var (
	//go:embed "templates/*"
	postTemp embed.FS
)

type PostRender struct {
	temp    *template.Template
	mdParse *parser.Parser
}

func NewPostRend() (*PostRender, error) {
	temp, err := template.ParseFS(postTemp, "templates/*.gohtml")
	if err != nil {
		return nil, err
	}

	ext := parser.CommonExtensions | parser.AutoHeadingIDs
	parse := parser.NewWithExtensions(ext)

	return &PostRender{temp: temp, mdParse: parse}, nil
}

func (r *PostRender) Render(w io.Writer, p Post) error {
	return r.temp.ExecuteTemplate(w, "blog.gohtml", newPostVM(p, r))
}

func (r *PostRender) RenderIndex(w io.Writer, pts []Post) error {
	return r.temp.ExecuteTemplate(w, "index.gohtml", pts)
}

type PostViewModel struct {
	Post
	HTMLBody template.HTML
}

func newPostVM(p Post, r *PostRender) PostViewModel {
	vm := PostViewModel{Post: p}
	vm.HTMLBody = template.HTML(markdown.ToHTML([]byte(p.Body), r.mdParse, nil))
	return vm
}
