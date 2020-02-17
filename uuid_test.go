package uuid

import (
	fmt "fmt"
	"reflect"
	"testing"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson/bsontype"
)

var uuidString = "4915e120-3594-b29e-bd92-62abff23e1c6"
var uuidBytes = [16]byte{0x49, 0x15, 0xe1, 0x20, 0x35, 0x94, 0xb2, 0x9e, 0xbd, 0x92, 0x62, 0xab, 0xff, 0x23, 0xe1, 0xc6}

func TestNew(t *testing.T) {
	tests := []struct {
		name string
		want int
	}{
		{"New UUID", 16},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := New()
			if got := u.Uuid.Size(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New().Size() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUUID_Size(t *testing.T) {
	type fields struct {
		UUID *_uuid
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{"Valid UUID", fields{&_uuid{uuid.UUID(uuidBytes)}}, 16},
		{"Nil UUID", fields{nil}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := tt.fields.UUID
			if got := u.Size(); got != tt.want {
				t.Errorf("_uuid.Size() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUUID_MarshalTo(t *testing.T) {
	type fields struct {
		UUID *_uuid
	}
	type args struct {
		data []byte
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    int
		wantErr bool
	}{
		{"Valid UUID", fields{&_uuid{uuid.UUID(uuidBytes)}}, args{make([]byte, 0)}, 16, false},
		{"Nil UUID", fields{nil}, args{make([]byte, 0)}, 0, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := tt.fields.UUID
			got, err := u.MarshalTo(tt.args.data)
			if (err != nil) != tt.wantErr {
				t.Errorf("_uuid.MarshalTo() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("_uuid.MarshalTo() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUUID_Unmarshal(t *testing.T) {
	type args struct {
		data []byte
	}
	tests := []struct {
		name    string
		args    args
		want    [16]byte
		wantErr bool
	}{
		{"Valid UUID", args{uuidBytes[:]}, uuidBytes, false},
		{"Zero length data", args{make([]byte, 0)}, [16]byte{}, false},
		{"Invalid data", args{[]byte{0x00}}, [16]byte{}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &_uuid{}
			if err := u.Unmarshal(tt.args.data); (err != nil) != tt.wantErr {
				t.Errorf("_uuid.Unmarshal() error = %v, wantErr %v", err, tt.wantErr)
			}
			if u.UUID != tt.want {
				t.Errorf("_uuid.Unmarshal() = %v, want %v", u.UUID, tt.want)
			}
		})
	}
}

func TestUUID_MarshalBSONValue(t *testing.T) {
	header := []byte{0x10, 0x00, 0x00, 0x00, 0x04}

	type fields struct {
		UUID uuid.UUID
	}
	tests := []struct {
		name    string
		fields  fields
		want    bsontype.Type
		want1   []byte
		wantErr bool
	}{
		{"Valid UUID", fields{uuidBytes}, bsontype.Binary, append(header, uuidBytes[:]...), false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := _uuid{
				UUID: tt.fields.UUID,
			}
			got, got1, err := u.MarshalBSONValue()
			if (err != nil) != tt.wantErr {
				t.Errorf("_uuid.MarshalBSONValue() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("_uuid.MarshalBSONValue() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("_uuid.MarshalBSONValue() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestUUID_UnmarshalBSONValue(t *testing.T) {
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
		t.Run(tt.name, func(t *testing.T) {
			u := &_uuid{}
			if err := u.UnmarshalBSONValue(tt.args.bsonType, tt.args.data); (err != nil) != tt.wantErr {
				t.Errorf("_uuid.UnmarshalBSONValue() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestUUID_MarshalJSON(t *testing.T) {
	type fields struct {
		UUID uuid.UUID
	}
	tests := []struct {
		name    string
		fields  fields
		want    []byte
		wantErr bool
	}{
		{"Valid UUID", fields{uuid.UUID(uuidBytes)}, []byte(fmt.Sprintf("\"%s\"", uuidString)), false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := _uuid{
				UUID: tt.fields.UUID,
			}
			got, err := u.MarshalJSON()
			if (err != nil) != tt.wantErr {
				t.Errorf("UUID.MarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UUID.MarshalJSON() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUUID_UnmarshalJSON(t *testing.T) {
	type fields struct {
		UUID uuid.UUID
	}
	type args struct {
		data []byte
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{"Valid UUID", fields{uuid.UUID(uuidBytes)}, args{[]byte(uuidString)}, false},
		{"Invalid UUID", fields{}, args{[]byte("")}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &_uuid{
				UUID: tt.fields.UUID,
			}
			if err := u.UnmarshalJSON(tt.args.data); (err != nil) != tt.wantErr {
				t.Errorf("UUID.UnmarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
