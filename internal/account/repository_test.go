package account_test

import (
	"context"
	"errors"
	"regexp"
	"testing"
	"time"

	"github.com/aasumitro/posbe/internal/account"
	"github.com/aasumitro/posbe/internal/model"
	"github.com/pashagolub/pgxmock/v4"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type accountRepositoryTestSuite struct {
	suite.Suite
	mock        pgxmock.PgxPoolIface
	accountRepo account.IAccountRepository
}

// SetupSuite is useful in cases where the setup code is time-consuming and isn't modified in any of the tests.
// An example of when this could be useful is if you were testing code that reads from a database,
// and all the tests used the same data and only ran SELECT statements. In this scenario,
// SetupSuite could be used once to load the database with data.
func (suite *accountRepositoryTestSuite) SetupSuite() {
	var err error
	suite.mock, err = pgxmock.NewPool()
	require.NoError(suite.T(), err)
	defer suite.mock.Close()
	suite.accountRepo = account.NewAccountRepository(suite.mock)
}

func (suite *accountRepositoryTestSuite) AfterTest(_, _ string) {
	require.NoError(suite.T(), suite.mock.ExpectationsWereMet())
}

func (suite *accountRepositoryTestSuite) TestAccountRepository_RoleList_ExpectedReturnDataRows() {
	roles := suite.mock.
		NewRows([]string{"id", "name", "description", "usage"}).
		AddRow(1, "test", "test 1", 1).
		AddRow(2, "test 2", "test 2", 0)
	q := "SELECT roles.id, roles.name, roles.description, COUNT(users.role_id) as usage "
	q += "FROM roles LEFT OUTER JOIN users ON users.role_id = roles.id "
	q += "GROUP BY roles.id ORDER BY roles.id ASC"
	expectedQuery := regexp.QuoteMeta(q)
	suite.mock.ExpectQuery(expectedQuery).WillReturnRows(roles)
	res, err := suite.accountRepo.GetAllRoles(context.TODO())
	require.Nil(suite.T(), err)
	require.NoError(suite.T(), err)
	require.NotNil(suite.T(), res)
}
func (suite *accountRepositoryTestSuite) TestAccountRepository_RoleList_ExpectedReturnError() {
	q := "SELECT roles.id, roles.name, roles.description, COUNT(users.role_id) as usage "
	q += "FROM roles LEFT OUTER JOIN users ON users.role_id = roles.id "
	q += "GROUP BY roles.id ORDER BY roles.id ASC"
	expectedQuery := regexp.QuoteMeta(q)

	roles := suite.mock.
		NewRows([]string{"id", "name", "description", "usage"}).
		AddRow(1, "test", "test 1", 1).
		AddRow("bad_id", nil, nil, nil)

	suite.T().Run("error from query", func(t *testing.T) {
		suite.mock.ExpectQuery(expectedQuery).WillReturnError(errors.New(""))
		res, err := suite.accountRepo.GetAllRoles(context.TODO())
		require.NotNil(suite.T(), err)
		require.Nil(suite.T(), res)
	})

	suite.T().Run("error from scan", func(t *testing.T) {
		suite.mock.ExpectQuery(expectedQuery).WillReturnRows(roles)
		res, err := suite.accountRepo.GetAllRoles(context.TODO())
		require.Nil(suite.T(), res)
		require.NotNil(suite.T(), err)
	})
}

func (suite *accountRepositoryTestSuite) TestAccountRepository_UserList_ExpectedReturnDataRows() {
	users := suite.mock.
		NewRows([]string{"id", "users.role_id", "name", "username", "email", "role_id", "role_name", "role_description", "created_at"}).
		AddRow(1, 1, "lorem ipsum", "lorem", "lorem@ipsum.id", 1, "test", "test 12345", time.Now().Unix()).
		AddRow(2, 2, "ipsum lorem", "ipsum", "ipsum@lorem.id", 1, "test", "test 12345", time.Now().Unix())
	q := "SELECT u.id, u.role_id, u.name, u.username, u.email, "
	q += "r.id as role_id, r.name as role_name, r.description, u.created_at FROM users as u "
	q += "JOIN roles as r ON r.id = u.role_id"
	expectedQuery := regexp.QuoteMeta(q)
	suite.mock.ExpectQuery(expectedQuery).WillReturnRows(users)
	res, err := suite.accountRepo.GetAllUsers(context.TODO())
	require.Nil(suite.T(), err)
	require.NoError(suite.T(), err)
	require.NotNil(suite.T(), res)
}
func (suite *accountRepositoryTestSuite) TestAccountRepository_UserList_ExpectedReturnError() {
	q := "SELECT u.id, u.role_id, u.name, u.username, u.email, "
	q += "r.id as role_id, r.name as role_name, r.description, u.created_at FROM users as u "
	q += "JOIN roles as r ON r.id = u.role_id"
	expectedQuery := regexp.QuoteMeta(q)

	users := suite.mock.
		NewRows([]string{"id", "users.role_id", "name", "username", "email", "role_id", "role_name", "role_description", "created_at"}).
		AddRow(1, 1, "lorem ipsum", "lorem", "lorem@ipsum.id", 1, "test", "test 12345", time.Now().Unix()).
		AddRow("Bad_ID", nil, nil, nil, nil, nil, nil, nil, nil)

	suite.T().Run("error from query", func(t *testing.T) {
		suite.mock.ExpectQuery(expectedQuery).WillReturnError(errors.New(""))
		res, err := suite.accountRepo.GetAllUsers(context.TODO())
		require.NotNil(suite.T(), err)
		require.Nil(suite.T(), res)
	})

	suite.T().Run("error from scan", func(t *testing.T) {
		suite.mock.ExpectQuery(expectedQuery).WillReturnRows(users)
		res, err := suite.accountRepo.GetAllUsers(context.TODO())
		require.Nil(suite.T(), res)
		require.NotNil(suite.T(), err)
	})
}

func (suite *accountRepositoryTestSuite) TestAccountRepository_FindUser_ExpectedSuccess() {
	tests := []struct {
		name  string
		args  string
		key   model.FindWith
		value any
	}{
		{
			name:  "test find with username",
			args:  "u.id = $1",
			key:   model.FindWithID,
			value: 1,
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
	}
	for _, tt := range tests {
		user := suite.mock.
			NewRows([]string{"id", "users.role_id", "name", "username", "email", "password", "role_id", "role_name", "role_description"}).
			AddRow(1, 1, "lorem ipsum", "lorem", "lorem@ipsum.id", "qwe123", 1, "test", "test 12345")
		q := "SELECT u.id, u.role_id, u.name, u.username, u.email, u.password, "
		q += "r.id as role_id, r.name as role_name, r.description FROM users as u "
		q += "JOIN roles as r ON r.id = u.role_id WHERE "
		q += tt.args
		q += " LIMIT 1"
		expectedQuery := regexp.QuoteMeta(q)
		suite.mock.ExpectQuery(expectedQuery).WithArgs(tt.value).WillReturnRows(user)
		res, err := suite.accountRepo.FindUserBy(context.TODO(), tt.key, tt.value)
		require.Nil(suite.T(), err)
		require.NoError(suite.T(), err)
		require.NotNil(suite.T(), res)
	}
}
func (suite *accountRepositoryTestSuite) TestAccountRepository_FindUser_ExpectedError() {
	user := suite.mock.
		NewRows([]string{"id", "users.role_id", "name", "username", "email", "password", "role_id", "role_name", "role_description"}).
		AddRow("BAd_ID", nil, nil, nil, nil, nil, nil, nil, nil)

	q := "SELECT u.id, u.role_id, u.name, u.username, u.email, u.password, "
	q += "r.id as role_id, r.name as role_name, r.description FROM users as u "
	q += "JOIN roles as r ON r.id = u.role_id WHERE u.id = $1"
	expectedQuery := regexp.QuoteMeta(q)

	suite.T().Run("error from key", func(t *testing.T) {
		res, err := suite.accountRepo.FindUserBy(context.TODO(), model.FindWithSKU, 1)
		require.Nil(suite.T(), res)
		require.NotNil(suite.T(), err)
	})

	suite.T().Run("error from query", func(t *testing.T) {
		suite.mock.ExpectQuery(expectedQuery).WithArgs(1).WillReturnRows(user)
		res, err := suite.accountRepo.FindUserBy(context.TODO(), model.FindWithID, 1)
		require.Nil(suite.T(), res)
		require.NotNil(suite.T(), err)
	})
}

func (suite *accountRepositoryTestSuite) TestAccountRepository_CreateUser_ExpectedSuccess() {
	user := model.User{ID: 1, RoleID: 1, Name: "test 123", Username: "test", Email: "test@test.id", Password: "12345"}

	rows := suite.mock.
		NewRows([]string{"id", "users.role_id", "name", "username", "email", "role_id", "role_name", "role_description"}).
		AddRow(1, 1, "lorem ipsum", "lorem", "lorem@ipsum.id", 1, "test", "test 12345")

	q := "WITH u AS (INSERT INTO users(role_id, name, username, email, password, created_at) "
	q += "values ($1, $2, $3, $4, $5, $6) RETURNING *) "
	q += "SELECT u.id, u.role_id, u.name, u.username, u.email, "
	q += "r.id as role_id, r.name as role_name, r.description FROM u "
	q += "JOIN roles as r ON r.id = u.role_id"

	now := time.Now().Unix()
	expectedQuery := regexp.QuoteMeta(q)
	suite.mock.ExpectQuery(expectedQuery).
		WithArgs(user.RoleID, user.Name, user.Username,
			user.Email, user.Password, now).
		WillReturnRows(rows)
	res, err := suite.accountRepo.InsertUser(context.TODO(), user)
	require.Nil(suite.T(), err)
	require.NotNil(suite.T(), res)
}
func (suite *accountRepositoryTestSuite) TestAccountRepository_CreateUser_ExpectedError() {
	user := model.User{ID: 1, RoleID: 1, Name: "test 123", Username: "test", Email: "test@test.id", Password: "12345"}

	rows := suite.mock.
		NewRows([]string{"id", "users.role_id", "name", "username", "email", "role_id", "role_name", "role_description"}).
		AddRow("BadID", nil, nil, nil, nil, nil, nil, nil)

	q := "WITH u AS (INSERT INTO users(role_id, name, username, email, password, created_at) "
	q += "values ($1, $2, $3, $4, $5, $6) RETURNING *) "
	q += "SELECT u.id, u.role_id, u.name, u.username, u.email, "
	q += "r.id as role_id, r.name as role_name, r.description FROM u "
	q += "JOIN roles as r ON r.id = u.role_id"

	expectedQuery := regexp.QuoteMeta(q)
	now := time.Now().Unix()
	suite.mock.ExpectQuery(expectedQuery).
		WithArgs(user.RoleID, user.Name, user.Username,
			user.Email, user.Password, now).
		WillReturnRows(rows)
	res, err := suite.accountRepo.InsertUser(context.TODO(), user)
	require.Nil(suite.T(), res)
	require.NotNil(suite.T(), err)
}

func (suite *accountRepositoryTestSuite) TestAccountRepository_UpdateUser_ExpectedSuccess() {
	user := model.User{ID: 1, RoleID: 1, Name: "test 123", Username: "test", Email: "test@test.id", Password: "12345"}

	rows := suite.mock.
		NewRows([]string{"id", "users.role_id", "name", "username", "email", "password", "role_id", "role_name", "role_description"}).
		AddRow(1, 1, "lorem ipsum", "lorem", "lorem@ipsum.id", "secret", 1, "test", "test 12345")

	q := `
	WITH u AS (
	    UPDATE users SET role_id = $1, name = $2, username = $3, email = $4, password = $5, updated_at = $6
	    WHERE id = $7
	    RETURNING *
	)
	SELECT u.id, u.role_id, u.name, u.username, u.email, u.password,
	       r.id as role_id, r.name as role_name, r.description
	FROM u
	JOIN roles as r ON r.id = u.role_id
	`

	expectedQuery := regexp.QuoteMeta(q)
	suite.mock.ExpectQuery(expectedQuery).
		WithArgs(user.RoleID, user.Name, user.Username, user.Email, user.Password, time.Now().Unix(), user.ID).
		WillReturnRows(rows).WillReturnError(nil)
	res, err := suite.accountRepo.UpdateUserByID(context.TODO(), user)
	require.Nil(suite.T(), err)
	require.NotNil(suite.T(), res)
}
func (suite *accountRepositoryTestSuite) TestAccountRepository_UpdateUser_ExpectedError() {
	suite.T().Run("no user id", func(t *testing.T) {
		user := model.User{RoleID: 1, Name: "test 123", Username: "test", Email: "test@test.id", Password: "12345"}
		res, err := suite.accountRepo.UpdateUserByID(context.TODO(), user)
		require.Nil(suite.T(), res)
		require.NotNil(suite.T(), err)
	})

	suite.T().Run("no clauses", func(t *testing.T) {
		user := model.User{ID: 1}
		res, err := suite.accountRepo.UpdateUserByID(context.TODO(), user)
		require.Nil(suite.T(), res)
		require.NotNil(suite.T(), err)
	})

	suite.T().Run("error when insert", func(t *testing.T) {
		user := model.User{ID: 1, RoleID: 1, Name: "test 123", Username: "test", Email: "test@test.id", Password: "12345"}
		rows := suite.mock.
			NewRows([]string{"id", "users.role_id", "name", "username", "email", "password", "role_id", "role_name", "role_description"}).
			AddRow("BadID", nil, nil, nil, nil, nil, nil, nil, nil)
		q := `
	WITH u AS (
	    UPDATE users SET role_id = $1, name = $2, username = $3, email = $4, password = $5, updated_at = $6
	    WHERE id = $7
	    RETURNING *
	)
	SELECT u.id, u.role_id, u.name, u.username, u.email, u.password,
	       r.id as role_id, r.name as role_name, r.description
	FROM u
	JOIN roles as r ON r.id = u.role_id
	`
		expectedQuery := regexp.QuoteMeta(q)
		suite.mock.ExpectQuery(expectedQuery).
			WithArgs(user.RoleID, user.Name, user.Username, user.Email, user.Password, time.Now().Unix(), user.ID).
			WillReturnRows(rows)
		res, err := suite.accountRepo.UpdateUserByID(context.TODO(), user)
		require.Nil(suite.T(), res)
		require.NotNil(suite.T(), err)
	})
}

func (suite *accountRepositoryTestSuite) TestAccountRepository_DeleteUser_ExpectedSuccess() {
	expectedQuery := regexp.QuoteMeta("DELETE FROM users")
	suite.mock.ExpectExec(expectedQuery).
		WithArgs(int64(1)).
		WillReturnResult(pgxmock.NewResult("DELETE", 1))
	user := model.User{ID: 1, RoleID: 1, Username: "test", Password: "12345"}
	err := suite.accountRepo.DeleteUserByID(context.TODO(), user)
	require.Nil(suite.T(), err)
}

func TestRoleRepository(t *testing.T) {
	suite.Run(t, new(accountRepositoryTestSuite))
}
