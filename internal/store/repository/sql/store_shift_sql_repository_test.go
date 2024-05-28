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

type shiftRepositoryTestSuite struct {
	suite.Suite
	mock      sqlmock.Sqlmock
	shiftRepo model.IStoreShiftRepository
}

func (suite *shiftRepositoryTestSuite) SetupSuite() {
	var err error

	config.PostgresPool, suite.mock, err = sqlmock.New(
		sqlmock.QueryMatcherOption(
			sqlmock.QueryMatcherRegexp))
	require.NoError(suite.T(), err)

	suite.shiftRepo = repoSql.NewStoreShiftSQLRepository()
}

func (suite *shiftRepositoryTestSuite) AfterTest(_, _ string) {
	require.NoError(suite.T(), suite.mock.ExpectationsWereMet())
}

func TestShiftRepository(t *testing.T) {
	suite.Run(t, new(shiftRepositoryTestSuite))
}

func (suite *shiftRepositoryTestSuite) TestShiftRepository_All_ExpectReturnData() {
	shifts := suite.mock.
		NewRows([]string{"id", "name", "start_time", "end_time", "created_at", "updated_at"}).
		AddRow(1, "test", time.Now().Unix(), time.Now().Unix(), time.Now().Unix(), time.Now().Unix()).
		AddRow(2, "test 2", time.Now().Unix(), time.Now().Unix(), time.Now().Unix(), time.Now().Unix())
	q := "SELECT * FROM shifts"
	expectedQuery := regexp.QuoteMeta(q)
	suite.mock.ExpectQuery(expectedQuery).WillReturnRows(shifts)
	res, err := suite.shiftRepo.All(context.TODO())
	require.Nil(suite.T(), err)
	require.NoError(suite.T(), err)
	require.NotNil(suite.T(), res)
}
func (suite *shiftRepositoryTestSuite) TestShiftRepository_All_ExpectErrorQuery() {
	q := "SELECT * FROM shifts"
	expectedQuery := regexp.QuoteMeta(q)
	suite.mock.ExpectQuery(expectedQuery).WillReturnError(errors.New(""))
	res, err := suite.shiftRepo.All(context.TODO())
	require.NotNil(suite.T(), err)
	require.Nil(suite.T(), res)
}
func (suite *shiftRepositoryTestSuite) TestShiftRepository_All_ExpectErrorScan() {
	shifts := suite.mock.
		NewRows([]string{"id", "name", "start_time", "end_time", "created_at", "updated_at"}).
		AddRow(1, "test", time.Now().Unix(), time.Now().Unix(), time.Now().Unix(), time.Now().Unix()).
		AddRow(nil, nil, nil, nil, nil, nil)
	q := "SELECT * FROM shifts"
	expectedQuery := regexp.QuoteMeta(q)
	suite.mock.ExpectQuery(expectedQuery).WillReturnRows(shifts)
	res, err := suite.shiftRepo.All(context.TODO())
	require.Nil(suite.T(), res)
	require.NotNil(suite.T(), err)
}

func (suite *shiftRepositoryTestSuite) TestShiftRepository_Create_ExpectSuccess() {
	shift := &model.Shift{
		Name:      "shift 1",
		StartTime: time.Now().Add(-1 * time.Hour).Unix(),
		EndTime:   time.Now().Add(1 * time.Hour).Unix(),
	}
	rows := suite.mock.
		NewRows([]string{"id", "name", "start_time", "end_time", "created_at", "updated_at"}).
		AddRow(1, "shift 1", time.Now().Unix(), time.Now().Unix(), time.Now().Unix(), time.Now().Unix())
	q := "INSERT INTO shifts "
	q += "(name, start_time, end_time, created_at) "
	q += " VALUES ($1, $2, $3, $4) RETURNING *"
	expectedQuery := regexp.QuoteMeta(q)
	suite.mock.ExpectQuery(expectedQuery).
		WithArgs(shift.Name, shift.StartTime,
			shift.EndTime, time.Now().Unix()).
		WillReturnRows(rows).
		WillReturnError(nil)
	res, err := suite.shiftRepo.Create(context.TODO(), shift)
	require.Nil(suite.T(), err)
	require.NotNil(suite.T(), res)
}
func (suite *shiftRepositoryTestSuite) TestShiftRepository_Create_ExpectError() {
	shift := &model.Shift{
		Name:      "shift 1",
		StartTime: time.Now().Add(-1 * time.Hour).Unix(),
		EndTime:   time.Now().Add(1 * time.Hour).Unix(),
	}
	rows := suite.mock.
		NewRows([]string{"id", "name", "start_time", "end_time", "created_at", "updated_at"}).
		AddRow(nil, nil, nil, nil, nil, nil)
	q := "INSERT INTO shifts "
	q += "(name, start_time, end_time, created_at) "
	q += " VALUES ($1, $2, $3, $4) RETURNING *"
	expectedQuery := regexp.QuoteMeta(q)
	suite.mock.ExpectQuery(expectedQuery).
		WithArgs(shift.Name, shift.StartTime,
			shift.EndTime, time.Now().Unix()).
		WillReturnRows(rows).
		WillReturnError(nil)
	res, err := suite.shiftRepo.Create(context.TODO(), shift)
	require.Nil(suite.T(), res)
	require.NotNil(suite.T(), err)
}

