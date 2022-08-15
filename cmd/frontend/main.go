package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/dailywire/monorepo/v2/frontend"
	"github.com/dailywire/monorepo/v2/http"
)

func main() {
	siteGen, err := frontend.NewTemplateSiteGenerator("frontend/templates/index.html")
	require("site generator", err)

	_, shutdown := requireHTTPServer(
		"frontend",
		":5000",
		frontend.NewServer(
			siteGen,
			frontend.NewHTTPContentFetcher("http://localhost:5001"), // Obviously this shouldn't be hardcoded!
		),
	)
	defer shutdown()

	interrupChan := make(chan os.Signal, 1)
	signal.Notify(interrupChan, syscall.SIGINT, syscall.SIGTERM)

	x := <-interrupChan
	log.Print("\nFrontend Exit:", x)
}

func requireHTTPServer(name, address string, handler http.Handler) (svr *http.Server, shutdown func()) {
	svr = http.NewServer(address, handler)
	require(fmt.Sprintf("%s server startup", name), svr.Start())

	log.Print(fmt.Sprintf("%s server started on ", name), address)

	return svr, func() {
		log.Print(fmt.Sprintf("%s server shutting down", name))
		assert(fmt.Sprintf("%s server shutdown", name), svr.Stop())
	}
}

func require(desc string, err error) {
	if err != nil {
		log.Print(fmt.Sprintf("%s: %s", desc, err))
		os.Exit(1)
	}
}

func assert(desc string, err error) {
	if err != nil {
		log.Print(fmt.Sprintf("%s: %s", desc, err))
		os.Exit(1)
	}
}
