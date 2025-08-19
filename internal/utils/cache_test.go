package utils_test

import (
	"context"
	"encoding/json"
	"errors"
	"testing"
	"time"

	"github.com/aasumitro/posbe/internal/model"
	"github.com/aasumitro/posbe/internal/utils"
	"github.com/alicebob/miniredis/v2"
	"github.com/redis/go-redis/v9"
	"github.com/stretchr/testify/assert"
)

func TestRedisCache_CacheFirstData(t *testing.T) {
	ctx := context.TODO()
	mr := miniredis.RunT(t)
	rdb := redis.NewClient(&redis.Options{Addr: mr.Addr()})
	// Sample test role
	testRole := model.Role{ID: 1, Name: "test"}

	tests := []struct {
		name     string
		supply   *utils.CacheDataSupplied[model.Role]
		prepare  func()
		wantData model.Role
		wantErr  assert.ErrorAssertionFunc
	}{
		{
			name: "TEST RETURN OBJECT (from callback)",
			supply: &utils.CacheDataSupplied[model.Role]{
				Key: "role1",
				TTL: time.Minute,
				CbF: func() (model.Role, error) {
					return testRole, nil
				},
			},
			prepare:  func() {},
			wantData: testRole,
			wantErr:  assert.NoError,
		},
		{
			name: "TEST RETURN ERROR",
			supply: &utils.CacheDataSupplied[model.Role]{
				Key: "role2",
				TTL: time.Minute,
				CbF: func() (model.Role, error) {
					return model.Role{}, errors.New("lorem ipsum")
				},
			},
			prepare:  func() {},
			wantData: model.Role{},
			wantErr:  assert.Error,
		},
		{
			name: "TEST RETURN OBJECT (from redis)",
			supply: &utils.CacheDataSupplied[model.Role]{
				Key: "role3",
				TTL: time.Minute,
				CbF: func() (model.Role, error) {
					return model.Role{}, nil // should not be called
				},
			},
			prepare: func() {
				// manually store in redis
				b, _ := json.Marshal(testRole)
				_ = rdb.Set(ctx, "role3", b, 0).Err()
			},
			wantData: testRole,
			wantErr:  assert.NoError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.prepare()
			got, err := utils.CacheFirstData(ctx, rdb, tt.supply)
			if !tt.wantErr(t, err) {
				return
			}
			assert.Equal(t, tt.wantData, got)
		})
	}
}
