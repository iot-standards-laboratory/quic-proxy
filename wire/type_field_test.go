package wire

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTypeField(t *testing.T) {
	assert := assert.New(t)

	typeValue := newTypeField(FRAMETYPE_DATA, PROTOCOL_UDP, 1)
	assert.Equal(typeField(0xb), typeValue)

	assert.Equal(FRAMETYPE_DATA, typeValue.frameType())
	assert.NotEqual(FRAMETYPE_HEADER, typeValue.frameType())
	assert.Equal(PROTOCOL_UDP, typeValue.underlyingProtocol())
	assert.True(typeValue.isFin())
}
