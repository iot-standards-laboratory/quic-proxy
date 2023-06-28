package wire

type optionsFrame struct {
	optionLength uint64
	optionDelta  []byte
}
