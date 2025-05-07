package localtime

import (
	"errors"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/bsontype"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/x/bsonx/bsoncore"
	"time"
)

// MongoTime is alias type for time.Time
// mysql 时间格式化 和 gorm的Scan/Value的实现类
type MongoTime time.Time

// UnmarshalJSON implements json unmarshal interface.
func (t *MongoTime) UnmarshalJSON(data []byte) (err error) {
	now, err := time.ParseInLocation(`"`+time.DateTime+`"`, string(data), time.Local)
	*t = MongoTime(now)
	return
}

// MarshalJSON implements json marshal interface.
func (t MongoTime) MarshalJSON() ([]byte, error) {
	b := make([]byte, 0, len(time.DateTime)+2)
	if time.Time(t).IsZero() {
		b = append(b, '"', '"')
		return b, nil
	}
	b = append(b, '"')
	b = time.Time(t).AppendFormat(b, time.DateTime)
	b = append(b, '"')
	return b, nil
}

// MarshalBSONValue mongodb是存储bson格式，因此需要实现序列化bsonvalue(这里不能实现MarshalBSON，MarshalBSON是处理Document的)，将时间转换成mongodb能识别的primitive.DateTime
func (t *MongoTime) MarshalBSONValue() (bsontype.Type, []byte, error) {
	targetTime := primitive.NewDateTimeFromTime(time.Time(*t))
	return bson.MarshalValue(targetTime)
}

// UnmarshalBSONValue 实现bson反序列化，从mongodb中读取数据转换成time.Time格式，这里用到了bsoncore中的方法读取数据转换成datetime然后再转换成time.MongoTime
func (t *MongoTime) UnmarshalBSONValue(t2 bsontype.Type, data []byte) error {
	v, _, valid := bsoncore.ReadValue(data, t2)
	if valid == false {
		return errors.New(fmt.Sprintf("%s, %s, %s", "读取数据失败:", t2, data))
	}
	if v.Type == bsontype.DateTime {
		*t = MongoTime(v.Time())
	}
	return nil
}
