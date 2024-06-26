package test

import (
	"gitee.com/coco/go808/protocol"
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
	"time"
)

func TestT808_0x8803_EncodeDecode(t *testing.T) {
	message := protocol.T808_0x8803{
		Type:       protocol.T808_0x0800_MediaTypeAudio,
		ChannelID:  56,
		Event:      87,
		StartTime:  time.Unix(time.Now().Unix(), 0),
		EndTime:    time.Unix(time.Now().Unix(), 0),
		RemoveFlag: 1,
	}
	data, err := message.Encode()
	if err != nil {
		assert.Error(t, err, "encode error")
	}

	var message2 protocol.T808_0x8803
	_, err = message2.Decode(data)
	if err != nil {
		assert.Error(t, err, "decode error")
	}
	assert.True(t, reflect.DeepEqual(message, message2))
}
