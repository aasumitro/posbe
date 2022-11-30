package utils_test

import (
	"fmt"
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
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pwd, err := utils.HashPassword(tt.args.s)
			fmt.Println(pwd)
			assert.Nil(t, err)
			valid, err := utils.ComparePasswords(pwd, tt.args.s)
			assert.Nil(t, err)
			assert.Equalf(t, tt.want, valid, "Hash(%v)", tt.args.s)
		})
	}
}
