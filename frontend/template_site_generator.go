package frontend

import (
	"html/template"
	"io"

	"github.com/dailywire/monorepo/v2/cms"
)

type TemplateSiteGenerator struct {
	templatePath string
}

func NewTemplateSiteGenerator(templatePath string) (*TemplateSiteGenerator, error) {
	return &TemplateSiteGenerator{
		templatePath: templatePath,
	}, nil
}

func (t *TemplateSiteGenerator) Render(dest io.Writer, content cms.Content) error {
	tmpl, err := template.ParseFiles(t.templatePath)
	if err != nil {
		return err
	}

	args := templateArgsFromContent(content)

	return tmpl.Execute(dest, args)
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
