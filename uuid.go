//go:generate protoc --gogofaster_out=plugins=grpc:. --proto_path=$GOPATH/pkg/mod:. uuid.proto

package uuid

import (
	"github.com/google/uuid"
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
