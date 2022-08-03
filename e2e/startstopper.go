package e2e

type StartStopper interface {
	Start() error
	Stop() error
}
