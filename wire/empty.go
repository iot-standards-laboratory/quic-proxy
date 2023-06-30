package wire

import "quicproxy/ccoapvarint"

type emptyFrame struct {
	code uint64
}

func parseEmptyFrame(r ccoapvarint.Reader) (*emptyFrame, error) {
	code, err := ccoapvarint.Read(r)
	if err != nil {
		return nil, err
	}

	return &emptyFrame{code: code}, nil
}
