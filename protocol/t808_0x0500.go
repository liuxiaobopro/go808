package protocol

// 车辆控制
type T808_0x0500 struct {
	AnswerMessageSerialNo uint16
	Result                T808_0x0200
}

func (entity *T808_0x0500) MsgID() MsgID {
	return MsgT808_0x0500
}

func (entity *T808_0x0500) Encode() ([]byte, error) {
	writer := NewWriter()

	// 写入消息序列号
	writer.WriteUint16(entity.AnswerMessageSerialNo)

	// 写入定位信息
	data, err := entity.Result.Encode()
	if err != nil {
		return nil, err
	}
	writer.Write(data)
	return writer.Bytes(), nil
}

func (entity *T808_0x0500) Decode(data []byte) (int, error) {
	if len(data) <= 3 {
		return 0, ErrInvalidBody
	}
	reader := NewReader(data)

	// 读取消息序列号
	responseMessageSerialNo, err := reader.ReadUint16()
	if err != nil {
		return 0, err
	}

	// 读取位置信息
	var result T808_0x0200
	size, err := result.Decode(data[len(data)-reader.Len():])
	if err != nil {
		return 0, err
	}

	// 更新Entity信息
	entity.Result = result
	entity.AnswerMessageSerialNo = responseMessageSerialNo
	return len(data) - reader.Len() + size, nil
}
