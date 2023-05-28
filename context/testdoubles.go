package context3

import (
	"context"
	"errors"
	"log"
	"net/http"
	"time"
)

type SpyOnStore struct {
	response string
}

func (s *SpyOnStore) Fetch(ctx context.Context) (string, error) {
	dat := make(chan string, 1)

	go func() {
		var res string
		for _, c := range s.response {
			select {
			case <-ctx.Done():
				log.Println("spy store cancelled")
				return
			default:
				time.Sleep(10 * time.Millisecond)
				res += string(c)
			}
		}
		dat <- res
	}()

	select {
	case <-ctx.Done():
		return "", ctx.Err()
	case res := <-dat:
		return res, nil
	}
}

type SpyResponseWrite struct {
	written bool
}

func (s *SpyResponseWrite) Header() http.Header {
	s.written = true
	return nil
}

func (s *SpyResponseWrite) Write([]byte) (int, error) {
	s.written = true
	return 0, errors.New("not implemented")
}

func (s *SpyResponseWrite) WriteHeader(status int) {
	s.written = true
}
