package frontend

import (
	"io"
	"net/http"

	"github.com/dailywire/monorepo/v2/cms"
)

type ContentFetcher interface {
	FetchContent() (cms.Content, error)
}

type Renderer interface {
	Render(io.Writer, cms.Content) error
}

type Server struct {
	renderer       Renderer
	contentFetcher ContentFetcher
}

func NewServer(renderer Renderer, contentFetcher ContentFetcher) *Server {
	return &Server{
		renderer:       renderer,
		contentFetcher: contentFetcher,
	}
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	content, err := s.contentFetcher.FetchContent()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "text/html")
	if err := s.renderer.Render(w, content); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}
