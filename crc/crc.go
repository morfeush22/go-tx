package crc

import (
	"fmt"
	"math"
)

// PolyT is type of polynomial generator
type PolyT uint32

const Byte = 8

var TableSize = int(math.Pow(2, Byte))

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
func GenerateCRC(message []byte, lookupTable []PolyT, poly PolyT, polyInit PolyT, polyFinal PolyT) PolyT {
	var register = polyInit
	for _, messageByte := range message {
		fmt.Println(byte(register) ^ messageByte)
		fmt.Println(TableSize)
		register = (register >> Byte) ^ lookupTable[byte(register)^messageByte]
	}
	return register ^ polyFinal
}
