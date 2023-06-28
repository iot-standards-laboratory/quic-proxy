package wire

type dataFrame struct {
	length  uint64
	payload []byte
}
