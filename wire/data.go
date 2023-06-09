package wire

import (
	"io"
	"quicproxy/ccoapvarint"
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

func parseDataFrame(r ccoapvarint.Reader) (*dataFrame, error) {
	l, err := ccoapvarint.Read(r)
	if err != nil {
		return nil, err
	}

	br, err := r.ReadBytes(l)
	if err != nil {
		return nil, err
	}

	payload, err := io.ReadAll(br)
	if err != nil {
		return nil, err
	}
	return &dataFrame{payload: payload}, nil
}
