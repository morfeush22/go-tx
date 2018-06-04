package modulator

type Signal struct {
	InPhase    []byte
	Quadrature []byte
}

type Modulator interface {
	Modulate() Signal
}
