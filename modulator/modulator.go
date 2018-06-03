package modulator

type Signal struct {
	inPhase    []byte
	quadrature []byte
}

type Modulator interface {
	Modulate() Signal
}
