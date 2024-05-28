package sql_test

import (
	"context"
	"errors"
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/aasumitro/posbe/config"
	repoSql "github.com/aasumitro/posbe/internal/store/repository/sql"
	"github.com/aasumitro/posbe/pkg/model"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type storePrefRepositoryTestSuite struct {
	suite.Suite
	mock      sqlmock.Sqlmock
	storePref model.IStorePrefRepository
}

func (suite *storePrefRepositoryTestSuite) SetupSuite() {
	var err error

	config.PostgresPool, suite.mock, err = sqlmock.New(
		sqlmock.QueryMatcherOption(
			sqlmock.QueryMatcherRegexp))
	require.NoError(suite.T(), err)

	suite.storePref = repoSql.NewStorePrefSQLRepository()
}

func (suite *storePrefRepositoryTestSuite) AfterTest(_, _ string) {
	require.NoError(suite.T(), suite.mock.ExpectationsWereMet())
}

func (suite *storePrefRepositoryTestSuite) TestStorePrefsRepository_Find_ExpectedSuccess() {
	pref := suite.mock.
		NewRows([]string{"key", "value", "created_at", "updated_at"}).
		AddRow("test", "test", 123, 123)
	q := "SELECT * FROM store_prefs WHERE key = $1 LIMIT 1"
	expectedQuery := regexp.QuoteMeta(q)
	suite.mock.ExpectQuery(expectedQuery).WillReturnRows(pref)
	res, err := suite.storePref.Find(context.TODO(), "test")
	require.Nil(suite.T(), err)
	require.NoError(suite.T(), err)
	require.NotNil(suite.T(), res)
}

func (suite *storePrefRepositoryTestSuite) TestStorePrefsRepository_Find_ExpectedError() {
	pref := suite.mock.
		NewRows([]string{"key", "value", "created_at", "updated_at"}).
		AddRow(nil, nil, nil, nil)
	q := "SELECT * FROM store_prefs WHERE key = $1 LIMIT 1"
	expectedQuery := regexp.QuoteMeta(q)
	suite.mock.ExpectQuery(expectedQuery).WillReturnRows(pref)
	res, err := suite.storePref.Find(context.TODO(), "test")
	require.Nil(suite.T(), res)
	require.NotNil(suite.T(), err)
}

func (suite *storePrefRepositoryTestSuite) TestStorePrefsRepository_All_ExpectedSuccess() {
	pref := suite.mock.
		NewRows([]string{"key", "value", "created_at", "updated_at"}).
		AddRow("test", "test", 123, 123)
	q := "SELECT * FROM store_prefs"
	expectedQuery := regexp.QuoteMeta(q)
	suite.mock.ExpectQuery(expectedQuery).WillReturnRows(pref)
	res, err := suite.storePref.All(context.TODO())
	require.Nil(suite.T(), err)
	require.NoError(suite.T(), err)
	require.NotNil(suite.T(), res)
}

func (suite *storePrefRepositoryTestSuite) TestStorePrefsRepository_All_ExpectedErrorFromQuery() {
	pref := suite.mock.
		NewRows([]string{"key", "value", "created_at", "updated_at"}).
		AddRow("test", "test", 123, 123)
	q := "SELECT * FROM store_prefs"
	expectedQuery := regexp.QuoteMeta(q)
	suite.mock.ExpectQuery(expectedQuery).
		WillReturnRows(pref).
		WillReturnError(errors.New(""))
	res, err := suite.storePref.All(context.TODO())
	require.NotNil(suite.T(), err)
	require.Nil(suite.T(), res)
}

func (suite *storePrefRepositoryTestSuite) TestStorePrefsRepository_All_ExpectedErrorFromScan() {
	pref := suite.mock.
		NewRows([]string{"key", "value", "created_at", "updated_at"}).
		AddRow(nil, nil, nil, nil)
	q := "SELECT * FROM store_prefs"
	expectedQuery := regexp.QuoteMeta(q)
	suite.mock.ExpectQuery(expectedQuery).WillReturnRows(pref)
	res, err := suite.storePref.All(context.TODO())
	require.Nil(suite.T(), res)
	require.NotNil(suite.T(), err)
}

func (suite *storePrefRepositoryTestSuite) TestStorePrefsRepository_Update_ExpectedSuccess() {
	pref := suite.mock.
		NewRows([]string{"key", "value", "created_at", "updated_at"}).
		AddRow("test", "test", 123, 123)
	q := "UPDATE store_prefs SET value = $1, updated_at = $2 WHERE key = $3 RETURNING *"
	expectedQuery := regexp.QuoteMeta(q)
	suite.mock.ExpectQuery(expectedQuery).
		WithArgs("test", time.Now().Unix(), "test").
		WillReturnRows(pref).
		WillReturnError(nil)
	res, err := suite.storePref.Update(context.TODO(), "test", "test")
	require.Nil(suite.T(), err)
	require.NoError(suite.T(), err)
	require.NotNil(suite.T(), res)
}

func (suite *storePrefRepositoryTestSuite) TestStorePrefsRepository_Update_ExpectedError() {
	pref := suite.mock.
		NewRows([]string{"key", "value", "created_at", "updated_at"}).
		AddRow(nil, nil, nil, nil)
	q := "UPDATE store_prefs SET value = $1, updated_at = $2 WHERE key = $3 RETURNING *"
	expectedQuery := regexp.QuoteMeta(q)
	suite.mock.ExpectQuery(expectedQuery).
		WithArgs("test", time.Now().Unix(), "test").
		WillReturnRows(pref)
	res, err := suite.storePref.Update(context.TODO(), "test", "test")
	require.Nil(suite.T(), res)
	require.NotNil(suite.T(), err)
}

func TestStorePrefRepository(t *testing.T) {
	suite.Run(t, new(storePrefRepositoryTestSuite))
}
