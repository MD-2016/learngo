package context3

import (
	"context"
	"errors"
	"fmt"
	"net/http"
)

type Store interface {
	Fetch(cont context.Context) (string, error)
}

func Server(sto Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		dat, err := sto.Fetch(r.Context())

		if err != nil {
			fmt.Sprintln(errors.New("not able to fetch"))
			return
		}

		fmt.Fprintf(w, dat)
	}
}
