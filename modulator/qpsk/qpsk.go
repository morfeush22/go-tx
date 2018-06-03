package qpsk

import "github.com/morfeush22/go-tx/modulator"

type QPSKModulator struct {
}

func GenerateLookupTable() (lookupTable [256]byte) {
	for i := range lookupTable {
		inPhase := i&0x80 | (i&0x20)<<1 | (i&0x8)<<2 | (i&0x2)<<3
		quadrature := (i&0x40)>>3 | (i&0x10)>>2 | (i&0x04)>>1 | i&0x01
		lookupTable[i] = byte(inPhase) | byte(quadrature)
	}
	return
}

func (m QPSKModulator) Modulate(signal []byte) modulator.Signal {
	return modulator.Signal{}
}
