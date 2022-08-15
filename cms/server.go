package cms

import (
	"encoding/json"
	"net/http"
)

type ContentGetter interface {
	GetFrontpageArticles() ([]Article, error)
}

type Server struct {
	contentGetter ContentGetter
}

func NewServer(contentFetcher ContentGetter) *Server {
	return &Server{
		contentGetter: contentFetcher,
	}
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	articles, err := s.contentGetter.GetFrontpageArticles()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	content := &Content{
		FrontpageArticles: articles,
	}

	jsn, err := json.Marshal(content)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	if _, err := w.Write(jsn); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}
