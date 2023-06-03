package main_test

import (
	"fmt"
	"net/http"
	"testing"
	"time"

	"learngo/go-specs-greet/adapters"
	"learngo/go-specs-greet/adapters/httpserver"
	"learngo/go-specs-greet/specifications"
)

func TestGreeterServer(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}
	var (
		port   = "8080"
		driver = httpserver.Driver{
			BaseURL: fmt.Sprintf("http://localhost:%s", port),
			Client: &http.Client{
				Timeout: 1 * time.Second,
			},
		}
	)

	adapters.StartDockerServer(t, port, "httpserver")
	specifications.GreetSpecification(t, driver)
	specifications.CurseSpecification(t, driver)
}
