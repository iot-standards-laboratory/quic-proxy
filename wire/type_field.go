package wire

type typeField uint8

const (
	MASKFIN   uint8 = 0x01
	MASKUPRT  uint8 = 0x02
	MASKFTYPE uint8 = 0x0c
)

const (
	FRAMETYPE_HEADER uint8 = iota
	FRAMETYPE_OPTION
	FRAMETYPE_DATA
	FRAMETYPE_EMPTY
)

const (
	PROTOCOL_TCP uint8 = iota
	PROTOCOL_UDP
)

func (t typeField) frameType() uint8 {
	return uint8(t) & MASKFTYPE >> 2
}

func (t typeField) underlyingProtocol() uint8 {
	return uint8(t) & MASKUPRT >> 1
}

func (t typeField) isFin() bool {
	return uint8(t)&MASKFIN != 0
}

// make a type field for frame
func newTypeField(frameType, protocol, fin uint8) typeField {
	return typeField((frameType << 2) + (protocol << 1) + fin)
}
