package httpserver

import (
	"fmt"
	"net/http"

	"github.com/MD-2016/go-specs-greet/domain/interactions"
)

const (
	greetPath = "/greet"
	cursePath = "/curse"
)

func replyWith(f func(name string) (interaction string)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		name := r.URL.Query().Get("name")
		fmt.Fprint(w, f(name))
	}
}

func NewHandler() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc(greetPath, replyWith(interactions.Greet))
	mux.HandleFunc(cursePath, replyWith(interactions.Curse))
}
