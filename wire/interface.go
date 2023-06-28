package wire

import (
	"io"
)

type SubframeType uint8

type FrameI interface {
	Type() SubframeType
	write(b io.Writer) error
	SetFin()
	IsFin() bool
	UnderlyingProtocol() uint8
}

// func (f *_CCoAPFrame) Type() SubframeType {
// 	return f.subtype & MASKFTYPE
// }

// func (f *_CCoAPFrame) SetFin() {
// 	f.subtype = f.subtype | MASKFIN
// }

// func (f *_CCoAPFrame) GetUnderlyingProtocol() uint8 {
// 	return uint8(f.subtype & MASKUPRT)
// }

// func (f *_CCoAPFrame) IsFin() bool {
// 	return f.subtype&MASKFIN != 0
// }

// func (f *_CCoAPFrame) write(w io.Writer) error {
// 	b := ccoapvarint.NewWriter(w)
// 	err := ccoapvarint.WriteBytes(b, uint64(f.seq), 1)
// 	if err != nil {
// 		return err
// 	}

// 	err = ccoapvarint.WriteBytes(b, uint64(f.subtype), 1)
// 	if err != nil {
// 		return err
// 	}

// 	err = ccoapvarint.WriteBytes(b, uint64(f.length), 2)
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }

// func ParseFrame(r io.Reader) (CCoAPFrameI, error) {
// 	// codeByte, err := r.ReadByte()
// 	// if err != nil {
// 	// 	return nil, err
// 	// }
// 	reader := ccoapvarint.NewReader(r)

// 	intFrame, err := reader.ReadBytesAsInt64(4)
// 	if err != nil {
// 		return nil, err
// 	}
// 	length := uint16(intFrame)
// 	intFrame >>= 16

// 	subtype := SubframeType(intFrame)
// 	intFrame >>= 8
// 	seq := uint8(intFrame)

// 	// switch subtype & MASKFTYPE {
// 	// case HEADERS:
// 	// 	remainBytes, err := reader.ReadBytes(uint64(length) - 4)
// 	// 	if err != nil {
// 	// 		return nil, err
// 	// 	}

// 	// 	frame, err := parseHeadersframe(remainBytes)
// 	// 	if err != nil {
// 	// 		return nil, err
// 	// 	}
// 	// 	frame._CCoAPFrame = &_CCoAPFrame{length, subtype, seq}
// 	// 	return frame, nil
// 	// case OPTIONS:
// 	// 	remainBytes, err := reader.ReadBytes(uint64(length) - 4)
// 	// 	if err != nil {
// 	// 		return nil, err
// 	// 	}

// 	// 	frame, err := parseOptionsFrame(remainBytes)
// 	// 	if err != nil {
// 	// 		return nil, err
// 	// 	}
// 	// 	frame._CCoAPFrame = &_CCoAPFrame{length, subtype, seq}
// 	// 	return frame, nil
// 	// case DATA:
// 	// 	frame, err := parseDataframe(r, length)
// 	// 	if err != nil {
// 	// 		return nil, err
// 	// 	}
// 	// 	frame._CCoAPFrame = &_CCoAPFrame{length, subtype, seq}
// 	// 	return frame, nil
// 	// default:
// 	// 	return nil, fmt.Errorf("wrong subtype: %v", subtype)
// 	// }

// }
