package utils_test

import (
	"github.com/aasumitro/posbe/pkg/utils"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPasswordUtils(t *testing.T) {
	type args struct {
		s string
	}

	tests := []struct {
		name    string
		args    args
		want    bool
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name:    "test hash make and verify",
			args:    args{s: "secret"},
			want:    true,
			wantErr: assert.NoError,
		},
		{
			name:    "test hash make and verify",
			args:    args{s: "12345"},
			want:    true,
			wantErr: assert.NoError,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := utils.Password{Stored: "", Supplied: tt.args.s}
			pwd, err := u.HashPassword()
			u.Stored = pwd
			assert.Nil(t, err)
			valid, err := u.ComparePasswords()
			assert.Nil(t, err)
			assert.Equalf(t, tt.want, valid, "Hash(%v)", tt.args.s)
		})
	}
}
