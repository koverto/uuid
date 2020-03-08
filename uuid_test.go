package uuid

import (
	fmt "fmt"
	"reflect"
	"testing"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson/bsontype"
)

const uuidString = "4915e120-3594-b29e-bd92-62abff23e1c6"

func uuidBytes() [16]byte {
	return [16]byte{0x49, 0x15, 0xe1, 0x20, 0x35, 0x94, 0xb2, 0x9e, 0xbd, 0x92, 0x62, 0xab, 0xff, 0x23, 0xe1, 0xc6}
}

func uuidZero() [16]byte {
	return [16]byte{0x0}
}

func TestNew(t *testing.T) {
	tests := []struct {
		name string
		want int
	}{
		{"New UUID", 16},
	}
	for _, tt := range tests {
		test := tt
		t.Run(test.name, func(t *testing.T) {
			u := New()
			if got := u.Uuid.Size(); !reflect.DeepEqual(got, test.want) {
				t.Errorf("New().Size() = %v, want %v", got, test.want)
			}
		})
	}
}

func TestUUID_Size(t *testing.T) {
	uuidBytes := uuidBytes()
	tests := []struct {
		name   string
		fields *_uuid
		want   int
	}{
		{"Valid UUID", &_uuid{uuid.UUID(uuidBytes)}, 16},
		{"Nil UUID", nil, 0},
	}

	for _, tt := range tests {
		test := tt
		t.Run(test.name, func(t *testing.T) {
			if got := test.fields.Size(); got != test.want {
				t.Errorf("_uuid.Size() = %v, want %v", got, test.want)
			}
		})
	}
}

func TestUUID_MarshalTo(t *testing.T) {
	uuidBytes := uuidBytes()
	tests := []struct {
		name    string
		fields  *_uuid
		args    []byte
		want    int
		wantErr bool
	}{
		{"Valid UUID", &_uuid{uuid.UUID(uuidBytes)}, make([]byte, 0), 16, false},
		{"Nil UUID", nil, make([]byte, 0), 0, false},
	}

	for _, tt := range tests {
		test := tt
		t.Run(test.name, func(t *testing.T) {
			got, err := test.fields.MarshalTo(test.args)
			if (err != nil) != test.wantErr {
				t.Errorf("_uuid.MarshalTo() error = %v, wantErr %v", err, test.wantErr)
				return
			}
			if got != test.want {
				t.Errorf("_uuid.MarshalTo() = %v, want %v", got, test.want)
			}
		})
	}
}

func TestUUID_Unmarshal(t *testing.T) {
	uuidBytes := uuidBytes()
	tests := []struct {
		name    string
		args    []byte
		want    [16]byte
		wantErr bool
	}{
		{"Valid UUID", uuidBytes[:], uuidBytes, false},
		{"Zero length data", make([]byte, 0), [16]byte{}, false},
		{"Invalid data", []byte{0x00}, [16]byte{}, true},
	}

	for _, tt := range tests {
		test := tt
		t.Run(test.name, func(t *testing.T) {
			u := &_uuid{}
			if err := u.Unmarshal(test.args); (err != nil) != test.wantErr {
				t.Errorf("_uuid.Unmarshal() error = %v, wantErr %v", err, test.wantErr)
			}
			if u.UUID != test.want {
				t.Errorf("_uuid.Unmarshal() = %v, want %v", u.UUID, test.want)
			}
		})
	}
}

func TestUUID_MarshalBSONValue(t *testing.T) {
	uuidBytes := uuidBytes()
	header := []byte{0x10, 0x00, 0x00, 0x00, 0x04}

	tests := []struct {
		name    string
		want    bsontype.Type
		wantErr bool
		want1   []byte
		fields  uuid.UUID
	}{
		{"Valid UUID", bsontype.Binary, false, append(header, uuidBytes[:]...), uuidBytes},
	}

	for _, tt := range tests {
		test := tt
		t.Run(test.name, func(t *testing.T) {
			u := _uuid{
				UUID: test.fields,
			}
			got, got1, err := u.MarshalBSONValue()
			if (err != nil) != test.wantErr {
				t.Errorf("_uuid.MarshalBSONValue() error = %v, wantErr %v", err, test.wantErr)
				return
			}
			if !reflect.DeepEqual(got, test.want) {
				t.Errorf("_uuid.MarshalBSONValue() got = %v, want %v", got, test.want)
			}
			if !reflect.DeepEqual(got1, test.want1) {
				t.Errorf("_uuid.MarshalBSONValue() got1 = %v, want %v", got1, test.want1)
			}
		})
	}
}

