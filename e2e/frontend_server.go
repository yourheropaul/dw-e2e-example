package e2e

import (
	"net/http/httptest"

	"github.com/dailywire/monorepo/v2/cms"
	"github.com/dailywire/monorepo/v2/frontend"
	"github.com/stretchr/testify/assert"
)

type frontendServer struct {
	cmsServer  *cmsServer
	httpServer *httptest.Server
}

func newFrontendServer(cmsServer *cmsServer) *frontendServer {
	return &frontendServer{
		httpServer: httptest.NewServer(
			frontend.NewServer(
				frontend.NewTemplateSiteGenerator(frontend.TemplatesFS),
				frontend.NewHTTPContentFetcher(cmsServer.httpServer.URL),
			),
		),
	}
}

// This is a very basic, example way to assert the top story!
func (f *frontendServer) assertTopStory(b *browser, article cms.Article) error {
	headline, err := b.Text("//article[@id='top-story']//h2")
	if err != nil {
		return err
	}
	if err := assertExpectedAndActual(assert.Equal, article.Headline, headline); err != nil {
		return err
	}

	byline, err := b.Text("//article[@id='top-story']/a/span")
	if err != nil {
		return err
	}
	if err := assertExpectedAndActual(assert.Equal, "By "+article.ByLine, byline); err != nil {
		return err
	}

	return nil
}

func (c *frontendServer) Close() {
	c.httpServer.Close()
}
