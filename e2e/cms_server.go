package e2e

import (
	"net/http/httptest"

	"github.com/dailywire/monorepo/v2/cms"
)

type cmsServer struct {
	database   *database
	httpServer *httptest.Server
}

func newCmsServer(database *database) *cmsServer {
	server := cms.NewServer(database)
	return &cmsServer{
		database:   database,
		httpServer: httptest.NewServer(server),
	}
}

func (c *cmsServer) Close() {
	c.httpServer.Close()
}
