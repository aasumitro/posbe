package utils_test

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/aasumitro/posbe/config"
	"github.com/aasumitro/posbe/internal/utils"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
)

func TestAuthN(t *testing.T) {
	viper.Reset()
	viper.SetConfigFile("../../misc/conf/.env.example")
	viper.SetConfigType("dotenv")
	config.LoadWith(context.Background())
	gin.SetMode(gin.TestMode)

	// Create a valid JWT for testing
	claims := map[string]any{
		"id":        "user-123",
		"role_id":   "role-456",
		"role_name": "admin",
		"exp":       time.Now().Add(time.Hour).Unix(),
	}
	token, _ := utils.NewJWT(claims, config.Instance.JWTSecretKey, 30)

	tests := []struct {
		name           string
		cookieValue    string
		setCookie      bool
		expectedStatus int
		expectUserID   bool
	}{
		{
			name:           "Missing authn cookie",
			setCookie:      false,
			expectedStatus: http.StatusUnauthorized,
			expectUserID:   false,
		},
		{
			name:           "Invalid token in cookie",
			setCookie:      true,
			cookieValue:    "invalid.jwt.token",
			expectedStatus: http.StatusUnauthorized,
			expectUserID:   false,
		},
		{
			name:           "Valid token in cookie",
			setCookie:      true,
			cookieValue:    token,
			expectedStatus: http.StatusOK,
			expectUserID:   true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			router := gin.New()
			router.Use(utils.AuthN())
			router.GET("/test", func(c *gin.Context) {
				if tc.expectUserID {
					uid, exists := c.Get("user_id")
					assert.True(t, exists)
					assert.Equal(t, "user-123", uid)
				}
				c.Status(http.StatusOK)
			})

			req, _ := http.NewRequest(http.MethodGet, "/test", nil)
			if tc.setCookie {
				req.AddCookie(&http.Cookie{
					Name:  "authn",
					Value: tc.cookieValue,
				})
			}

			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)

			assert.Equal(t, tc.expectedStatus, w.Code)
		})
	}
}

func TestAuthZ(t *testing.T) {
	gin.SetMode(gin.TestMode)

	tests := []struct {
		name          string
		acceptedRoles []string
		contextRole   any
		expectedCode  int
	}{
		{
			name:          "Allow any role when accepted is *",
			acceptedRoles: []string{"*"},
			contextRole:   nil, // no role_name
			expectedCode:  http.StatusOK,
		},
		{
			name:          "No role_name in context",
			acceptedRoles: []string{"admin"},
			contextRole:   nil,
			expectedCode:  http.StatusUnauthorized,
		},
		{
			name:          "Role not in accepted list",
			acceptedRoles: []string{"admin"},
			contextRole:   "user",
			expectedCode:  http.StatusUnauthorized,
		},
		{
			name:          "Role in accepted list",
			acceptedRoles: []string{"admin", "user"},
			contextRole:   "user",
			expectedCode:  http.StatusOK,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			router := gin.New()

			// Apply middleware
			router.Use(func(c *gin.Context) {
				if tc.contextRole != nil {
					c.Set("role_name", tc.contextRole)
				}
			})
			router.Use(utils.AuthZ(tc.acceptedRoles))
			router.GET("/test", func(c *gin.Context) {
				c.Status(http.StatusOK)
			})

			req, _ := http.NewRequest(http.MethodGet, "/test", nil)
			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)

			assert.Equal(t, tc.expectedCode, w.Code)
		})
	}
}
