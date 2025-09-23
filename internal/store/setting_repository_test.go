package store_test

import (
	"context"
	"database/sql"
	"errors"
	"regexp"
	"testing"
	"time"

	"github.com/aasumitro/posbe/internal/model"
	"github.com/aasumitro/posbe/internal/store"
	"github.com/pashagolub/pgxmock/v4"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type settingRepositoryTestCase struct {
	suite.Suite
	mock        pgxmock.PgxPoolIface
	settingRepo store.IStoreSettingRepository
}

func (suite *settingRepositoryTestCase) SetupTest() {
	var err error
	suite.mock, err = pgxmock.NewPool()
	require.NoError(suite.T(), err)
	defer suite.mock.Close()
	suite.settingRepo = store.NewSettingRepository(suite.mock)
}

func (suite *settingRepositoryTestCase) TestSettingRepository_FindByKey_ExpectSuccess() {
	setting := suite.mock.
		NewRows([]string{"key", "value", "created_at", "updated_at"}).
		AddRow("key", "value", time.Now().Unix(), time.Now().Unix())
	q := "SELECT * FROM store_prefs WHERE key = $1 LIMIT 1"
	expectedQuery := regexp.QuoteMeta(q)
	suite.mock.ExpectQuery(expectedQuery).WithArgs("key").WillReturnRows(setting)
	res, err := suite.settingRepo.FindByKey(context.TODO(), "key")
	require.NotNil(suite.T(), res)
	require.NoError(suite.T(), err)
	require.Nil(suite.T(), err)
}
func (suite *settingRepositoryTestCase) TestSettingRepository_FindByKey_ExpectError() {
	q := "SELECT * FROM store_prefs WHERE key = $1 LIMIT 1"
	expectedQuery := regexp.QuoteMeta(q)
	suite.mock.ExpectQuery(expectedQuery).WithArgs("key").WillReturnError(sql.ErrNoRows)
	res, err := suite.settingRepo.FindByKey(context.TODO(), "key")
	require.Nil(suite.T(), res)
	require.Error(suite.T(), err)
	require.NotNil(suite.T(), err)
}

func (suite *settingRepositoryTestCase) TestSettingRepository_GetAll_ExpectSuccess() {
	setting := suite.mock.
		NewRows([]string{"key", "value", "created_at", "updated_at"}).
		AddRow("key", "value", time.Now().Unix(), time.Now().Unix()).
		AddRow("other", "other", time.Now().Unix(), time.Now().Unix())
	q := "SELECT * FROM store_prefs"
	expectedQuery := regexp.QuoteMeta(q)
	suite.mock.ExpectQuery(expectedQuery).WillReturnRows(setting)
	res, err := suite.settingRepo.GetAll(context.TODO())
	require.NotNil(suite.T(), res)
	require.NoError(suite.T(), err)
	require.Nil(suite.T(), err)
}
func (suite *settingRepositoryTestCase) TestSettingRepository_GetAll_ExpectError() {
	q := "SELECT * FROM store_prefs"
	expectedQuery := regexp.QuoteMeta(q)
	suite.T().Run("error from query", func(t *testing.T) {
		suite.mock.ExpectQuery(expectedQuery).WillReturnError(errors.New(""))
		res, err := suite.settingRepo.GetAll(context.TODO())
		require.Nil(suite.T(), res)
		require.Error(suite.T(), err)
		require.NotNil(suite.T(), err)
	})
	suite.T().Run("error from scan", func(t *testing.T) {
		settings := suite.mock.
			NewRows([]string{"key", "value", "created_at", "updated_at"}).
			AddRow("somekey", "somevalue", "not-a-timestamp", time.Now().Unix())
		suite.mock.ExpectQuery(expectedQuery).WillReturnRows(settings)
		res, err := suite.settingRepo.GetAll(context.TODO())
		require.Nil(suite.T(), res)
		require.Error(suite.T(), err)
		require.NotNil(suite.T(), err)
	})
}

func (suite *settingRepositoryTestCase) TestSettingRepository_Update_ExpectSuccess() {
	q := `
	INSERT INTO store_prefs (key, value, created_at, updated_at)
			VALUES ($1, $2, extract(epoch from now()), extract(epoch from now()))
			ON CONFLICT (key) DO UPDATE
			SET value = EXCLUDED.value,
			    updated_at = extract(epoch from now())
	`
	expectedQuery := regexp.QuoteMeta(q)
	suite.mock.ExpectExec(expectedQuery).WithArgs("k", "v").WillReturnResult(pgxmock.NewResult("UPDATE", 1))
	err := suite.settingRepo.Update(context.TODO(), &model.StoreSetting{"k": "v"})
	require.NoError(suite.T(), err)
	require.Nil(suite.T(), err)
}
func (suite *settingRepositoryTestCase) TestSettingRepository_Update_ExpectError() {
	q := `
	INSERT INTO store_prefs (key, value, created_at, updated_at)
			VALUES ($1, $2, extract(epoch from now()), extract(epoch from now()))
			ON CONFLICT (key) DO UPDATE
			SET value = EXCLUDED.value,
			    updated_at = extract(epoch from now())
	`
	expectedQuery := regexp.QuoteMeta(q)
	suite.T().Run("no params", func(t *testing.T) {
		err := suite.settingRepo.Update(context.TODO(), &model.StoreSetting{})
		require.NoError(suite.T(), err)
		require.Nil(suite.T(), err)
	})
	suite.T().Run("error from query", func(t *testing.T) {
		suite.mock.ExpectExec(expectedQuery).WithArgs("k", "v").WillReturnError(errors.New(""))
		err := suite.settingRepo.Update(context.TODO(), &model.StoreSetting{"k": "v"})
		require.Error(suite.T(), err)
		require.NotNil(suite.T(), err)
	})
}

func (suite *settingRepositoryTestCase) AfterTest(_, _ string) {
	require.NoError(suite.T(), suite.mock.ExpectationsWereMet())
}

func TestStoreSettingRepository(t *testing.T) {
	suite.Run(t, new(settingRepositoryTestCase))
}
