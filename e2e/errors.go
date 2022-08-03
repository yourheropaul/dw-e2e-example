package e2e

import "github.com/cucumber/godog"

func nonPendingError(err error) error {
	if err == godog.ErrPending {
		return nil
	}

	return err
}
