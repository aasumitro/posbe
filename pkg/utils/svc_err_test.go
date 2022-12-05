package utils_test

import (
	"errors"
	"github.com/aasumitro/posbe/domain"
	"github.com/aasumitro/posbe/pkg/utils"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestValidateDataRow(t *testing.T) {
	type args[T any] struct {
		data *T
		err  error
	}
	type testCase[T any] struct {
		name          string
		args          args[T]
		wantValueData *T
		wantErrData   *utils.ServiceError
	}
	tests := []testCase[domain.Role]{
		{
			name: "Validate Row Should Success",
			args: args[domain.Role]{
				data: &domain.Role{ID: 1, Name: "ipsum"},
				err:  nil,
			},
			wantValueData: &domain.Role{ID: 1, Name: "ipsum"},
			wantErrData:   nil,
		},
		{
			name: "Validate Row Should Error",
			args: args[domain.Role]{
				data: nil,
				err:  errors.New("LOREM"),
			},
			wantValueData: nil,
			wantErrData:   &utils.ServiceError{Code: 500, Message: "LOREM"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotValueData, gotErrData := utils.ValidateDataRow(tt.args.data, tt.args.err)
			assert.Equalf(t, tt.wantValueData, gotValueData, "ValidateDataRow(%v, %v)", tt.args.data, tt.args.err)
			assert.Equalf(t, tt.wantErrData, gotErrData, "ValidateDataRow(%v, %v)", tt.args.data, tt.args.err)
		})
	}
}

func TestValidateDataRows(t *testing.T) {
	type args[T any] struct {
		data []*T
		err  error
	}
	type testCase[T any] struct {
		name          string
		args          args[T]
		wantValueData []*T
		wantErrData   *utils.ServiceError
	}
	tests := []testCase[domain.Role]{
		{
			name: "Validate Row Should Success",
			args: args[domain.Role]{
				data: []*domain.Role{
					{ID: 1, Name: "ipsum"},
					{ID: 2, Name: "lorem"},
				},
				err: nil,
			},
			wantValueData: []*domain.Role{
				{ID: 1, Name: "ipsum"},
				{ID: 2, Name: "lorem"},
			},
			wantErrData: nil,
		},
		{
			name: "Validate Row Should Error",
			args: args[domain.Role]{
				data: nil,
				err:  errors.New("LOREM"),
			},
			wantValueData: nil,
			wantErrData:   &utils.ServiceError{Code: 500, Message: "LOREM"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotValueData, gotErrData := utils.ValidateDataRows(tt.args.data, tt.args.err)
			assert.Equalf(t, tt.wantValueData, gotValueData, "ValidateDataRows(%v, %v)", tt.args.data, tt.args.err)
			assert.Equalf(t, tt.wantErrData, gotErrData, "ValidateDataRows(%v, %v)", tt.args.data, tt.args.err)
		})
	}
}
