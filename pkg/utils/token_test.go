package utils_test

import (
	"github.com/aasumitro/posbe/pkg/utils"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestJSONWebToken_ClaimJWTToken(t *testing.T) {
	type fields struct {
		Issuer    string
		SecretKey []byte
		Payload   interface{}
		IssuedAt  time.Time
		ExpiredAt time.Time
	}

	issuedAt, _ := time.Parse("",
		"2022-12-05 17:57:44.321843 +0800 WITA m=+25.737606459")
	expiredAt, _ := time.Parse("",
		"2022-12-06 17:57:44.321851 +0800 WITA m=+86425.737614876")

	tests := []struct {
		name    string
		fields  fields
		want    string
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name: "NEW JWT TEST SHOULD SUCCESS",
			fields: fields{
				Issuer: "POSBE_TEST",
				Payload: map[string]string{
					"data": "hello world",
				},
				IssuedAt:  issuedAt,
				ExpiredAt: expiredAt,
			},
			want:    "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJQT1NCRV9URVNUIiwiZXhwIjotNjIxMzU1OTY4MDAsImlhdCI6LTYyMTM1NTk2ODAwLCJwYXlsb2FkIjp7ImRhdGEiOiJoZWxsbyB3b3JsZCJ9fQ.-_tfeKKhqSRP2H_pVg4f_spkX_Z1Lo1nuiu09OFFvO0",
			wantErr: assert.NoError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			j := &utils.JSONWebToken{
				Issuer:    tt.fields.Issuer,
				SecretKey: tt.fields.SecretKey,
				Payload:   tt.fields.Payload,
				IssuedAt:  tt.fields.IssuedAt,
				ExpiredAt: tt.fields.ExpiredAt,
			}
			got, err := j.ClaimJWTToken()
			if !tt.wantErr(t, err, "ClaimJWTToken()") {
				return
			}
			assert.Equalf(t, tt.want, got, "ClaimJWTToken()")
		})
	}
}
