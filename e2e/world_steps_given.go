package e2e

func (w *World) thereASeveralArticlesAlreadyPosted() error {
	if err := w.requireSetup(
		w.setupDatabase,
	); err != nil {
		return err
	}

	w.database.addRandomArticles(10)

	return nil
}

func (w *World) anArticleHasBeenPromotedToTheTopStory() error {
	if err := w.requireSetup(
		w.setupDatabase,
	); err != nil {
		return err
	}

	return w.database.setRandomArticleToTopStory()
}
