package e2e

func (w *World) iVisitTheHomePage() error {
	if err := w.requireSetup(
		w.setupCMSServer,
		w.setupFrontendServer,
		w.setupBrowser,
	); err != nil {
		return err
	}

	return w.browser.Visit(w.frontendServer.httpServer.URL)
}
