package context3

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestServer(t *testing.T) {
	dat := "hello, universe"

	t.Run("returns data from store", func(t *testing.T) {
		sto := &SpyOnStore{response: dat}
		serv := Server(sto)

		req := httptest.NewRequest(http.MethodGet, "/", nil)
		res := httptest.NewRecorder()

		serv.ServeHTTP(res, req)

		if res.Body.String() != dat {
			t.Errorf("got %s, want %s", res.Body.String(), dat)
		}
	})

	t.Run("tells store to cancel work if request is cancelled", func(t *testing.T) {
		sto := &SpyOnStore{response: dat}
		serv := Server(sto)

		req := httptest.NewRequest(http.MethodGet, "/", nil)

		cancelCon, cancel := context.WithCancel(req.Context())
		time.AfterFunc(5*time.Millisecond, cancel)
		req = req.WithContext(cancelCon)

		res := &SpyResponseWrite{}

		serv.ServeHTTP(res, req)

		if res.written {
			t.Errorf("a response should not be written")
		}
	})
}
