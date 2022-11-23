package sql

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/aasumitro/posbe/domain"
	"github.com/aasumitro/posbe/pkg/sql/query"
)

//goland:noinspection ALL
var (
	table = "roles"
)

type RoleSQLRepository struct {
	Db *sql.DB
}

// All
// repo.All(context)
//
// @usage
// all(context)
func (repo RoleSQLRepository) All(ctx context.Context) (roles []domain.Role, err error) {
	q := query.SQLSelectBuilder{}.Table(table).Build()
	fmt.Println(q)
	rows, err := repo.Db.QueryContext(ctx, q)
	if err != nil {
		return nil, err
	}

	defer func() {
		err := rows.Close()
		if err != nil {
			fmt.Println(err.Error())
		}
	}()

	for rows.Next() {
		var role domain.Role

		if err := rows.Scan(&role.ID, &role.Name, &role.Description); err != nil {
			return nil, err
		}

		roles = append(roles, role)
	}

	return roles, nil
}

// Find
// repo.Find(context, key, value)
//
// @variable
// var where = []string
//
// @usage
// find(context, ["id = 1"]) - search one value
// find(context, ["id 1"]) - search by id value int
// find(context, nil, "lorem") - key nil and value string = search by name
// find(context, nil, 1) - key nil and value int = search by id
func (repo RoleSQLRepository) Find(ctx context.Context, where string) (role *domain.Role, err error) {
	q := query.SQLSelectBuilder{}.Table(table).Where(where).Build()
	fmt.Println(q)
	row := repo.Db.QueryRowContext(ctx, q)

	var result domain.Role
	if err := row.Scan(&result.ID, &result.Name, &result.Description); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("row does not exist")
		}

		return nil, err
	}

	return &result, nil
}

//func (repo RoleSQLRepository) Create(ctx context.Context) {
//	//TODO implement me
//	panic("implement me")
//}
//
//func (repo RoleSQLRepository) Update(ctx context.Context) {
//	//TODO implement me
//	panic("implement me")
//}
//
//func (repo RoleSQLRepository) Delete(ctx context.Context) {
//	//TODO implement me
//	panic("implement me")
//}

func NewRoleSQlRepository(db *sql.DB) domain.RoleRepository {
	return &RoleSQLRepository{Db: db}
}
