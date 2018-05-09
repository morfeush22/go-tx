package crc

import (
	"encoding/binary"
	"math"
)

// PolyT is type of polynomial generator
type PolyT uint32

const Byte = 8

var TableSize = int(math.Pow(2, Byte))

// ToByte converts poly to byte representation
func (poly PolyT) ToByte() []byte {
	buff := make([]byte, 4)
	binary.LittleEndian.PutUint32(buff, uint32(poly))
	return buff
}

// GenerateCRCLookupTable generates common CRC lookup table
func GenerateCRCLookupTable(poly PolyT) (lookupTable []PolyT) {

	lookupTable = make([]PolyT, TableSize)

	for i := 0; i < TableSize; i++ {
		var register = PolyT(i)
		for j := 0; j < Byte; j++ {
			if (register & 0x1) == 1 {
				register = (register >> 1) ^ poly
			} else {
				register = register >> 1
			}
		}
		lookupTable[i] = register
	}

	return
}

//GenerateCRC generates CRC for message
func GenerateCRC(message []byte, lookupTable []PolyT, polyInit PolyT, polyFinal PolyT) PolyT {
	var register = polyInit
	for _, messageByte := range message {
		register = (register >> Byte) ^ lookupTable[byte(register)^messageByte]
	}
	return register ^ polyFinal
}
