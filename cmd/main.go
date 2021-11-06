package main

import (
	"fmt"
	"log"

	"github.com/nomkhonwaan/myblog-api-gateway/cmd/serve"
	"github.com/spf13/cobra"
)

var (
	// Version refers to the latest Git tag version.
	Version = "v0.0.1"

	// Revision refers to the latest Git commit hash.
	Revision = "development"
)

func main() {
	cmd := cobra.Command{Version: fmt.Sprintf("%s %s", Version, Revision)}
	cmd.AddCommand(serve.Cmd)

	if err := cmd.Execute(); err != nil {
		log.Fatalf("myblog-api-gateway: %s", err)
	}
}
