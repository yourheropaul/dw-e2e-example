package e2e

func (w *World) setupBrowser() (shutdownFunc, error) {
	if w.browser != nil {
		return nil, nil
	}

	logger, err := w.loggerForScenario("browser-console")
	if err != nil {
		return nil, err
	}

	w.browser = NewBrowser(logger)
	return func() {
		w.browser.Close()
		w.browser = nil
	}, nil
}
