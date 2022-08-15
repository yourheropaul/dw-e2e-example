package e2e

import "fmt"

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

func (w *World) setupDatabase() (shutdownFunc, error) {
	if w.database != nil {
		return nil, nil
	}

	w.database = &database{}
	return func() {
		w.database = nil
	}, nil
}

func (w *World) setupCMSServer() (shutdownFunc, error) {
	if w.cmsServer != nil {
		return nil, nil
	}

	if w.database == nil {
		return nil, fmt.Errorf("database is required")
	}

	w.cmsServer = newCmsServer(w.database)
	return func() {
		w.cmsServer.Close()
		w.cmsServer = nil
	}, nil
}

func (w *World) setupFrontendServer() (shutdownFunc, error) {
	if w.frontendServer != nil {
		return nil, nil
	}

	if w.cmsServer == nil {
		return nil, fmt.Errorf("CMS server is required")
	}

	w.frontendServer = newFrontendServer(w.cmsServer)
	return func() {
		w.frontendServer.Close()
		w.frontendServer = nil
	}, nil
}
