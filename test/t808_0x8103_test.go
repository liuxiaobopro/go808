package test

import (
	"gitee.com/coco/go808/protocol"
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

func TestT808_0x8103_EncodeDecode(t *testing.T) {
	message := protocol.T808_0x8103{
		Params: []*protocol.Param{
			new(protocol.Param).SetByte(0x0084, 24),
			new(protocol.Param).SetBytes(0x0110, []byte{1, 2, 3, 4, 5, 6, 7, 8}),
			new(protocol.Param).SetUint16(0x0031, 100),
			new(protocol.Param).SetUint32(0x0046, 64000),
			new(protocol.Param).SetString(0x0083, "车牌号码"),
		},
	}
	data, err := message.Encode()
	if err != nil {
		assert.Error(t, err, "encode error")
	}

	var message2 protocol.T808_0x8103
	_, err = message2.Decode(data)
	if err != nil {
		assert.Error(t, err, "decode error")
	}
	assert.True(t, reflect.DeepEqual(message, message2))
}
