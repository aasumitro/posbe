package utils

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io"
)

// MockJSONRequest
// accepted method GET, POST, PUT, DELETE
func MockJSONRequest(c *gin.Context, method string, cType string, content interface{}) {
	c.Request.Method = method
	c.Request.Header.Set("Content-Type", cType)
	b, _ := json.Marshal(content)
	c.Request.Body = io.NopCloser(bytes.NewBuffer(b))
}