func TestUUID_UnmarshalBSONValue(t *testing.T) {
	uuidBytes := uuidBytes()
	goodHeader := []byte{0x10, 0x00, 0x00, 0x00, 0x04}
	badTypeHeader := []byte{0x01, 0x00, 0x00, 0x00, 0x04}
	badSubtypeHeader := []byte{0x10, 0x00, 0x00, 0x00, 0x01}

	type args struct {
		bsonType bsontype.Type
		data     []byte
	}

	tests := []struct {
		name    string
		args    args
		want    [16]byte
		wantErr bool
	}{
		{"Good header", args{bsontype.Binary, append(goodHeader, uuidBytes[:]...)}, uuidBytes, false},
		{"Bad type", args{bsontype.Boolean, append(goodHeader, uuidBytes[:]...)}, [16]byte{}, true},
		{"Bad type header", args{bsontype.Binary, append(badTypeHeader, uuidBytes[:]...)}, [16]byte{}, true},
		{"Bad subtype header", args{bsontype.Binary, append(badSubtypeHeader, uuidBytes[:]...)}, [16]byte{}, true},
	}

	for _, tt := range tests {
		test := tt
		t.Run(test.name, func(t *testing.T) {
			u := &_uuid{}
			if err := u.UnmarshalBSONValue(test.args.bsonType, test.args.data); (err != nil) != test.wantErr {
				t.Errorf("_uuid.UnmarshalBSONValue() error = %v, wantErr %v", err, test.wantErr)
			}
		})
	}
}

func TestUUID_MarshalJSON(t *testing.T) {
	uuidBytes := uuidBytes()
	tests := []struct {
		name    string
		wantErr bool
		fields  uuid.UUID
		want    []byte
	}{
		{"Valid UUID", false, uuid.UUID(uuidBytes), []byte(fmt.Sprintf("\"%s\"", uuidString))},
	}

	for _, tt := range tests {
		test := tt
		t.Run(test.name, func(t *testing.T) {
			u := UUID{&_uuid{
				UUID: test.fields,
			}}
			got, err := u.MarshalJSON()
			if (err != nil) != test.wantErr {
				t.Errorf("UUID.MarshalJSON() error = %v, wantErr %v", err, test.wantErr)
				return
			}
			if !reflect.DeepEqual(got, test.want) {
				t.Errorf("UUID.MarshalJSON() = %v, want %v", got, test.want)
			}
		})
	}
}

func TestUUID_UnmarshalJSON(t *testing.T) {
	uuidBytes := uuidBytes()
	tests := []struct {
		name    string
		fields  uuid.UUID
		args    []byte
		wantErr bool
	}{
		{"Valid UUID", uuid.UUID(uuidBytes), []byte(uuidString), false},
		{"Invalid UUID", uuid.UUID{}, []byte(""), true},
	}

	for _, tt := range tests {
		test := tt
		t.Run(test.name, func(t *testing.T) {
			u := UUID{&_uuid{
				UUID: test.fields,
			}}
			if err := u.UnmarshalJSON(test.args); (err != nil) != test.wantErr {
				t.Errorf("UUID.UnmarshalJSON() error = %v, wantErr %v", err, test.wantErr)
			}
		})
	}
}

func TestParse(t *testing.T) {
	uuidBytes := uuidBytes()
	uuidZero := uuidZero()
	tests := []struct {
		name    string
		args    string
		want    *UUID
		wantErr bool
	}{
		{"Valid UUID", uuidString, &UUID{&_uuid{uuid.UUID(uuidBytes)}}, false},
		{"Invalid UUID", "invalid", &UUID{&_uuid{uuid.UUID(uuidZero)}}, true},
	}

	for _, tt := range tests {
		test := tt
		t.Run(test.name, func(t *testing.T) {
			got, err := Parse(test.args)
			if (err != nil) != test.wantErr {
				t.Errorf("Parse() error = %v, wantErr %v", err, test.wantErr)
				return
			}
			if !reflect.DeepEqual(got, test.want) {
				t.Errorf("Parse() = %v, want %v", got, test.want)
			}
		})
	}
}
