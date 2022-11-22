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

func (repo RoleSQLRepository) All(ctx context.Context) (roles []domain.Role, err error) {
	fmt.Println(ctx)
	// ===========

	q := query.SQLSelectBuilder{}.ForTable(table).Build()
	rows, err := repo.Db.Query(q)
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
// find(context, key, value)
//
// @variable
// var key = interface{} - accepted [string, int]
// var value = interface - accepted [string, int]
//
// @usage
// find(context, "name", "lorem") - search by name value string
// find(context, "id", 1) - search by id value int
// find(context, nil, "lorem") - key nil and value string = search by name
// find(context, nil, 1) - key nil and value int = search by id
func (repo RoleSQLRepository) Find(ctx context.Context, where []string) (role *domain.Role, err error) {
	// TODO:
	// 1. impl ctx for tr
	// 2. impl chain pattern for query
	fmt.Println(ctx)
	// ===========

	q := query.SQLSelectBuilder{}.
		ForTable(table).
		AddWhere(where).
		Build()
	row := repo.Db.QueryRow(q)

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
