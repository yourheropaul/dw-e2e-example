package frontend

import (
	"html/template"
	"io"
	"io/fs"

	"github.com/dailywire/monorepo/v2/cms"
)

type TemplateSiteGenerator struct {
	filesystem fs.FS
}

func NewTemplateSiteGenerator(filesystem fs.FS) (*TemplateSiteGenerator, error) {
	return &TemplateSiteGenerator{
		filesystem: filesystem,
	}, nil
}

func (t *TemplateSiteGenerator) Render(dest io.Writer, content cms.Content) error {
	tmpl, err := template.ParseFS(t.filesystem, "templates/*")
	if err != nil {
		return err
	}

	args := templateArgsFromContent(content)

	return tmpl.ExecuteTemplate(dest, "index.html", args)
}

type TemplateArgs struct {
	TopStory        cms.Article
	FeaturedStories []cms.Article
}

func templateArgsFromContent(input cms.Content) (output TemplateArgs) {
	output.TopStory = input.FrontpageArticles[0]
	output.FeaturedStories = input.FrontpageArticles[1:][:4]
	return
}
