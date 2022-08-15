package e2e

func (w *World) iShouldSeeTheTopStoryArticlePreviewDisplayedProminently() error {
	if err := w.requires(
		req("database", w.database),
		req("frontend", w.frontendServer),
		req("browser", w.browser),
	); err != nil {
		return err
	}

	newestTopStory, err := w.database.getNewestTopStory()
	if err != nil {
		return err
	}

	return w.frontendServer.assertTopStory(w.browser, newestTopStory)
}
