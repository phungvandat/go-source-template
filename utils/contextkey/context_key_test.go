package contextkey

import (
	"context"
	"reflect"
	"testing"
)

func TestGetUserID(t *testing.T) {
	t.Parallel()
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "context_key_get_user_id_success",
			args: args{
				ctx: context.WithValue(context.TODO(), CtxKeyUserID, "123"),
			},
			want: "123",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetUserID(tt.args.ctx); got != tt.want {
				t.Errorf("GetUserID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSetUserID(t *testing.T) {
	t.Parallel()
	type args struct {
		ctx    context.Context
		userID string
	}
	tests := []struct {
		name string
		args args
		want context.Context
	}{
		{
			name: "context_key_set_user_id_success",
			args: args{
				ctx:    context.TODO(),
				userID: "123",
			},
			want: context.WithValue(context.TODO(), CtxKeyUserID, "123"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SetUserID(tt.args.ctx, tt.args.userID); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetUserID() = %v, want %v", got, tt.want)
			}
		})
	}
}
