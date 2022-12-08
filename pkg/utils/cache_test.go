package utils

import (
	"context"
	"github.com/go-redis/redis/v9"
	"testing"
)

func TestRedisCache_CacheFirstData(t *testing.T) {
	type fields struct {
		Ctx     context.Context
		RdpConn *redis.Client
	}
	type args struct {
		i *CacheDataSupplied
	}
	tests := []struct {
		name     string
		fields   fields
		args     args
		wantData any
		wantErr  *ServiceError
	}{
		{},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			//cache := &RedisCache{
			//	Ctx:     tt.fields.Ctx,
			//	RdpConn: tt.fields.RdpConn,
			//}
			//gotData, gotErr := cache.CacheFirstData(tt.args.i)
			//assert.Equalf(t, tt.wantData, gotData, "CacheFirstData(%v)", tt.args.i)
			//assert.Equalf(t, tt.wantErr, gotErr, "CacheFirstData(%v)", tt.args.i)
		})
	}
}
