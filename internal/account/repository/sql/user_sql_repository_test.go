package sql_test

import (
	"context"
	"errors"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/aasumitro/posbe/config"
	repoSql "github.com/aasumitro/posbe/internal/account/repository/sql"
	"github.com/aasumitro/posbe/pkg/model"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type userRepositoryTestSuite struct {
	suite.Suite
	mock     sqlmock.Sqlmock
	userRepo model.ICRUDRepository[model.User]
}

func (suite *userRepositoryTestSuite) SetupSuite() {
	var err error
	config.PostgresPool, suite.mock, err = sqlmock.New(
		sqlmock.QueryMatcherOption(
			sqlmock.QueryMatcherRegexp))
	require.NoError(suite.T(), err)
	suite.userRepo = repoSql.NewUserSQLRepository()
}

func (suite *userRepositoryTestSuite) AfterTest(_, _ string) {
	require.NoError(suite.T(), suite.mock.ExpectationsWereMet())
}

func (suite *userRepositoryTestSuite) TestUserRepository_All_ExpectedReturnDataRows() {
	users := suite.mock.
		NewRows([]string{"id", "users.role_id", "name", "username", "email", "phone", "role_id", "role_name", "role_description"}).
		AddRow(1, 1, "lorem ipsum", "lorem", "lorem@ipsum.id", "+6275555", 1, "test", "test 12345").
		AddRow(2, 2, "ipsum lorem", "ipsum", "ipsum@lorem.id", "+6278888", 1, "test", "test 12345")
	q := "SELECT u.id, u.role_id, u.name, u.username, u.email, u.phone, "
	q += "r.id as role_id, r.name as role_name, r.description FROM users as u "
	q += "JOIN roles as r ON r.id = u.role_id"
	expectedQuery := regexp.QuoteMeta(q)
	suite.mock.ExpectQuery(expectedQuery).WillReturnRows(users)
	res, err := suite.userRepo.All(context.TODO())
	require.Nil(suite.T(), err)
	require.NoError(suite.T(), err)
	require.NotNil(suite.T(), res)
}

func (suite *userRepositoryTestSuite) TestUserRepository_All_ExpectedReturnErrorFromQuery() {
	q := "SELECT u.id, u.role_id, u.name, u.username, u.email, u.phone, "
	q += "r.id as role_id, r.name as role_name, r.description FROM users as u "
	q += "JOIN roles as r ON r.id = u.role_id"
	expectedQuery := regexp.QuoteMeta(q)
	suite.mock.ExpectQuery(expectedQuery).WillReturnError(errors.New(""))
	res, err := suite.userRepo.All(context.TODO())
	require.NotNil(suite.T(), err)
	require.Nil(suite.T(), res)
}

func (suite *userRepositoryTestSuite) TestUserRepository_All_ExpectedReturnErrorFromScan() {
	users := suite.mock.
		NewRows([]string{"id", "users.role_id", "name", "username", "email", "phone", "role_id", "role_name", "role_description"}).
		AddRow(1, 1, "lorem ipsum", "lorem", "lorem@ipsum.id", "+6275555", 1, "test", "test 12345").
		AddRow(nil, nil, nil, nil, nil, nil, nil, nil, nil)
	q := "SELECT u.id, u.role_id, u.name, u.username, u.email, u.phone, "
	q += "r.id as role_id, r.name as role_name, r.description FROM users as u "
	q += "JOIN roles as r ON r.id = u.role_id"
	expectedQuery := regexp.QuoteMeta(q)
	suite.mock.ExpectQuery(expectedQuery).WillReturnRows(users)
	res, err := suite.userRepo.All(context.TODO())
	require.Nil(suite.T(), res)
	require.NotNil(suite.T(), err)
}

func (suite *userRepositoryTestSuite) TestUserRepository_Find_ExpectedSuccess() {
	user := suite.mock.
		NewRows([]string{"id", "users.role_id", "name", "username", "email", "phone", "password", "role_id", "role_name", "role_description"}).
		AddRow(1, 1, "lorem ipsum", "lorem", "lorem@ipsum.id", "+6275555", "qwe123", 1, "test", "test 12345")
	q := "SELECT u.id, u.role_id, u.name, u.username, u.email, u.phone, u.password, "
	q += "r.id as role_id, r.name as role_name, r.description FROM users as u "
	q += "JOIN roles as r ON r.id = u.role_id WHERE u.id = $1"
	expectedQuery := regexp.QuoteMeta(q)
	suite.mock.ExpectQuery(expectedQuery).WillReturnRows(user)
	res, err := suite.userRepo.Find(context.TODO(), model.FindWithID, 1)
	require.Nil(suite.T(), err)
	require.NoError(suite.T(), err)
	require.NotNil(suite.T(), res)
}

func (suite *userRepositoryTestSuite) TestUserRepository_FindBY_ExpectedSuccess() {
	tests := []struct {
		name  string
		args  string
		key   model.FindWith
		value string
	}{
		{
			name:  "test find with name",
			args:  "u.name LIKE %$1%",
			key:   model.FindWithName,
			value: "lorem",
		},
		{
			name:  "test find with username",
			args:  "u.username = $1",
			key:   model.FindWithUsername,
			value: "lorem",
		},
		{
			name:  "test find with email",
			args:  "u.email = $1",
			key:   model.FindWithEmail,
			value: "lorem@ipsum.id",
		},
		{
			name:  "test find with phone",
			args:  "u.phone = $1",
			key:   model.FindWithPhone,
			value: "+6275555",
		},
	}
	for _, tt := range tests {
		user := suite.mock.
			NewRows([]string{"id", "users.role_id", "name", "username", "email", "phone", "password", "role_id", "role_name", "role_description"}).
			AddRow(1, 1, "lorem ipsum", "lorem", "lorem@ipsum.id", "+6275555", "qwe123", 1, "test", "test 12345")
		q := "SELECT u.id, u.role_id, u.name, u.username, u.email, u.phone, u.password, "
		q += "r.id as role_id, r.name as role_name, r.description FROM users as u "
		q += "JOIN roles as r ON r.id = u.role_id WHERE "
		q += tt.args
		q += " LIMIT 1"
		expectedQuery := regexp.QuoteMeta(q)
		suite.mock.ExpectQuery(expectedQuery).WillReturnRows(user)
		res, err := suite.userRepo.Find(context.TODO(), tt.key, tt.value)
		require.Nil(suite.T(), err)
		require.NoError(suite.T(), err)
		require.NotNil(suite.T(), res)
	}
}

func (suite *userRepositoryTestSuite) TestUserRepository_Find_ExpectedError() {
	user := suite.mock.
		NewRows([]string{"id", "users.role_id", "name", "username", "email", "phone", "password", "role_id", "role_name", "role_description"}).
		AddRow(nil, nil, nil, nil, nil, nil, nil, nil, nil, nil)
	q := "SELECT u.id, u.role_id, u.name, u.username, u.email, u.phone, u.password, "
	q += "r.id as role_id, r.name as role_name, r.description FROM users as u "
	q += "JOIN roles as r ON r.id = u.role_id WHERE u.id = $1"
	expectedQuery := regexp.QuoteMeta(q)
	suite.mock.ExpectQuery(expectedQuery).WillReturnRows(user)
	res, err := suite.userRepo.Find(context.TODO(), model.FindWithID, 1)
	require.Nil(suite.T(), res)
	require.NotNil(suite.T(), err)
}

func (suite *userRepositoryTestSuite) TestUserRepository_Create_ExpectedSuccess() {
	user := &model.User{ID: 1, RoleID: 1, Name: "test 123", Username: "test", Email: "test@test.id", Phone: "+627888", Password: "12345"}
	rows := suite.mock.
		NewRows([]string{"id", "users.role_id", "name", "username", "email", "phone", "password", "role_id", "role_name", "role_description"}).
		AddRow(1, 1, "lorem ipsum", "lorem", "lorem@ipsum.id", "+6275555", "qwe123", 1, "test", "test 12345")
	q := "WITH u AS (INSERT INTO users(role_id, name, username, email, phone, password) "
	q += "values ($1, $2, $3, $4, $5, $6) RETURNING *) "
	q += "SELECT u.id, u.role_id, u.name, u.username, u.email, u.phone, u.password, "
	q += "r.id as role_id, r.name as role_name, r.description FROM u "
	q += "JOIN roles as r ON r.id = u.role_id"
	expectedQuery := regexp.QuoteMeta(q)
	suite.mock.ExpectQuery(expectedQuery).
		WithArgs(user.RoleID, user.Name, user.Username, user.Email, user.Phone, user.Password).
		WillReturnRows(rows).WillReturnError(nil)
	res, err := suite.userRepo.Create(context.TODO(), user)
	require.Nil(suite.T(), err)
	require.NotNil(suite.T(), res)
}

func (suite *userRepositoryTestSuite) TestUserRepository_Create_ExpectedError() {
	user := &model.User{ID: 1, RoleID: 1, Name: "test 123", Username: "test", Email: "test@test.id", Phone: "+627888", Password: "12345"}
	rows := suite.mock.
		NewRows([]string{"id", "users.role_id", "name", "username", "email", "phone", "password", "role_id", "role_name", "role_description"}).
		AddRow(nil, nil, nil, nil, nil, nil, nil, nil, nil, nil)
	q := "WITH u AS (INSERT INTO users(role_id, name, username, email, phone, password) "
	q += "values ($1, $2, $3, $4, $5, $6) RETURNING *) "
	q += "SELECT u.id, u.role_id, u.name, u.username, u.email, u.phone, u.password, "
	q += "r.id as role_id, r.name as role_name, r.description FROM u "
	q += "JOIN roles as r ON r.id = u.role_id"
	expectedQuery := regexp.QuoteMeta(q)
	suite.mock.ExpectQuery(expectedQuery).
		WithArgs(user.RoleID, user.Name, user.Username, user.Email, user.Phone, user.Password).
		WillReturnRows(rows).
		WillReturnError(nil)
	res, err := suite.userRepo.Create(context.TODO(), user)
	require.Nil(suite.T(), res)
	require.NotNil(suite.T(), err)
}

func (suite *userRepositoryTestSuite) TestUserRepository_Update_ExpectedSuccess() {
	user := &model.User{ID: 1, RoleID: 1, Name: "test 123", Username: "test", Email: "test@test.id", Phone: "+627888", Password: "12345"}
	rows := suite.mock.
		NewRows([]string{"id", "users.role_id", "name", "username", "email", "phone", "password", "role_id", "role_name", "role_description"}).
		AddRow(1, 1, "lorem ipsum", "lorem", "lorem@ipsum.id", "+6275555", "qwe123", 1, "test", "test 12345")
	q := "WITH u AS (UPDATE users SET role_id = $1, name = $2, username = $3, email = $4, "
	q += "phone = $5, password = $6 WHERE id = $7 RETURNING *) "
	q += "SELECT u.id, u.role_id, u.name, u.username, u.email, u.phone, u.password, "
	q += "r.id as role_id, r.name as role_name, r.description FROM u "
	q += "JOIN roles as r ON r.id = u.role_id"
	expectedQuery := regexp.QuoteMeta(q)
	suite.mock.ExpectQuery(expectedQuery).
		WithArgs(user.RoleID, user.Name, user.Username, user.Email, user.Phone, user.Password, user.ID).
		WillReturnRows(rows).WillReturnError(nil)
	res, err := suite.userRepo.Update(context.TODO(), user)
	require.Nil(suite.T(), err)
	require.NotNil(suite.T(), res)
}

func (suite *userRepositoryTestSuite) TestUserRepository_Update_ExpectedError() {
	user := &model.User{ID: 1, RoleID: 1, Name: "test 123", Username: "test", Email: "test@test.id", Phone: "+627888", Password: "12345"}
	rows := suite.mock.
		NewRows([]string{"id", "users.role_id", "name", "username", "email", "phone", "password", "role_id", "role_name", "role_description"}).
		AddRow(nil, nil, nil, nil, nil, nil, nil, nil, nil, nil)
	q := "WITH u AS (UPDATE users SET role_id = $1, name = $2, username = $3, email = $4, "
	q += "phone = $5, password = $6 WHERE id = $7 RETURNING *) "
	q += "SELECT u.id, u.role_id, u.name, u.username, u.email, u.phone, u.password, "
	q += "r.id as role_id, r.name as role_name, r.description FROM u "
	q += "JOIN roles as r ON r.id = u.role_id"
	expectedQuery := regexp.QuoteMeta(q)
	suite.mock.ExpectQuery(expectedQuery).
		WithArgs(user.RoleID, user.Name, user.Username, user.Email, user.Phone, user.Password, user.ID).
		WillReturnRows(rows).WillReturnError(nil)
	res, err := suite.userRepo.Update(context.TODO(), user)
	require.Nil(suite.T(), res)
	require.NotNil(suite.T(), err)
}

func (suite *userRepositoryTestSuite) TestUserRepository_Delete_ExpectedSuccess() {
	expectedQuery := regexp.QuoteMeta("DELETE FROM users")
	suite.mock.ExpectExec(expectedQuery).
		WithArgs(1).
		WillReturnResult(sqlmock.NewResult(0, 1))
	user := &model.User{ID: 1, RoleID: 1, Username: "test", Password: "12345"}
	err := suite.userRepo.Delete(context.TODO(), user)
	require.Nil(suite.T(), err)
}

func TestUserRepository(t *testing.T) {
	suite.Run(t, new(userRepositoryTestSuite))
}
