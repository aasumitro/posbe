package sql

import (
	"context"
	"database/sql"
	"time"

	"github.com/aasumitro/posbe/config"
	"github.com/aasumitro/posbe/pkg/model"
)

type FloorSQLRepository struct {
	Db *sql.DB
}

func (repo FloorSQLRepository) All(ctx context.Context) (floors []*model.Floor, err error) {
	q := "SELECT floors.id, floors.name, COUNT(tables.floor_id) "
	q += "as total_tables, COUNT(rooms.floor_id) as total_rooms, "
	q += "floors.created_at, floors.updated_at "
	q += "FROM floors LEFT OUTER JOIN tables ON tables.floor_id = floors.id "
	q += "LEFT OUTER JOIN rooms ON rooms.floor_id = floors.id "
	q += "GROUP BY floors.id ORDER BY floors.id ASC"
	rows, err := repo.Db.QueryContext(ctx, q)
	if err != nil {
		return nil, err
	}
	defer func(rows *sql.Rows) { _ = rows.Close() }(rows)
	for rows.Next() {
		var floor model.Floor
		if err := rows.Scan(
			&floor.ID, &floor.Name,
			&floor.TotalTables, &floor.TotalRooms,
			&floor.CreatedAt, &floor.UpdatedAt,
		); err != nil {
			return nil, err
		}
		floors = append(floors, &floor)
	}
	return floors, nil
}

func (repo FloorSQLRepository) Find(ctx context.Context, _ model.FindWith, val any) (floor *model.Floor, err error) {
	q := "SELECT floors.id, floors.name, COUNT(tables.floor_id) "
	q += "as total_tables, COUNT(rooms.floor_id) as total_rooms, "
	q += "floors.created_at, floors.updated_at "
	q += "FROM floors LEFT OUTER JOIN tables ON tables.floor_id = floors.id "
	q += "LEFT OUTER JOIN rooms ON rooms.floor_id = floors.id "
	q += "WHERE floors.id = $1 GROUP BY floors.id LIMIT 1"
	row := repo.Db.QueryRowContext(ctx, q, val)
	floor = &model.Floor{}
	if err := row.Scan(
		&floor.ID, &floor.Name,
		&floor.TotalTables, &floor.TotalRooms,
		&floor.CreatedAt, &floor.UpdatedAt,
	); err != nil {
		return nil, err
	}
	return floor, nil
}

func (repo FloorSQLRepository) Create(ctx context.Context, params *model.Floor) (floor *model.Floor, err error) {
	q := "INSERT INTO floors (name, created_at) values ($1, $2) RETURNING *"
	row := repo.Db.QueryRowContext(ctx, q, params.Name, time.Now().Unix())
	floor = &model.Floor{}
	if err := row.Scan(
		&floor.ID, &floor.Name,
		&floor.CreatedAt, &floor.UpdatedAt,
	); err != nil {
		return nil, err
	}
	return floor, nil
}

func (repo FloorSQLRepository) Update(ctx context.Context, params *model.Floor) (floor *model.Floor, err error) {
	q := "UPDATE floors SET name = $1, updated_at = $2 WHERE id = $3 RETURNING *"
	row := repo.Db.QueryRowContext(ctx, q, params.Name, time.Now().Unix(), params.ID)
	floor = &model.Floor{}
	if err := row.Scan(
		&floor.ID, &floor.Name,
		&floor.CreatedAt, &floor.UpdatedAt,
	); err != nil {
		return nil, err
	}
	return floor, nil
}

func (repo FloorSQLRepository) Delete(ctx context.Context, params *model.Floor) error {
	q := "DELETE FROM floors WHERE id = $1"
	_, err := repo.Db.ExecContext(ctx, q, params.ID)
	return err
}

func NewFloorSQLRepository() model.ICRUDRepository[model.Floor] {
	return &FloorSQLRepository{Db: config.PostgresPool}
}
