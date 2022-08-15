package frontend

import (
	"encoding/json"
	"net/http"

	"github.com/dailywire/monorepo/v2/cms"
)

type HTTPContentFetcher struct {
	URL string
}

func NewHTTPContentFetcher(URL string) *HTTPContentFetcher {
	return &HTTPContentFetcher{
		URL: URL,
	}
}

func (h HTTPContentFetcher) FetchContent() (cms.Content, error) {
	res, err := http.Get(h.URL)
	if err != nil {
		return cms.Content{}, err
	}

	defer res.Body.Close()

	content := cms.Content{}
	err = json.NewDecoder(res.Body).Decode(&content)

	return content, err
}
