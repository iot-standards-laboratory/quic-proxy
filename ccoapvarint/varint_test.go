package ccoapvarint_test

import (
	"bytes"
	"quicproxy/ccoapvarint"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAppendWithLen(t *testing.T) {
	assert := assert.New(t)

	// b := make([]byte, 0, 1)
	b := ccoapvarint.Append(nil, 13)
	l, err := ccoapvarint.Read(bytes.NewReader(b))

	assert.NoError(err)
	assert.Equal(l, uint64(13))
	assert.Equal(len(b), 1)

	b = ccoapvarint.Append(nil, 1073741824)
	l, err = ccoapvarint.Read(bytes.NewReader(b))

	assert.NoError(err)
	assert.Equal(l, uint64(1073741824))
	assert.Equal(len(b), 8)
}
