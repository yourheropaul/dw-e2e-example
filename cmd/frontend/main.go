package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/dailywire/monorepo/v2/cms"
	"github.com/dailywire/monorepo/v2/frontend"
	"github.com/dailywire/monorepo/v2/http"
)

func main() {
	siteGen, err := frontend.NewTemplateSiteGenerator("frontend/templates/index.html")
	require("site generator", err)

	_, shutdown := requireHTTPServer(
		"webserver",
		":5000",
		frontend.NewServer(siteGen, &mockContentFetcher{}), // TODO: CMS server
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

type mockContentFetcher struct{}

func (m mockContentFetcher) FetchContent() (cms.Content, error) {
	return cms.Content{
		Articles: []cms.Article{
			{
				Headline: "This is a headline",
				ByLine:   "Some Author",
				ImageURL: "https://i.picsum.photos/id/1062/1920/1080.jpg?hmac=BwtGVMQ3zWdOyaoIN3-8Cm41N1f-Ey9OmMcicGmwyVA",
				TopStory: true,
			},
			{
				Headline: "This is another headline",
				ByLine:   "Some Author",
				ImageURL: "https://i.picsum.photos/id/1062/1920/1080.jpg?hmac=BwtGVMQ3zWdOyaoIN3-8Cm41N1f-Ey9OmMcicGmwyVA",
				Featured: true,
			},
			{
				Headline: "3This is a headline",
				ByLine:   "4Some Author",
				ImageURL: "https://i.picsum.photos/id/1062/1920/1080.jpg?hmac=BwtGVMQ3zWdOyaoIN3-8Cm41N1f-Ey9OmMcicGmwyVA",
				Featured: true,
			},
			{
				Headline: "This is a headline",
				ByLine:   "Some Author",
				ImageURL: "https://i.picsum.photos/id/1062/1920/1080.jpg?hmac=BwtGVMQ3zWdOyaoIN3-8Cm41N1f-Ey9OmMcicGmwyVA",
				Featured: true,
			},
			{
				Headline: "--This is a headline",
				ByLine:   "Some Author",
				ImageURL: "https://i.picsum.photos/id/1062/1920/1080.jpg?hmac=BwtGVMQ3zWdOyaoIN3-8Cm41N1f-Ey9OmMcicGmwyVA",
			},
			{
				Headline: "@@This is a headline",
				ByLine:   "Some Author",
				ImageURL: "https://i.picsum.photos/id/1062/1920/1080.jpg?hmac=BwtGVMQ3zWdOyaoIN3-8Cm41N1f-Ey9OmMcicGmwyVA",
				Featured: true,
			},
		},
	}, nil
}