func (suite *shiftRepositoryTestSuite) TestShiftRepository_Update_ExpectSuccess() {
	shift := &model.Shift{
		ID:        1,
		Name:      "shift 1",
		StartTime: time.Now().Add(-1 * time.Hour).Unix(),
		EndTime:   time.Now().Add(1 * time.Hour).Unix(),
	}
	rows := suite.mock.
		NewRows([]string{"id", "name", "start_time", "end_time", "created_at", "updated_at"}).
		AddRow(1, "shift 1", time.Now().Unix(), time.Now().Unix(), time.Now().Unix(), time.Now().Unix())
	q := "UPDATE shifts SET "
	q += "name = $1, start_time = $2, "
	q += "end_time = $3, updated_at = $4 "
	q += " WHERE id = $5 RETURNING *"
	expectedQuery := regexp.QuoteMeta(q)
	suite.mock.ExpectQuery(expectedQuery).
		WithArgs(shift.Name, shift.StartTime,
			shift.EndTime, time.Now().Unix(), shift.ID).
		WillReturnRows(rows).
		WillReturnError(nil)
	res, err := suite.shiftRepo.Update(context.TODO(), shift)
	require.Nil(suite.T(), err)
	require.NotNil(suite.T(), res)
}
func (suite *shiftRepositoryTestSuite) TestShiftRepository_Update_ExpectError() {
	shift := &model.Shift{
		ID:        1,
		Name:      "shift 1",
		StartTime: time.Now().Add(-1 * time.Hour).Unix(),
		EndTime:   time.Now().Add(1 * time.Hour).Unix(),
	}
	rows := suite.mock.
		NewRows([]string{"id", "name", "start_time", "end_time", "created_at", "updated_at"}).
		AddRow(nil, nil, nil, nil, nil, nil)
	q := "UPDATE shifts SET "
	q += "name = $1, start_time = $2, "
	q += "end_time = $3, updated_at = $4 "
	q += " WHERE id = $5 RETURNING *"
	expectedQuery := regexp.QuoteMeta(q)
	suite.mock.ExpectQuery(expectedQuery).
		WithArgs(shift.Name, shift.StartTime,
			shift.EndTime, time.Now().Unix(), shift.ID).
		WillReturnRows(rows).
		WillReturnError(nil)
	res, err := suite.shiftRepo.Update(context.TODO(), shift)
	require.Nil(suite.T(), res)
	require.NotNil(suite.T(), err)
}

func (suite *shiftRepositoryTestSuite) TestShiftRepository_Delete_ExpectSuccess() {}
func (suite *shiftRepositoryTestSuite) TestShiftRepository_Delete_ExpectError()   {}

func (suite *shiftRepositoryTestSuite) TestShiftRepository_Open_ExpectSuccess() {
	rows := suite.mock.
		NewRows([]string{"id"}).
		AddRow(1)
	shift := &model.StoreShiftForm{
		ShiftID: 1,
		UserID:  1,
		Cash:    200000,
	}
	q := "INSERT INTO store_shifts "
	q += "(shift_id, open_at, open_by, open_cash, created_at) "
	q += " VALUES ($1, $2, $3, $4, $5) RETURNING id"
	expectedQuery := regexp.QuoteMeta(q)
	suite.mock.ExpectQuery(expectedQuery).
		WithArgs(shift.ShiftID, time.Now().Unix(),
			shift.UserID, shift.Cash,
			time.Now().Unix()).
		WillReturnRows(rows).
		WillReturnError(nil)
	err := suite.shiftRepo.OpenShift(context.TODO(), shift)
	require.Nil(suite.T(), err)
}
func (suite *shiftRepositoryTestSuite) TestShiftRepository_Close_ExpectSuccess() {
	rows := suite.mock.
		NewRows([]string{"id"}).
		AddRow(1)
	shift := &model.StoreShiftForm{
		ID:      1,
		ShiftID: 1,
		UserID:  1,
		Cash:    200000,
	}
	q := "UPDATE store_shifts SET "
	q += "close_at = $1, close_by = $2, "
	q += "close_cash = $3, updated_at = $4 "
	q += " WHERE id = $5 AND shift_id = $6 RETURNING id"
	expectedQuery := regexp.QuoteMeta(q)
	suite.mock.ExpectQuery(expectedQuery).
		WithArgs(time.Now().Unix(),
			shift.UserID, shift.Cash,
			time.Now().Unix(),
			shift.ID, shift.ShiftID).
		WillReturnRows(rows).
		WillReturnError(nil)
	err := suite.shiftRepo.CloseShift(context.TODO(), shift)
	require.Nil(suite.T(), err)
}
