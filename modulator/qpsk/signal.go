package qpsk

import "encoding/json"

type Signal struct {
	InPhase    []byte
	Quadrature []byte
}

// Marshal converts signal to JSON representation
func (s Signal) Marshal() ([]byte, error) {
	return json.Marshal(map[string][]byte{
		"inPhase":    s.InPhase,
		"quadrature": s.Quadrature,
	})
}
