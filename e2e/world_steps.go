package e2e

type stepper interface {
	Step(match interface{}, function interface{})
}

func (w *World) RegisterSteps(sc stepper) {
	// Given
	sc.Step(`^there a several articles already posted$`, w.thereASeveralArticlesAlreadyPosted)
	sc.Step(`^an article has been promoted to the Top Story$`, w.anArticleHasBeenPromotedToTheTopStory)

	// When
	sc.Step(`^I visit the home page$`, w.iVisitTheHomePage)

	// Then
	sc.Step(`^I should see the Top Story article preview displayed prominently$`, w.iShouldSeeTheTopStoryArticlePreviewDisplayedProminently)
}
