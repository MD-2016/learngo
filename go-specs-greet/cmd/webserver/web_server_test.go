package main_test

import (
	"fmt"
	"testing"

	"github.com/MD-2016/go-specs-greet/adapters"
	"github.com/MD-2016/go-specs-greet/adapters/webserver"
	"github.com/MD-2016/go-specs-greet/specifications"
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
