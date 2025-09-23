package utils_test

import (
	"testing"

	"github.com/aasumitro/posbe/internal/utils"
	"github.com/golang-jwt/jwt/v5"
	"github.com/stretchr/testify/assert"
)

func TestParseJWT(t *testing.T) {
	scenarios := []struct {
		token        string
		secret       string
		expectError  bool
		expectClaims jwt.MapClaims
	}{
		// invalid formatted JWT
		{
			"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJuYW1lIjoidGVzdCJ9",
			"test",
			true,
			nil,
		},
		// properly formatted JWT with INVALID claims and INVALID secret
		// {"name": "test", "exp": 1516239022}
		{
			"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJuYW1lIjoidGVzdCIsImV4cCI6MTUxNjIzOTAyMn0.xYHirwESfSEW3Cq2BL47CEASvD_p_ps3QCA54XtNktU",
			"invalid",
			true,
			nil,
		},
		// properly formatted JWT with INVALID claims and VALID secret
		// {"name": "test", "exp": 1516239022}
		{
			"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJuYW1lIjoidGVzdCIsImV4cCI6MTUxNjIzOTAyMn0.xYHirwESfSEW3Cq2BL47CEASvD_p_ps3QCA54XtNktU",
			"test",
			true,
			nil,
		},
		// properly formatted JWT with VALID claims and INVALID secret
		// {"name": "test", "exp": 1898636137}
		{
			"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJuYW1lIjoidGVzdCIsImV4cCI6MTg5ODYzNjEzN30.gqRkHjpK5s1PxxBn9qPaWEWxTbpc1PPSD-an83TsXRY",
			"invalid",
			true,
			nil,
		},
		// properly formatted EXPIRED JWT with VALID secret
		// {"name": "test", "exp": 1652097610}
		{
			"eyJhbGciOiJIUzI1NiJ9.eyJuYW1lIjoidGVzdCIsImV4cCI6OTU3ODczMzc0fQ.0oUUKUnsQHs4nZO1pnxQHahKtcHspHu4_AplN2sGC4A",
			"test",
			true,
			nil,
		},
		// properly formatted JWT with VALID claims and VALID secret
		// {"name": "test", "exp": 1898636137}
		{
			"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJuYW1lIjoidGVzdCIsImV4cCI6MTg5ODYzNjEzN30.gqRkHjpK5s1PxxBn9qPaWEWxTbpc1PPSD-an83TsXRY",
			"test",
			false,
			jwt.MapClaims{"name": "test", "exp": 1898636137.0},
		},
		// properly formatted JWT with VALID claims (without exp) and VALID secret
		// {"name": "test"}
		{
			"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJuYW1lIjoidGVzdCJ9.ml0QsTms3K9wMygTu41ZhKlTyjmW9zHQtoS8FUsCCjU",
			"test",
			false,
			jwt.MapClaims{"name": "test"},
		},
	}

	for _, scenario := range scenarios {
		result, err := utils.ParseJWT(scenario.token, scenario.secret)
		if scenario.expectError && err == nil {
			assert.NoError(t, err)
		}
		if !scenario.expectError && err != nil {
			assert.Error(t, err)
		}
		assert.Equal(t, len(result), len(scenario.expectClaims))
		for k, v := range scenario.expectClaims {
			v2, ok := result[k]
			assert.True(t, ok)
			assert.Equal(t, v, v2)
		}
	}
}

func TestNewJWT(t *testing.T) {
	scenarios := []struct {
		claims      jwt.MapClaims
		key         string
		duration    int64
		expectError bool
	}{
		// empty, zero duration
		{jwt.MapClaims{}, "", 0, true},
		// empty, 10 seconds duration
		{jwt.MapClaims{}, "", 10, false},
		// non-empty, 10 seconds duration
		{jwt.MapClaims{"name": "test"}, "test", 10, false},
	}

	for _, scenario := range scenarios {
		token, tokenErr := utils.NewJWT(scenario.claims, scenario.key, scenario.duration)
		assert.NoError(t, tokenErr)
		claims, parseErr := utils.ParseJWT(token, scenario.key)
		hasParseErr := parseErr != nil
		assert.Equal(t, hasParseErr, scenario.expectError)
		if scenario.expectError {
			continue
		}
		_, ok := claims["exp"]
		assert.True(t, ok)
		delete(claims, "exp")
		assert.Equal(t, len(claims), len(scenario.claims))
		for j := range claims {
			assert.Equal(t, claims[j], scenario.claims[j])
		}
	}
}
