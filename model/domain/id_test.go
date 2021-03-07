package domain

import (
	"reflect"
	"testing"

	uuid "github.com/satori/go.uuid"
)

func TestIDFromString(t *testing.T) {
	t.Parallel()
	var (
		id = uuid.NewV4()
	)
	type args struct {
		s string
	}
	tests := []struct {
		name    string
		args    args
		want    ID
		wantErr bool
	}{
		{
			name: "id_from_string_success",
			args: args{
				s: id.String(),
			},
			want: ID(id),
		},
		{
			name: "id_from_string_failed",
			args: args{
				s: "123",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := IDFromString(tt.args.s)
			if (err != nil) != tt.wantErr {
				t.Errorf("IDFromString() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("IDFromString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestID_IsZero(t *testing.T) {
	t.Parallel()
	var (
		id = NewID()
	)
	tests := []struct {
		name string
		u    *ID
		want bool
	}{
		{
			name: "id_is_zero_success_1",
			u:    nil,
			want: true,
		},
		{
			name: "id_is_zero_success_2",
			u:    (*ID)(&uuid.Nil),
			want: true,
		},
		{
			name: "id_is_zero_failed",
			u:    &id,
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.u.IsZero(); got != tt.want {
				t.Errorf("ID.IsZero() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestID_String(t *testing.T) {
	var (
		str     = "5fa03df5-ae8d-4119-95a5-0abad2413165"
		id, err = IDFromString(str)
	)

	if err != nil {
		t.Fatalf("failed get id from string by error %v", err)
	}

	tests := []struct {
		name string
		u    ID
		want string
	}{
		{
			name: "id_string_success",
			u:    id,
			want: str,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.u.String(); got != tt.want {
				t.Errorf("ID.String() = %v, want %v", got, tt.want)
			}
		})
	}
}
