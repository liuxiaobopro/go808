package protocol

import (
	"go808/errors"
)

// 多媒体数据上传应答
type T808_0x8800 struct {
	MediaID  uint32
	RetryIDs []uint16
}

func (entity *T808_0x8800) MsgID() MsgID {
	return MsgT808_0x8800
}

func (entity *T808_0x8800) Encode() ([]byte, error) {
	writer := NewWriter()

	// 写入媒体ID
	writer.WriteUint32(entity.MediaID)

	if len(entity.RetryIDs) == 0 {
		return writer.Bytes(), nil
	}

	// 写入重传消息数
	writer.WriteByte(byte(len(entity.RetryIDs)))

	// 写入重传消息ID
	for _, id := range entity.RetryIDs {
		writer.WriteUint16(id)
	}
	return writer.Bytes(), nil
}

func (entity *T808_0x8800) Decode(data []byte) (int, error) {
	if len(data) < 4 {
		return 0, errors.ErrEntityDecodeFail
	}
	reader := NewReader(data)

	// 读取媒体ID
	var err error
	entity.MediaID, err = reader.ReadUint32()
	if err != nil {
		return 0, errors.ErrEntityDecodeFail
	}

	if reader.Len() == 0 {
		return len(data) - reader.Len(), nil
	}

	// 读取重传消息数
	count, err := reader.ReadByte()
	if err != nil {
		return 0, errors.ErrEntityDecodeFail
	}

	// 读取重传消息ID
	for i := 0; i < int(count); i++ {
		id, err := reader.ReadUint16()
		if err != nil {
			return 0, errors.ErrEntityDecodeFail
		}
		entity.RetryIDs = append(entity.RetryIDs, id)
	}
	return len(data) - reader.Len(), nil
}