package sql

import (
	"context"
	"database/sql"
	"github.com/aasumitro/posbe/domain"
	"github.com/aasumitro/posbe/pkg/config"
	"time"
)

type RoomSQLRepository struct {
	Db *sql.DB
}

func (repo RoomSQLRepository) AllWhere(ctx context.Context, key domain.FindWith, val any) (rooms []*domain.Room, err error) {
	q := "SELECT * FROM rooms "
	if key == domain.FindWithRelationId {
		q += "WHERE floor_id = $1"
	}
	rows, err := repo.Db.QueryContext(ctx, q, val)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var room domain.Room

		if err := rows.Scan(
			&room.ID, &room.FloorId, &room.Name,
			&room.XPos, &room.YPos, &room.WSize,
			&room.HSize, &room.Capacity, &room.Price,
			&room.CreatedAt, &room.UpdatedAt,
		); err != nil {
			return nil, err
		}

		rooms = append(rooms, &room)
	}

	return rooms, nil
}

func (repo RoomSQLRepository) All(ctx context.Context) (rooms []*domain.Room, err error) {
	q := "SELECT * FROM rooms"
	rows, err := repo.Db.QueryContext(ctx, q)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var room domain.Room

		if err := rows.Scan(
			&room.ID, &room.FloorId, &room.Name,
			&room.XPos, &room.YPos, &room.WSize,
			&room.HSize, &room.Capacity, &room.Price,
			&room.CreatedAt, &room.UpdatedAt,
		); err != nil {
			return nil, err
		}

		rooms = append(rooms, &room)
	}

	return rooms, nil
}

func (repo RoomSQLRepository) Find(ctx context.Context, key domain.FindWith, val any) (room *domain.Room, err error) {
	q := "SELECT * FROM rooms WHERE "
	switch key {
	case domain.FindWithId:
		q += "id = $1 "
	case domain.FindWithRelationId:
		q += "floor_id = $1 "
	}
	q += "LIMIT 1"
	row := repo.Db.QueryRowContext(ctx, q, val)

	room = &domain.Room{}
	if err := row.Scan(
		&room.ID, &room.FloorId, &room.Name,
		&room.XPos, &room.YPos, &room.WSize,
		&room.HSize, &room.Capacity, &room.Price,
		&room.CreatedAt, &room.UpdatedAt,
	); err != nil {
		return nil, err
	}

	return room, nil
}

func (repo RoomSQLRepository) Create(ctx context.Context, params *domain.Room) (room *domain.Room, err error) {
	q := "INSERT INTO rooms (floor_id, name, x_pos,  "
	q += "y_pos, w_size, h_size, capacity, price, created_at) "
	q += "values ($1, $2, $3, $4, $5, $6, $7, $8, $9) RETURNING *"
	row := repo.Db.QueryRowContext(ctx, q,
		params.FloorId, params.Name, params.XPos,
		params.YPos, params.WSize, params.HSize,
		params.Capacity, params.Price, time.Now().Unix())

	room = &domain.Room{}
	if err := row.Scan(&room.ID, &room.FloorId, &room.Name,
		&room.XPos, &room.YPos, &room.WSize,
		&room.HSize, &room.Capacity, &room.Price,
		&room.CreatedAt, &room.UpdatedAt,
	); err != nil {
		return nil, err
	}

	return room, nil
}

func (repo RoomSQLRepository) Update(ctx context.Context, params *domain.Room) (room *domain.Room, err error) {
	q := "UPDATE rooms SET "
	q += "floor_id = $1, name = $2, x_pos = $3, "
	q += "y_pos = $4, w_size = $5, h_size = $6, "
	q += "capacity= $7, price = $8, updated_at = $9 "
	q += "WHERE id = $10 RETURNING *"
	row := repo.Db.QueryRowContext(ctx, q,
		params.FloorId, params.Name, params.XPos,
		params.YPos, params.WSize, params.HSize,
		params.Capacity, params.Price, time.Now().Unix(), params.ID)

	room = &domain.Room{}
	if err := row.Scan(&room.ID, &room.FloorId, &room.Name,
		&room.XPos, &room.YPos, &room.WSize,
		&room.HSize, &room.Capacity, &room.Price,
		&room.CreatedAt, &room.UpdatedAt,
	); err != nil {
		return nil, err
	}

	return room, nil
}

func (repo RoomSQLRepository) Delete(ctx context.Context, params *domain.Room) error {
	q := "DELETE FROM rooms WHERE id = $1"
	_, err := repo.Db.ExecContext(ctx, q, params.ID)
	return err
}

func NewRoomSQLRepository() domain.ICRUDAddOnRepository[domain.Room] {
	return &RoomSQLRepository{Db: config.DbPool}
}
