package uuid

import (
	"reflect"
	"testing"

	"github.com/google/uuid"
)

// 4915e120-3594-b29e-bd92-62abff23e1c6
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
