package racer

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestRacer(t *testing.T) {
	t.Run("compares server speeds, returning fastest url", func(t *testing.T) {
		slowServ := makeDelayServer(20 * time.Millisecond)
		fastServ := makeDelayServer(0 * time.Millisecond)

		defer slowServ.Close()
		defer fastServ.Close()

		slowURL := slowServ.URL
		fastURL := fastServ.URL

		want := fastURL
		got, err := Racer(slowURL, fastURL)

		if err != nil {
			t.Fatalf("did not expect error but got one %v", err)
		}

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})

	t.Run("returns error if server doesn't respond in 10s", func(t *testing.T) {
		serv := makeDelayServer(25 * time.Millisecond)

		defer serv.Close()

		_, err := ConfigRacer(serv.URL, serv.URL, 20*time.Millisecond)

		if err == nil {
			t.Errorf("expected an error but didn't get it")
		}
	})
}

func makeDelayServer(delay time.Duration) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(delay)
		w.WriteHeader(http.StatusOK)
	}))
}
