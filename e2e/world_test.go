package e2e

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_AnSuccessSetupFuncDoesntCauseASetupError(t *testing.T) {
	errorTestFunc := func() (shutdownFunc, error) {
		return nil, nil
	}
	w := &World{}
	require.Nil(t, w.requireSetup(errorTestFunc))
}

func Test_AnFailingSetupFuncDoesCauseASetupError(t *testing.T) {
	errorTestFunc := func() (shutdownFunc, error) {
		return nil, fmt.Errorf("error")
	}
	w := &World{}
	require.NotNil(t, w.requireSetup(errorTestFunc))
}
