package main_test

import (
	"fmt"
	"testing"

	"learngo/go-specs-greet/adapters"
	"learngo/go-specs-greet/adapters/webserver"
	"learngo/go-specs-greet/specifications"

	"github.com/alecthomas/assert/v2"
)

func TestGreeterWeb(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}
	var (
		port            = "8081"
		driver, cleanup = webserver.NewDriver(fmt.Sprintf("http://localhost:%s", port))
	)

	t.Cleanup(func() {
		assert.NoError(t, cleanup())
	})

	adapters.StartDockerServer(t, port, "webserver")
	specifications.CurseSpecification(t, driver)
	specifications.GreetSpecification(t, driver)
}
