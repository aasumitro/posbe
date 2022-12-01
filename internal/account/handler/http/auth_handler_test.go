package http_test

import (
	"github.com/stretchr/testify/suite"
	"testing"
)

type authHandlerTestSuite struct {
	suite.Suite
}

func (suite *authHandlerTestSuite) SetupSuite() {
	// TODO
}

func (suite *authHandlerTestSuite) TestAuthHandler_T_K() {
	// TODO
}

func TestAuthHandlerService(t *testing.T) {
	suite.Run(t, new(authHandlerTestSuite))
}
