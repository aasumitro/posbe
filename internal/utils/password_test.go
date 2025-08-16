package utils_test

import (
	"crypto/rand"
	"errors"
	"testing"

	"github.com/aasumitro/posbe/internal/utils"
	"github.com/stretchr/testify/assert"
)

// Mock rand.Reader that always returns an error
type mockRandReader struct {
	err error
}

func (r *mockRandReader) Read(_ []byte) (n int, err error) {
	return 0, r.err
}

func TestHash_Must_Success(t *testing.T) {
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
			pwd, err := utils.MakePassword(utils.Parallelization, tt.args.s)
			assert.Nil(t, err)
			valid, err := utils.ComparePassword(utils.Parallelization, pwd, tt.args.s)
			assert.Nil(t, err)
			assert.Equalf(t, tt.want, valid, "Hash(%v)", tt.args.s)
		})
	}
}

func TestHash_Make_Error(t *testing.T) {
	t.Run("ERROR FROM READER", func(t *testing.T) {
		t.Skip()
		expectedErr := errors.New("random read error")
		mockRand := &mockRandReader{err: expectedErr}

		origRandReader := rand.Reader
		defer func() { rand.Reader = origRandReader }()
		rand.Reader = mockRand

		secret, err := utils.MakePassword(utils.Parallelization, "password123")
		assert.NotNil(t, err)
		assert.Empty(t, secret)
		assert.Equal(t, err.Error(), expectedErr.Error())
	})
	t.Run("ERROR FROM HASH", func(t *testing.T) {
		secret, err := utils.MakePassword(-1, "password123")
		assert.NotNil(t, err)
		assert.Empty(t, secret)
	})
}

func TestHash_Compare_Error(t *testing.T) {
	t.Run("ERROR WHEN VALIDATE LEN", func(t *testing.T) {
		valid, err := utils.ComparePassword(utils.Parallelization, "", "")
		assert.NotNil(t, err)
		assert.False(t, valid)
	})

	t.Run("ERROR WHEN HEX DECODE", func(t *testing.T) {
		valid, err := utils.ComparePassword(utils.Parallelization, "1234.gibberish", "")
		assert.NotNil(t, err)
		assert.False(t, valid)
		assert.Equal(t, err.Error(), utils.ErrorPasswordUnableToVerify.Error())
	})

	t.Run("ERROR WHEN HEX DECODE", func(t *testing.T) {
		valid, err := utils.ComparePassword(-1, "1234.1234", "")
		assert.NotNil(t, err)
		assert.False(t, valid)
		assert.Equal(t, err.Error(), utils.ErrorPasswordUnableToVerify.Error())
	})
}
