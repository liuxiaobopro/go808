package test

import (
	"gitee.com/coco/go808/protocol"
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

func TestT808_0x8401_EncodeDecode(t *testing.T) {
	message := protocol.T808_0x8401{
		Type: protocol.T808_0x8401_TypeModify,
		Contacts: []protocol.T808_0x8401_Contact{
			{
				Flag:    protocol.T808_0x8401_ContactFlagBoth,
				Number:  "+8619942351698",
				Contact: "老王",
			},
			{
				Flag:    protocol.T808_0x8401_ContactFlagBoth,
				Number:  "+8619942351698",
				Contact: "老王",
			},
		},
	}
	data, err := message.Encode()
	if err != nil {
		assert.Error(t, err, "encode error")
	}

	var message2 protocol.T808_0x8401
	_, err = message2.Decode(data)
	if err != nil {
		assert.Error(t, err, "decode error")
	}
	assert.True(t, reflect.DeepEqual(message, message2))
}
