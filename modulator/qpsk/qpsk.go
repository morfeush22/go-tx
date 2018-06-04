package qpsk

import (
	"github.com/morfeush22/go-tx/modulator"
	"math"
)

type Modulator struct {
}

const (
	inPhaseMask    = 0xf0
	quadratureMask = 0x0f
)

var lookupTable = GenerateLookupTable()

func GenerateLookupTable() (lookupTable [256]byte) {
	for i := range lookupTable {
		inPhase := i&0x80 | (i&0x20)<<1 | (i&0x8)<<2 | (i&0x2)<<3
		quadrature := (i&0x40)>>3 | (i&0x10)>>2 | (i&0x04)>>1 | i&0x01
		lookupTable[i] = byte(inPhase) | byte(quadrature)
	}
	return
}

func (m Modulator) Modulate(inSignal []byte) modulator.Signal {
	iqLen := int(math.Ceil(float64(len(inSignal)) / 2))
	outSignal := modulator.Signal{InPhase: make([]byte, iqLen), Quadrature: make([]byte, iqLen)}

	for i, b := range inSignal {
		iqIndex := i / 2
		iq := lookupTable[b]
		outSignal.InPhase[iqIndex] = (outSignal.InPhase[iqIndex] << 4) | (iq&inPhaseMask)>>4
		outSignal.Quadrature[iqIndex] = (outSignal.Quadrature[iqIndex] << 4) | (iq & quadratureMask)
	}

	return outSignal
}
