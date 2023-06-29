package ccoapvarint

import (
	"fmt"
	"io"
	"quicproxy/protocol"
)

const (
	// Min is the minimum value allowed for a QUIC varint.
	Min = 0

	// Max is the maximum allowed value for a QUIC varint (2^62-1).
	Max = maxVarInt8

	maxVarInt1 = 63                  // 0x3f
	maxVarInt2 = 16383               // 0x3fff
	maxVarInt4 = 1073741823          // 0x3ffffffff
	maxVarInt8 = 4611686018427387903 // 0x3fffffffffffffff
)

// Read reads a number in the QUIC varint format from r.
func Read(r io.ByteReader) (uint64, error) {
	firstByte, err := r.ReadByte()
	if err != nil {
		return 0, err
	}
	// the first two bits of the first byte encode the length
	len := 1 << ((firstByte & 0xc0) >> 6)
	b := uint64(firstByte & (0xff - 0xc0))

	for i := 0; i < len-1; i++ {
		rb, err := r.ReadByte()
		if err != nil {
			return 0, err
		}
		b <<= 8
		b += uint64(rb)
	}
	return b, nil
}

// Append appends i in the QUIC varint format.
func Append(b []byte, i uint64) []byte {
	if i <= maxVarInt1 {
		return append(b, uint8(i))
	}
	if i <= maxVarInt2 {
		return append(b, []byte{uint8(i>>8) | 0x40, uint8(i)}...)
	}
	if i <= maxVarInt4 {
		return append(b, []byte{uint8(i>>24) | 0x80, uint8(i >> 16), uint8(i >> 8), uint8(i)}...)
	}
	if i <= maxVarInt8 {
		return append(b, []byte{
			uint8(i>>56) | 0xc0, uint8(i >> 48), uint8(i >> 40), uint8(i >> 32),
			uint8(i >> 24), uint8(i >> 16), uint8(i >> 8), uint8(i),
		}...)
	}
	panic(fmt.Sprintf("%#x doesn't fit into 62 bits", i))
}

// AppendWithLen append i in the QUIC varint format with the desired length.
func AppendWithLen(b []byte, i uint64, length protocol.ByteCount) []byte {
	if length != 1 && length != 2 && length != 4 && length != 8 {
		panic("invalid varint length")
	}
	l := Len(i)
	if l == length {
		return Append(b, i)
	}
	if l > length {
		panic(fmt.Sprintf("cannot encode %d in %d bytes", i, length))
	}
	if length == 2 {
		b = append(b, 0b01000000)
	} else if length == 4 {
		b = append(b, 0b10000000)
	} else if length == 8 {
		b = append(b, 0b11000000)
	}
	for j := protocol.ByteCount(1); j < length-l; j++ {
		b = append(b, 0)
	}
	for j := protocol.ByteCount(0); j < l; j++ {
		b = append(b, uint8(i>>(8*(l-1-j))))
	}
	return b
}

// Len determines the number of bytes that will be needed to write the number i.
func Len(i uint64) protocol.ByteCount {
	if i <= maxVarInt1 {
		return 1
	}
	if i <= maxVarInt2 {
		return 2
	}
	if i <= maxVarInt4 {
		return 4
	}
	if i <= maxVarInt8 {
		return 8
	}
	// Don't use a fmt.Sprintf here to format the error message.
	// The function would then exceed the inlining budget.
	panic(struct {
		message string
		num     uint64
	}{"value doesn't fit into 62 bits: ", i})
}
