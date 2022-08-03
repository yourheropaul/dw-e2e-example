package e2e

import (
	"github.com/cucumber/godog"
	"github.com/dailywire/monorepo/v2/e2e"
)

func InitializeScenario(ctx *godog.ScenarioContext) {
	w := e2e.NewWorld()

	ctx.Before(w.BeforeScenario)
	ctx.After(w.AfterScenario)
	ctx.BeforeStep(w.AfterStep)
	w.RegisterSteps(ctx)
}
