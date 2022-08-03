package e2e

import (
	"context"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/cucumber/godog"
	"github.com/cucumber/messages-go/v16"
)

type shutdownFunc func()
type setupFunc func() (shutdown shutdownFunc, err error)

type World struct {
	scenario           *godog.Scenario
	stepIndex          int
	logger             *Log
	scenarioResetFuncs []shutdownFunc
	browser            *Browser
}

func NewWorld() *World {
	return &World{}
}

func (w *World) BeforeScenario(ctx context.Context, scenario *godog.Scenario) (context.Context, error) {
	w.scenario = scenario

	combinedLogFile, err := w.logFileForScenario("combined")
	if err != nil {
		return ctx, err
	}

	outputs := []io.WriteCloser{combinedLogFile}

	if _, present := os.LookupEnv("LOG_TO_STDOUT"); present {
		outputs = append(outputs, os.Stdout)
	}

	w.logger = NewLog(outputs...)

	return ctx, nil
}

func (w *World) AfterScenario(ctx context.Context, scenario *godog.Scenario, scenarioError error) (context.Context, error) {
	if err := w.reset(); err != nil {
		return ctx, err
	}

	return ctx, nonPendingError(scenarioError)
}

func (w *World) AfterStep(step *messages.PickleStep) {
	w.stepIndex++

	if w.browser != nil {
		w.screenshotForStep(step)
	}
}

func (w *World) screenshotForStep(step *messages.PickleStep) {
	screenshot, err := w.browser.FullScreenshot()
	if err != nil {
		log.Printf("screenshot error: %s", err)
		return
	}

	filename := fmt.Sprintf(
		"e2e-screenshots/%s/%s/%02d-%s.png",
		strings.TrimPrefix(w.scenario.Uri, "features/"),
		w.scenario.Name,
		w.stepIndex,
		step.Text,
	)

	if err := os.MkdirAll(filepath.Dir(filename), os.ModePerm); err != nil {
		log.Printf("failed to create dir: %s", err)
		return
	}

	if err := ioutil.WriteFile(filename, screenshot, os.ModePerm); err != nil {
		log.Printf("screenshot error: %s", err)
		return
	}
}

func (w *World) logFileForScenario(serviceName string) (io.WriteCloser, error) {
	filename := fmt.Sprintf(
		"e2e-logs/%s/%s/%s.log",
		strings.TrimPrefix(w.scenario.Uri, "features/"),
		w.scenario.Name,
		serviceName,
	)

	if err := os.MkdirAll(filepath.Dir(filename), os.ModePerm); err != nil {
		return nil, fmt.Errorf("Failed to create dir: %w", err)
	}

	return os.OpenFile(filename, os.O_CREATE|os.O_WRONLY, os.ModePerm)
}

func (w *World) loggerForScenario(serviceName string) (*ServiceLogger, error) {
	f, err := w.logFileForScenario(serviceName)
	if err != nil {
		return nil, err
	}

	return w.logger.NewServiceLogger(serviceName, f), nil
}

func (w *World) reset() error {
	w.scenario = nil
	w.stepIndex = 0

	for _, resetFunc := range w.scenarioResetFuncs {
		resetFunc()
	}

	w.scenarioResetFuncs = []shutdownFunc{}

	return w.logger.Close()
}

func (w *World) requireSetup(setups ...setupFunc) error {
	for _, setup := range setups {
		reset, err := setup()

		if err != nil {
			return err
		}

		if reset != nil {
			w.scenarioResetFuncs = append(w.scenarioResetFuncs, reset)
		}
	}

	return nil
}

type requirement struct {
	err     string
	subject interface{}
}

func req(err string, subject interface{}) requirement {
	return requirement{err: err, subject: subject}
}

func (w *World) requires(requirements ...requirement) error {
	for _, requirement := range requirements {
		if requirement.subject == nil {
			return fmt.Errorf("requirement not met: %s", requirement.err)
		}
	}

	return nil
}
