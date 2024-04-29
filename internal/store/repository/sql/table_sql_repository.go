package sql

import (
	"context"
	"database/sql"
	"time"

	"github.com/aasumitro/posbe/config"
	"github.com/aasumitro/posbe/pkg/model"
)

type TableSQLRepository struct {
	Db *sql.DB
}

func (repo TableSQLRepository) AllWhere(ctx context.Context, key model.FindWith, val any) (tables []*model.Table, err error) {
	q := "SELECT * FROM tables "
	if key == model.FindWithRelationID {
		q += "WHERE floor_id = $1"
	}
	rows, err := repo.Db.QueryContext(ctx, q, val)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var table model.Table
		if err := rows.Scan(
			&table.ID, &table.FloorID, &table.Name,
			&table.XPos, &table.YPos, &table.WSize,
			&table.HSize, &table.Capacity, &table.Type,
			&table.CreatedAt, &table.UpdatedAt,
		); err != nil {
			return nil, err
		}
		tables = append(tables, &table)
	}
	return tables, nil
}

func (repo TableSQLRepository) All(ctx context.Context) (tables []*model.Table, err error) {
	q := "SELECT * FROM tables"
	rows, err := repo.Db.QueryContext(ctx, q)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var table model.Table
		if err := rows.Scan(
			&table.ID, &table.FloorID, &table.Name,
			&table.XPos, &table.YPos, &table.WSize,
			&table.HSize, &table.Capacity, &table.Type,
			&table.CreatedAt, &table.UpdatedAt,
		); err != nil {
			return nil, err
		}
		tables = append(tables, &table)
	}

	return tables, nil
}

func (repo TableSQLRepository) Find(ctx context.Context, key model.FindWith, val any) (table *model.Table, err error) {
	q := "SELECT * FROM tables WHERE "
	//goland:noinspection ALL
	switch key {
	case model.FindWithID:
		q += "id = $1 "
	case model.FindWithRelationID:
		q += "floor_id = $1 "
	}
	q += "LIMIT 1"
	row := repo.Db.QueryRowContext(ctx, q, val)
	table = &model.Table{}
	if err := row.Scan(
		&table.ID, &table.FloorID, &table.Name,
		&table.XPos, &table.YPos, &table.WSize,
		&table.HSize, &table.Capacity, &table.Type,
		&table.CreatedAt, &table.UpdatedAt,
	); err != nil {
		return nil, err
	}
	return table, nil
}

func (repo TableSQLRepository) Create(ctx context.Context, params *model.Table) (table *model.Table, err error) {
	q := "INSERT INTO tables (floor_id, name, x_pos,  "
	q += "y_pos, w_size, h_size, capacity, type, created_at) "
	q += "values ($1, $2, $3, $4, $5, $6, $7, $8, $9) RETURNING *"
	row := repo.Db.QueryRowContext(ctx, q,
		params.FloorID, params.Name, params.XPos,
		params.YPos, params.WSize, params.HSize,
		params.Capacity, params.Type, time.Now().Unix())
	table = &model.Table{}
	if err := row.Scan(&table.ID, &table.FloorID, &table.Name,
		&table.XPos, &table.YPos, &table.WSize,
		&table.HSize, &table.Capacity, &table.Type,
		&table.CreatedAt, &table.UpdatedAt,
	); err != nil {
		return nil, err
	}
	return table, nil
}

func (repo TableSQLRepository) Update(ctx context.Context, params *model.Table) (table *model.Table, err error) {
	q := "UPDATE tables SET "
	q += "floor_id = $1, name = $2, x_pos = $3, "
	q += "y_pos = $4, w_size = $5, h_size = $6, "
	q += "capacity= $7, type = $8, updated_at = $9 "
	q += "WHERE id = $10 RETURNING *"
	row := repo.Db.QueryRowContext(ctx, q,
		params.FloorID, params.Name, params.XPos,
		params.YPos, params.WSize, params.HSize,
		params.Capacity, params.Type,
		time.Now().Unix(), params.ID)
	table = &model.Table{}
	if err := row.Scan(&table.ID, &table.FloorID, &table.Name,
		&table.XPos, &table.YPos, &table.WSize,
		&table.HSize, &table.Capacity, &table.Type,
		&table.CreatedAt, &table.UpdatedAt,
	); err != nil {
		return nil, err
	}
	return table, nil
}

func (repo TableSQLRepository) Delete(ctx context.Context, params *model.Table) error {
	q := "DELETE FROM tables WHERE id = $1"
	_, err := repo.Db.ExecContext(ctx, q, params.ID)
	return err
}

func NewTableSQLRepository() model.ICRUDAddOnRepository[model.Table] {
	return &TableSQLRepository{Db: config.PostgresPool}
}
