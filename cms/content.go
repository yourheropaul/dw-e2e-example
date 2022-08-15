package cms

import "time"

type Article struct {
	Headline string
	ImageURL string
	ByLine   string
	Date     *time.Time
	TopStory bool
	Featured bool
}

type Content struct {
	FrontpageArticles []Article
}
