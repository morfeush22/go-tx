package crc

// PolyT is type of polynomial generator
type PolyT uint32

const Byte = 8
const TableSize = 2 ^ Byte

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
