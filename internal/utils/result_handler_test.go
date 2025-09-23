package utils_test

import (
	"errors"
	"testing"

	"github.com/aasumitro/posbe/internal/model"
	"github.com/aasumitro/posbe/internal/utils"
	"github.com/stretchr/testify/assert"
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
	tests := []testCase[model.Role]{
		{
			name: "Validate Row Should Success",
			args: args[model.Role]{
				data: &model.Role{ID: 1, Name: "ipsum"},
				err:  nil,
			},
			wantValueData: &model.Role{ID: 1, Name: "ipsum"},
			wantErrData:   nil,
		},
		{
			name: "Validate Row Should Error",
			args: args[model.Role]{
				data: nil,
				err:  errors.New("LOREM"),
			},
			wantValueData: nil,
			wantErrData:   &utils.ServiceError{Code: 500, Message: "LOREM"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotValueData, gotErrData := utils.HandleSingleResult("test", tt.args.data, tt.args.err)
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
	tests := []testCase[model.Role]{
		{
			name: "Validate Row Should Success",
			args: args[model.Role]{
				data: []*model.Role{
					{ID: 1, Name: "ipsum"},
					{ID: 2, Name: "lorem"},
				},
				err: nil,
			},
			wantValueData: []*model.Role{
				{ID: 1, Name: "ipsum"},
				{ID: 2, Name: "lorem"},
			},
			wantErrData: nil,
		},
		{
			name: "Validate Row Should Error",
			args: args[model.Role]{
				data: nil,
				err:  errors.New("LOREM"),
			},
			wantValueData: nil,
			wantErrData:   &utils.ServiceError{Code: 500, Message: "LOREM"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotValueData, gotErrData := utils.HandleMultipleResults("test", tt.args.data, tt.args.err)
			assert.Equalf(t, tt.wantValueData, gotValueData, "HandleMultipleResults(%v, %v)", tt.args.data, tt.args.err)
			assert.Equalf(t, tt.wantErrData, gotErrData, "HandleMultipleResults(%v, %v)", tt.args.data, tt.args.err)
		})
	}
}
