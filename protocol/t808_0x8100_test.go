package protocol

import (
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

func TestT808_0x8100_EncodeDecode(t *testing.T) {
	message := T808_0x8100{
		MessageSerialNo: 1234,
		Result:          T808_0x8100_ResultTerminalRegistered,
		AuthKey:         "鉴权码",
	}
	data, err := message.Encode()
	if err != nil {
		assert.Error(t, err, "encode error")
	}

	var message2 T808_0x8100
	_, err = message2.Decode(data)
	if err != nil {
		assert.Error(t, err, "decode error")
	}
	assert.True(t, reflect.DeepEqual(message, message2))
}