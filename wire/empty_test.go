package wire

import (
	"bytes"
	"quicproxy/ccoapvarint"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseEmptyFrame(t *testing.T) {
	b := make([]byte, 0, 1000)

	b = ccoapvarint.Append(b, uint64(100))
	assert.Equal(t, 2, len(b))

	ef, err := parseEmptyFrame(ccoapvarint.NewReader(bytes.NewReader(b)))
	assert.NoError(t, err)
	assert.Equal(t, ef.code, uint64(100))
}
