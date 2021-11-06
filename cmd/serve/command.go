package serve

import (
	"log"

	"github.com/fasthttp/router"
	"github.com/spf13/cobra"
	"github.com/valyala/fasthttp"
)

var (
	Cmd = &cobra.Command{
		Use:     "serve",
		Short:   "Listen and serve HTTP server as an API gateway",
		PreRunE: preRunE,
		RunE:    runE,
	}
)

func init() {
	Cmd.Flags().String("listen-address", "0.0.0.0:8080", "Specify the host/IP and port to which HTTP server binds for listening")
	Cmd.Flags().String("auth-service-endpoint", "localhost:8081", "Specify the auth-service endpoint")
	Cmd.Flags().String("blog-service-endpoint", "localhost:8082", "Specify the blog-service endpoint")
	Cmd.Flags().String("discussion-service-endpoint", "localhost:8083", "Specify the discussion-service endpoint")
	Cmd.Flags().String("storage-service-endpoint", "localhost:8084", "Specify the storage-service endpoint")
}

func preRunE(cmd *cobra.Command, _ []string) error {
	return nil
}

func runE(cmd *cobra.Command, _ []string) error {
	var (
		listenAddress, _ = cmd.Flags().GetString("listen-address")
	)

	r := router.New()

	log.Printf("[INFO]: server is listening on: %s\n", listenAddress)
	return fasthttp.ListenAndServe(listenAddress, r.Handler)
}
