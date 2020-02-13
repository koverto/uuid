//go:generate protoc --gogofaster_out=plugins=grpc:. --proto_path=$GOPATH/pkg/mod:. uuid.proto

package uuid

import (
	fmt "fmt"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson/bsontype"
	"go.mongodb.org/mongo-driver/x/bsonx"
)

func New() *UUID {
	id := _uuid{uuid.New()}
	return &UUID{&id}
}

type _uuid struct {
	uuid.UUID
}

// Size is required to implement the proto.Marshaler interface.
func (u *_uuid) Size() int {
	if u == nil {
		return 0
	}
	return 16
}

// MarshalTo is required to implement the proto.Marshaler interface.
func (u *_uuid) MarshalTo(data []byte) (int, error) {
	if u == nil {
		return 0, nil
	}
	copy(data, u.UUID[:])
	return 16, nil
}

// Unmarshal is required to implement the proto.Marshaler interface.
func (u *_uuid) Unmarshal(data []byte) error {
	if len(data) == 0 {
		u = nil
		return nil
	}

	uid, err := uuid.FromBytes(data)
	if err == nil {
		u.UUID = uid
	}

	return err
}

// MarshalBSONValue implements the bson.ValueMarshaler interface.
func (u _uuid) MarshalBSONValue() (bsontype.Type, []byte, error) {
	val := bsonx.Binary(0x04, u.UUID[:])
	return val.MarshalBSONValue()
}

// UnmarshalBSONValue implements the bson.ValueUnmarshaler interface.
func (u *_uuid) UnmarshalBSONValue(bsonType bsontype.Type, data []byte) error {
	if bsonType != bsontype.Binary || data[0] != 0x10 || data[4] != 0x04 {
		return fmt.Errorf("Could not unmarshal %v as a UUID", bsonType)
	}

	return u.Unmarshal(data[5:])
}
