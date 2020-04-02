package protocol

import (
	"bytes"
	"go808/errors"
	"io"
	"io/ioutil"
)

// 多媒体类型
type MediaType byte

var (
	MediaTypeImage MediaType = 0
	MediaTypeAudio MediaType = 1
	MediaTypeVideo MediaType = 2
)

// 多媒体编码
type MediaCoding byte

var (
	MediaCodingJPEG MediaCoding = 0
	MediaCodingTIF  MediaCoding = 1
	MediaCodingMP3  MediaCoding = 2
	MediaCodingWAV  MediaCoding = 3
	MediaCodingWMV  MediaCoding = 4
)

// 多媒体数据上传
type T808_0x0801 struct {
	MediaID   uint32
	Type      MediaType
	Coding    MediaCoding
	Event     byte
	ChannelID byte
	Location  T808_0x0200
	Packet    io.Reader
}

func (entity *T808_0x0801) MsgID() MsgID {
	return MsgT808_0x0801
}

func (entity *T808_0x0801) Encode() ([]byte, error) {
	writer := NewWriter()

	// 写入媒体ID
	writer.WriteUint32(entity.MediaID)

	// 写入媒体类型
	writer.WriteByte(byte(entity.Type))

	// 写入媒体编码
	writer.WriteByte(byte(entity.Coding))

	// 写入事件类型
	writer.WriteByte(entity.Event)

	// 写入通道ID
	writer.WriteByte(entity.ChannelID)

	// 写入定位信息
	entity.Location.Extras = nil
	data, err := entity.Location.Encode()
	if err != nil {
		return nil, err
	}
	writer.WriteBytes(data)

	// 写入数据包
	if entity.Packet != nil {
		data, err = ioutil.ReadAll(entity.Packet)
		if err != nil {
			return nil, err
		}
		writer.WriteBytes(data)
	}
	return writer.Bytes(), nil
}

func (entity *T808_0x0801) Decode(data []byte) (int, error) {
	if len(data) < 36 {
		return 0, errors.ErrEntityDecodeFail
	}
	reader := NewReader(data)

	// 读取媒体ID
	var err error
	entity.MediaID, err = reader.ReadUint32()
	if err != nil {
		return 0, errors.ErrEntityDecodeFail
	}

	// 读取媒体类型
	mediaType, err := reader.ReadByte()
	if err != nil {
		return 0, errors.ErrEntityDecodeFail
	}
	entity.Type = MediaType(mediaType)

	// 读取媒体编码
	coding, err := reader.ReadByte()
	if err != nil {
		return 0, errors.ErrEntityDecodeFail
	}
	entity.Coding = MediaCoding(coding)

	// 读取事件类型
	entity.Event, err = reader.ReadByte()
	if err != nil {
		return 0, errors.ErrEntityDecodeFail
	}

	// 读取通道ID
	entity.ChannelID, err = reader.ReadByte()
	if err != nil {
		return 0, errors.ErrEntityDecodeFail
	}

	// 读取定位信息
	buf, err := reader.ReadBytes(28)
	if err != nil {
		return 0, errors.ErrEntityDecodeFail
	}
	if _, err = entity.Location.Decode(buf); err != nil {
		return 0, errors.ErrEntityDecodeFail
	}
	return len(data) - reader.Len(), nil
}

func (entity *T808_0x0801) GetTag() uint32 {
	return entity.MediaID
}

func (entity *T808_0x0801) GetReader() io.Reader {
	return entity.Packet
}

func (entity *T808_0x0801) SetReader(reader io.Reader) {
	entity.Packet = reader
}

func (entity *T808_0x0801) DecodePacket(data []byte) error {
	n, err := entity.Decode(data)
	if err != nil {
		return err
	}
	entity.Packet = bytes.NewReader(data[n:])
	return nil
}