package sql_test

import (
	"context"
	"errors"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/aasumitro/posbe/domain"
	repoSql "github.com/aasumitro/posbe/internal/account/repository/sql"
	"github.com/aasumitro/posbe/pkg/config"
	_ "github.com/lib/pq"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"regexp"
	"testing"
)

type roleRepositoryTestSuite struct {
	suite.Suite
	mock     sqlmock.Sqlmock
	roleRepo domain.ICRUDRepository[domain.Role]
}

// SetupSuite is useful in cases where the setup code is time-consuming and isn't modified in any of the tests.
// An example of when this could be useful is if you were testing code that reads from a database,
// and all the tests used the same data and only ran SELECT statements. In this scenario,
// SetupSuite could be used once to load the database with data.
func (suite *roleRepositoryTestSuite) SetupSuite() {
	var (
		err error
	)

	config.Db, suite.mock, err = sqlmock.New(
		sqlmock.QueryMatcherOption(
			sqlmock.QueryMatcherRegexp))
	require.NoError(suite.T(), err)

	suite.roleRepo = repoSql.NewRoleSQlRepository()
}

func (suite *roleRepositoryTestSuite) AfterTest(_, _ string) {
	require.NoError(suite.T(), suite.mock.ExpectationsWereMet())
}

func (suite *roleRepositoryTestSuite) TestRoleRepository_All_ExpectedReturnDataRows() {
	roles := suite.mock.
		NewRows([]string{"id", "name", "description", "usage"}).
		AddRow(1, "test", "test 1", 1).
		AddRow(2, "test 2", "test 2", 0)
	q := "SELECT roles.id, roles.name, roles.description, COUNT(users.role_id) as usage "
	q += "FROM roles LEFT OUTER JOIN users ON users.role_id = roles.id "
	q += "GROUP BY roles.id ORDER BY roles.id ASC"
	expectedQuery := regexp.QuoteMeta(q)
	suite.mock.ExpectQuery(expectedQuery).WillReturnRows(roles)
	res, err := suite.roleRepo.All(context.TODO())
	require.Nil(suite.T(), err)
	require.NoError(suite.T(), err)
	require.NotNil(suite.T(), res)
}

func (suite *roleRepositoryTestSuite) TestRoleRepository_All_ExpectedReturnErrorFromQuery() {
	q := "SELECT roles.id, roles.name, roles.description, COUNT(users.role_id) as usage "
	q += "FROM roles LEFT OUTER JOIN users ON users.role_id = roles.id "
	q += "GROUP BY roles.id ORDER BY roles.id ASC"
	expectedQuery := regexp.QuoteMeta(q)
	suite.mock.ExpectQuery(expectedQuery).WillReturnError(errors.New(""))
	res, err := suite.roleRepo.All(context.TODO())
	require.NotNil(suite.T(), err)
	require.Nil(suite.T(), res)
}

func (suite *roleRepositoryTestSuite) TestRoleRepository_All_ExpectedReturnErrorFromScan() {
	roles := suite.mock.
		NewRows([]string{"id", "name", "description", "usage"}).
		AddRow(1, "test", "test 1", 1).
		AddRow(nil, nil, nil, nil)
	q := "SELECT roles.id, roles.name, roles.description, COUNT(users.role_id) as usage "
	q += "FROM roles LEFT OUTER JOIN users ON users.role_id = roles.id "
	q += "GROUP BY roles.id ORDER BY roles.id ASC"
	expectedQuery := regexp.QuoteMeta(q)
	suite.mock.ExpectQuery(expectedQuery).WillReturnRows(roles)
	res, err := suite.roleRepo.All(context.TODO())
	require.Nil(suite.T(), res)
	require.NotNil(suite.T(), err)
}

func (suite *roleRepositoryTestSuite) TestRoleRepository_Find_ExpectedSuccess() {
	role := suite.mock.
		NewRows([]string{"id", "name", "description", "usage"}).
		AddRow(1, "test", "test 1", 1)
	q := "SELECT roles.id, roles.name, roles.description, COUNT(users.role_id) as usage "
	q += "FROM roles LEFT OUTER JOIN users ON users.role_id = roles.id "
	q += "WHERE roles.id = $1 GROUP BY roles.id LIMIT 1"
	expectedQuery := regexp.QuoteMeta(q)
	suite.mock.ExpectQuery(expectedQuery).WillReturnRows(role)
	res, err := suite.roleRepo.Find(context.TODO(), domain.FindWithId, 1)
	require.Nil(suite.T(), err)
	require.NoError(suite.T(), err)
	require.NotNil(suite.T(), res)
}

