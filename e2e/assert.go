package e2e

import (
	"fmt"

	"github.com/stretchr/testify/assert"
)

// assertExpectedAndActual is a helper function to allow the step function to call
// assertion functions where you want to compare an expected and an actual value.
// nolint: unused, deadcode
func assertExpectedAndActual(a expectedAndActualAssertion, expected, actual interface{}, msgAndArgs ...interface{}) error {
	var t asserter
	a(&t, expected, actual, msgAndArgs...)
	return t.err
}

// nolint: unused, deadcode
type expectedAndActualAssertion func(t assert.TestingT, expected, actual interface{}, msgAndArgs ...interface{}) bool

// assertActual is a helper function to allow the step function to call
// assertion functions where you want to compare an actual value to a
// predined state like nil, empty or true/false.
func assertActual(a actualAssertion, got, expected interface{}, msgAndArgs ...interface{}) error {
	var t asserter
	a(&t, got, expected, msgAndArgs...)
	return t.err
}

type actualAssertion func(t assert.TestingT, got, expected interface{}, msgAndArgs ...interface{}) bool

// asserter is used to be able to retrieve the error reported by the called assertion
type asserter struct {
	err error
}

// Errorf is used by the called assertion to report an error
func (a *asserter) Errorf(format string, args ...interface{}) {
	a.err = fmt.Errorf(format, args...)
}
