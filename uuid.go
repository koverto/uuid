//go:generate protoc --gogofaster_out=plugins=grpc:. --proto_path=$GOPATH/pkg/mod:. uuid.proto

// Package uuid wraps github.com/google/uuid for use as a protobuf type and with
// implementations of various de/serialization interfaces.
package uuid

import (
	fmt "fmt"
	io "io"
	"strconv"

	"github.com/google/uuid"
	"github.com/prometheus/common/log"
	"go.mongodb.org/mongo-driver/bson/bsontype"
	"go.mongodb.org/mongo-driver/x/bsonx"
)

// Parse parses a UUID from a valid string representation.
func Parse(s string) (*UUID, error) {
	parsed, err := uuid.Parse(s)
	return &UUID{&_uuid{parsed}}, err
}

// New generates a new, random UUID.
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

	return len(u.UUID)
}

// MarshalTo is required to implement the proto.Marshaler interface.
func (u *_uuid) MarshalTo(data []byte) (int, error) {
	if u == nil {
		return 0, nil
	}

	copy(data, u.UUID[:])

	return len(u.UUID), nil
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
	val := bsonx.Binary(bsontype.BinaryUUID, u.UUID[:])
	return val.MarshalBSONValue()
}

// UnmarshalBSONValue implements the bson.ValueUnmarshaler interface.
func (u *_uuid) UnmarshalBSONValue(bsonType bsontype.Type, data []byte) error {
	if bsonType != bsontype.Binary || data[0] != 0x10 || data[4] != bsontype.BinaryUUID {
		return fmt.Errorf("could not unmarshal %v as a UUID", bsonType)
	}

	return u.Unmarshal(data[5:])
}

// UnmarshalGQL implements the graphql.Unmarshal interface.
func (u *UUID) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("value for unmarshalling was not a string: %v", v)
	}

	return u.UnmarshalJSON([]byte(str))
}

// MarshalGQL implements the graphql.Marshal interface.
func (u UUID) MarshalGQL(w io.Writer) {
	marshaled, _ := u.MarshalJSON()
	if _, err := w.Write(marshaled); err != nil {
		log.Errorf("Error marshalling %v to GraphQL: %s", u, err)
	}
}

// MarshalJSON implements the json.Marshaler interface.
func (u UUID) MarshalJSON() ([]byte, error) {
	return []byte(strconv.Quote(u.Uuid.String())), nil
}

// UnmarshalJSON implements the json.Unmarshaler interface.
func (u *UUID) UnmarshalJSON(data []byte) error {
	if parsed, err := uuid.Parse(string(data)); err == nil {
		u.Uuid = &_uuid{parsed}
	} else {
		return err
	}

	return nil
}
