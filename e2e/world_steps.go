package e2e

type stepper interface {
	Step(match interface{}, function interface{})
}

func (w *World) RegisterSteps(sc stepper) {
	// Given
	//	sc.Step(`^there\'s a network$`, w.theresANetwork)

	// When

	// Then
}
