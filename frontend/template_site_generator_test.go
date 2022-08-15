package frontend

import (
	"bytes"
	"testing"

	"github.com/dailywire/monorepo/v2/cms"
	"github.com/stretchr/testify/require"
)

var mockContent = cms.Content{
	FrontpageArticles: []cms.Article{
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
}

func Test_ATemplateSiteGeneratoCanRenderTheHomePage(t *testing.T) {
	generator := NewTemplateSiteGenerator(TemplatesFS)

	buf := bytes.Buffer{}
	generator.Render(&buf, mockContent)
	require.Equal(t, buf.Len(), 736419)
}
