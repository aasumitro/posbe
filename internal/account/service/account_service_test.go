package service_test

import (
	"database/sql"
	"github.com/stretchr/testify/suite"
	"testing"
)

type accountTestSuite struct {
	suite.Suite
	Db *sql.DB
}

func (suite *accountTestSuite) SetupSuite() {
	// TODO
}

func (suite *accountTestSuite) AfterTest(_, _ string) {
	// TODO
}

func (suite *accountTestSuite) TestAccountService_T_K() {
	// TODO
}

func TestAccountService(t *testing.T) {
	suite.Run(t, new(accountTestSuite))
}
