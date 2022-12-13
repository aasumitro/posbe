package sql

import (
	"context"
	"database/sql"
	"github.com/aasumitro/posbe/domain"
	"github.com/aasumitro/posbe/pkg/config"
	"time"
)

type TableSQLRepository struct {
	Db *sql.DB
}

func (repo TableSQLRepository) AllWhere(ctx context.Context, key domain.FindWith, val any) (tables []*domain.Table, err error) {
	q := "SELECT * FROM tables "
	if key == domain.FindWithRelationId {
		q += "WHERE floor_id = $1"
	}
	rows, err := repo.Db.QueryContext(ctx, q, val)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var table domain.Table

		if err := rows.Scan(
			&table.ID, &table.FloorId, &table.Name,
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

func (repo TableSQLRepository) All(ctx context.Context) (tables []*domain.Table, err error) {
	q := "SELECT * FROM tables"
	rows, err := repo.Db.QueryContext(ctx, q)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var table domain.Table

		if err := rows.Scan(
			&table.ID, &table.FloorId, &table.Name,
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

func (repo TableSQLRepository) Find(ctx context.Context, key domain.FindWith, val any) (table *domain.Table, err error) {
	q := "SELECT * FROM tables WHERE "
	switch key {
	case domain.FindWithId:
		q += "id = $1 "
	case domain.FindWithRelationId:
		q += "floor_id = $1 "
	}
	q += "LIMIT 1"
	row := repo.Db.QueryRowContext(ctx, q, val)

	table = &domain.Table{}
	if err := row.Scan(
		&table.ID, &table.FloorId, &table.Name,
		&table.XPos, &table.YPos, &table.WSize,
		&table.HSize, &table.Capacity, &table.Type,
		&table.CreatedAt, &table.UpdatedAt,
	); err != nil {
		return nil, err
	}

	return table, nil
}

func (repo TableSQLRepository) Create(ctx context.Context, params *domain.Table) (table *domain.Table, err error) {
	q := "INSERT INTO tables (floor_id, name, x_pos,  "
	q += "y_pos, w_size, h_size, capacity, type, created_at) "
	q += "values ($1, $2, $3, $4, $5, $6, $7, $8, $9) RETURNING *"
	row := repo.Db.QueryRowContext(ctx, q,
		params.FloorId, params.Name, params.XPos,
		params.YPos, params.WSize, params.HSize,
		params.Capacity, params.Type, time.Now().Unix())

	table = &domain.Table{}
	if err := row.Scan(&table.ID, &table.FloorId, &table.Name,
		&table.XPos, &table.YPos, &table.WSize,
		&table.HSize, &table.Capacity, &table.Type,
		&table.CreatedAt, &table.UpdatedAt,
	); err != nil {
		return nil, err
	}

	return table, nil
}

func (repo TableSQLRepository) Update(ctx context.Context, params *domain.Table) (table *domain.Table, err error) {
	q := "UPDATE tables SET "
	q += "floor_id = $1, name = $2, x_pos = $3, "
	q += "y_pos = $4, w_size = $5, h_size = $6, "
	q += "capacity= $7, type = $8, updated_at = $9 "
	q += "WHERE id = $10 RETURNING *"
	row := repo.Db.QueryRowContext(ctx, q,
		params.FloorId, params.Name, params.XPos,
		params.YPos, params.WSize, params.HSize,
		params.Capacity, params.Type,
		time.Now().Unix(), params.ID)

	table = &domain.Table{}
	if err := row.Scan(&table.ID, &table.FloorId, &table.Name,
		&table.XPos, &table.YPos, &table.WSize,
		&table.HSize, &table.Capacity, &table.Type,
		&table.CreatedAt, &table.UpdatedAt,
	); err != nil {
		return nil, err
	}

	return table, nil
}

func (repo TableSQLRepository) Delete(ctx context.Context, params *domain.Table) error {
	q := "DELETE FROM tables WHERE id = $1"
	_, err := repo.Db.ExecContext(ctx, q, params.ID)
	return err
}

func NewTableSQLRepository() domain.ICRUDAddOnRepository[domain.Table] {
	return &TableSQLRepository{Db: config.DbPool}
}
