package utils_test

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/aasumitro/posbe/domain"
	"github.com/aasumitro/posbe/pkg/utils"
	"github.com/alicebob/miniredis/v2"
	"github.com/go-redis/redis/v9"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRedisCache_CacheFirstData(t *testing.T) {
	type fields struct {
		Ctx     context.Context
		RdpConn *redis.Client
	}
	type args struct {
		i *utils.CacheDataSupplied
	}
	testData, _ := json.Marshal(domain.Role{ID: 1, Name: "test"})
	tests := []struct {
		name     string
		fields   fields
		args     args
		wantData any
		wantErr  assert.ErrorAssertionFunc
	}{
		{
			name: "TEST RETURN NIL",
			fields: fields{
				Ctx: context.TODO(),
				RdpConn: redis.NewClient(&redis.Options{
					Addr: miniredis.RunT(t).Addr(),
				}),
			},
			args: args{
				&utils.CacheDataSupplied{
					Key: "lorem",
					TTL: 0,
					CbF: func() (data any, err error) {
						return nil, err
					},
				},
			},
			wantData: nil,
			wantErr:  assert.NoError,
		},
		{
			name: "TEST RETURN OBJECT",
			fields: fields{
				Ctx: context.TODO(),
				RdpConn: redis.NewClient(&redis.Options{
					Addr: miniredis.RunT(t).Addr(),
				}),
			},
			args: args{
				&utils.CacheDataSupplied{
					Key: "lorem",
					TTL: 0,
					CbF: func() (data any, err error) {
						return domain.Role{ID: 1, Name: "test"}, err
					},
				},
			},
			wantData: domain.Role{ID: 1, Name: "test"},
			wantErr:  assert.NoError,
		},
		{
			name: "TEST RETURN STRING",
			fields: fields{
				Ctx: context.TODO(),
				RdpConn: redis.NewClient(&redis.Options{
					Addr: miniredis.RunT(t).Addr(),
				}),
			},
			args: args{
				&utils.CacheDataSupplied{
					Key: "lorem",
					TTL: 0,
					CbF: func() (data any, err error) {
						return testData, err
					},
				},
			},
			wantData: testData,
			wantErr:  assert.NoError,
		},
		{
			name: "TEST RETURN ERROR",
			fields: fields{
				Ctx: context.TODO(),
				RdpConn: redis.NewClient(&redis.Options{
					Addr: miniredis.RunT(t).Addr(),
				}),
			},
			args: args{
				&utils.CacheDataSupplied{
					Key: "lorem",
					TTL: 0,
					CbF: func() (data any, err error) {
						return nil, errors.New("lorem ipsum")
					},
				},
			},
			wantData: nil,
			wantErr:  assert.Error,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cache := &utils.RedisCache{
				Ctx:     tt.fields.Ctx,
				RdpConn: tt.fields.RdpConn,
			}

			if tt.name == "TEST RETURN STRING" {
				cache.RdpConn.Set(cache.Ctx, "lorem", testData, 0)
			}

			gotData, err := cache.CacheFirstData(tt.args.i)
			if !tt.wantErr(t, err, fmt.Sprintf("CacheFirstData(%v)", tt.args.i)) {
				return
			}
			if tt.name != "TEST RETURN STRING" {
				assert.Equalf(t, tt.wantData, gotData, "CacheFirstData(%v)", tt.args.i)
			} else {
				assert.Equalf(t, tt.wantData, []byte(gotData.(string)), "CacheFirstData(%v)", tt.args.i)
			}
		})
	}
}
