package e2e

import (
	"github.com/pkg/errors"
)

func (w *World) iViewTheBlockExplorerForMyAccount() error {
	if err := w.requireSetup(
		w.setupBrowser,
	); err != nil {
		return err
	}

	if err := w.browser.Visit("/"); err != nil {
		return errors.Wrap(err, "browser.Visit")
	}

	return nil
}
