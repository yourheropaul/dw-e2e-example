package cms

type Database struct{}

// This is a fake database, obviously!
func (m Database) GetFrontpageArticles() ([]Article, error) {
	return []Article{
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
	}, nil
}
