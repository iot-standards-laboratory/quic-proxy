// test code for data frame
package wire

import (
	"bytes"
	"quicproxy/ccoapvarint"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseDataFrame(t *testing.T) {
	b := make([]byte, 0, 1000)

	b = ccoapvarint.Append(b, uint64(len("data frame parse test")))
	buf := bytes.NewBuffer(b)
	writer := ccoapvarint.NewWriter(buf)

	w := ccoapvarint.NewWriter(writer)
	n, err := w.Write([]byte("data frame parse test"))
	assert.NoError(t, err)
	assert.Equal(t, n, len("data frame parse test"))

	df, err := parseDataFrame(ccoapvarint.NewReader(buf))

	assert.NoError(t, err)
	assert.Equal(t, string(df.payload), "data frame parse test")
}
