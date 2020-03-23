package protocol

import (
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

func TestT808_0x8001_EncodeDecode(t *testing.T) {
	message := T808_0x8001{
		MessageSerialNo: 1234,
		Result:          T808_0x8100_ResultTerminalRegistered,
		MsgID:           TypeT808_0x8103,
	}
	data, err := message.Encode()
	if err != nil {
		assert.Error(t, err, "encode error")
	}

	var message2 T808_0x8001
	_, err = message2.Decode(data)
	if err != nil {
		assert.Error(t, err, "decode error")
	}
	assert.True(t, reflect.DeepEqual(message, message2))
}
