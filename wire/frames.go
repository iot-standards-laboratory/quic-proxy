package wire

type headersFrame struct {
	code  uint8
	token []byte
}

type optionFrame struct {
	optionLength uint64
	optionDelta  []byte
}

type dataFrame struct {
	length  uint64
	payload []byte
}

type emptyFrame struct{}