func (suite *roleRepositoryTestSuite) TestRoleRepository_Find_ExpectedError() {
	role := suite.mock.
		NewRows([]string{"id", "name", "description", "usage"}).
		AddRow(nil, nil, nil, nil)
	q := "SELECT roles.id, roles.name, roles.description, COUNT(users.role_id) as usage "
	q += "FROM roles LEFT OUTER JOIN users ON users.role_id = roles.id "
	q += "WHERE roles.id = $1 GROUP BY roles.id LIMIT 1"
	expectedQuery := regexp.QuoteMeta(q)
	suite.mock.ExpectQuery(expectedQuery).WillReturnRows(role)
	res, err := suite.roleRepo.Find(context.TODO(), domain.FindWithId, 1)
	require.Nil(suite.T(), res)
	require.NotNil(suite.T(), err)
}

func (suite *roleRepositoryTestSuite) TestRoleRepository_Create_ExpectedSuccess() {
	role := &domain.Role{ID: 1, Name: "test", Description: "test"}
	rows := suite.mock.
		NewRows([]string{"id", "name", "description"}).
		AddRow(1, "test", "test 1")
	expectedQuery := regexp.QuoteMeta("INSERT INTO roles (name, description) values ($1, $2) RETURNING *")
	suite.mock.ExpectQuery(expectedQuery).
		WithArgs(role.Name, role.Description).
		WillReturnRows(rows).
		WillReturnError(nil)
	res, err := suite.roleRepo.Create(context.TODO(), role)
	require.Nil(suite.T(), err)
	require.NotNil(suite.T(), res)
}

func (suite *roleRepositoryTestSuite) TestRoleRepository_Create_ExpectedError() {
	role := &domain.Role{ID: 1, Name: "test", Description: "test"}
	rows := suite.mock.
		NewRows([]string{"id", "name", "description"}).
		AddRow(1, nil, nil)
	expectedQuery := regexp.QuoteMeta("INSERT INTO roles (name, description) values ($1, $2) RETURNING *")
	suite.mock.ExpectQuery(expectedQuery).
		WithArgs(role.Name, role.Description).
		WillReturnRows(rows).
		WillReturnError(nil)
	res, err := suite.roleRepo.Create(context.TODO(), role)
	require.Nil(suite.T(), res)
	require.NotNil(suite.T(), err)
}

func (suite *roleRepositoryTestSuite) TestRoleRepository_Update_ExpectedSuccess() {
	role := &domain.Role{ID: 1, Name: "test", Description: "test"}
	rows := suite.mock.
		NewRows([]string{"id", "name", "description"}).
		AddRow(1, "test", "test")
	expectedQuery := regexp.QuoteMeta("UPDATE roles SET name = $1, description = $2 WHERE id = $3 RETURNING *")
	suite.mock.ExpectQuery(expectedQuery).
		WithArgs(role.Name, role.Description, role.ID).
		WillReturnRows(rows).
		WillReturnError(nil)
	res, err := suite.roleRepo.Update(context.TODO(), role)
	require.Nil(suite.T(), err)
	require.NotNil(suite.T(), res)
}

func (suite *roleRepositoryTestSuite) TestRoleRepository_Update_ExpectedError() {
	role := &domain.Role{ID: 1, Name: "test", Description: "test"}
	rows := suite.mock.
		NewRows([]string{"id", "name", "description"}).
		AddRow(1, nil, nil)
	expectedQuery := regexp.QuoteMeta("UPDATE roles SET name = $1, description = $2 WHERE id = $3 RETURNING *")
	suite.mock.ExpectQuery(expectedQuery).
		WithArgs(role.Name, role.Description, role.ID).
		WillReturnRows(rows).
		WillReturnError(nil)
	res, err := suite.roleRepo.Update(context.TODO(), role)
	require.Nil(suite.T(), res)
	require.NotNil(suite.T(), err)
}

func (suite *roleRepositoryTestSuite) TestRoleRepository_Delete_ExpectedSuccess() {
	expectedQuery := regexp.QuoteMeta("DELETE FROM roles")
	suite.mock.ExpectExec(expectedQuery).
		WithArgs(1).
		WillReturnResult(sqlmock.NewResult(0, 1))
	role := &domain.Role{ID: 1, Name: "test", Description: "test"}
	err := suite.roleRepo.Delete(context.TODO(), role)
	require.Nil(suite.T(), err)
}

func TestRoleRepository(t *testing.T) {
	suite.Run(t, new(roleRepositoryTestSuite))
}
