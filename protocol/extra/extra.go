package extra

import (
	"fmt"
	"gitee.com/coco/go808/errors"
	log "github.com/sirupsen/logrus"
)

// 附加信息
type Entity interface {
	ID() byte
	Data() []byte
	Value() interface{}
	Decode(data []byte) (int, error)
}

func Decode(typ byte, data []byte) (Entity, int, error) {
	creator, ok := entityMapper[typ]
	if !ok {
		return nil, 0, errors.ErrTypeNotRegistered
	}

	entity := creator()
	count, err := entity.Decode(data)
	if err != nil {
		log.WithFields(log.Fields{
			"id":     fmt.Sprintf("0x%x", typ),
			"reason": err,
		}).Error("[JT/T808] decode extra type error")
		return nil, 0, err
	}
	return entity, count, nil
}
