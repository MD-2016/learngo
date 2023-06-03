package main_test

import (
	"fmt"
	"testing"

	"github.com/MD-2016/go-specs-greet/adapters"
	"github.com/MD-2016/go-specs-greet/adapters/grpcserver"
	"github.com/MD-2016/go-specs-greet/specifications"
)

func TestGreeterServer(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}
	var (
		port   = "50051"
		driver = grpcserver.Driver{Addr: fmt.Sprintf("localhost:%s", port)}
	)
	t.Cleanup(driver.Close)
	adapters.StartDockerServer(t, port, "grpcserver")
	specifications.GreetSpecification(t, &driver)
	specifications.CurseSpecification(t, &driver)
}
