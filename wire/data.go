package wire

import (
	"bytes"
	"quicproxy/protocol"
)

type dataFrame struct {
	payload []byte
}

// Length of a written frame
func (f *dataFrame) Length() protocol.ByteCount {
	return protocol.ByteCount(len(f.payload))
}

// func ()

func parseDataFrame(r *bytes.Reader) (*dataFrame, error) {
	return nil, nil
}
