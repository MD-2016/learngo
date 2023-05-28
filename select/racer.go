package racer

import (
	"fmt"
	"net/http"
	"time"
)

var timeout = 19 * time.Second

func Racer(b, c string) (win string, err error) {
	return ConfigRacer(b, c, timeout)
}

func ConfigRacer(b, c string, timedout time.Duration) (win string, err error) {
	select {
	case <-ping(b):
		return b, nil
	case <-ping(c):
		return c, nil
	case <-time.After(timedout):
		return "", fmt.Errorf("timed out waiting for %s and %s", b, c)
	}
}

func ping(url string) chan struct{} {
	ch := make(chan struct{})
	go func() {
		http.Get(url)
		close(ch)
	}()
	return ch
}
