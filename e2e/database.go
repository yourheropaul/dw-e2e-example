package e2e

import (
	"fmt"
	"math/rand"

	"github.com/brianvoe/gofakeit"
	"github.com/dailywire/monorepo/v2/cms"
)

type database struct {
	frontpageArticles []cms.Article

	// indexes of top story articles
	topStories []int
}

func (d *database) GetFrontpageArticles() ([]cms.Article, error) {
	return d.frontpageArticles, nil
}

func (d *database) addRandomArticles(number int) {
	for i := 0; i < number; i++ {
		d.frontpageArticles = append(d.frontpageArticles, newArticle())
	}
}

func (d *database) setRandomArticleToTopStory() error {
	if len(d.frontpageArticles) == 0 {
		return fmt.Errorf("no articles")
	}

	randomIndex := rand.Intn(len(d.frontpageArticles))
	d.topStories = append(d.topStories, randomIndex)
	d.frontpageArticles[randomIndex].TopStory = true

	return nil
}

func (d *database) getNewestTopStory() (cms.Article, error) {
	for _, article := range d.frontpageArticles {
		if article.TopStory {
			return article, nil
		}
	}

	return cms.Article{}, fmt.Errorf("no top story set")
}

func newArticle() cms.Article {
	return cms.Article{
		Headline: gofakeit.Question(),
		ByLine:   gofakeit.Name(),
		ImageURL: gofakeit.ImageURL(1920, 1080),
	}
}
