package sql

import (
	"context"
	"database/sql"
	"time"

	"github.com/aasumitro/posbe/config"
	"github.com/aasumitro/posbe/pkg/model"
)

type RoomSQLRepository struct {
	Db *sql.DB
}

func (repo RoomSQLRepository) AllWhere(ctx context.Context, key model.FindWith, val any) (rooms []*model.Room, err error) {
	q := "SELECT * FROM rooms "
	if key == model.FindWithRelationID {
		q += "WHERE floor_id = $1"
	}
	rows, err := repo.Db.QueryContext(ctx, q, val)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var room model.Room
		if err := rows.Scan(
			&room.ID, &room.FloorID, &room.Name,
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

func (repo RoomSQLRepository) All(ctx context.Context) (rooms []*model.Room, err error) {
	q := "SELECT * FROM rooms"
	rows, err := repo.Db.QueryContext(ctx, q)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var room model.Room
		if err := rows.Scan(
			&room.ID, &room.FloorID, &room.Name,
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

func (repo RoomSQLRepository) Find(ctx context.Context, key model.FindWith, val any) (room *model.Room, err error) {
	q := "SELECT * FROM rooms WHERE "
	//goland:noinspection ALL
	switch key {
	case model.FindWithID:
		q += "id = $1 "
	case model.FindWithRelationID:
		q += "floor_id = $1 "
	}
	q += "LIMIT 1"
	row := repo.Db.QueryRowContext(ctx, q, val)
	room = &model.Room{}
	if err := row.Scan(
		&room.ID, &room.FloorID, &room.Name,
		&room.XPos, &room.YPos, &room.WSize,
		&room.HSize, &room.Capacity, &room.Price,
		&room.CreatedAt, &room.UpdatedAt,
	); err != nil {
		return nil, err
	}
	return room, nil
}

func (repo RoomSQLRepository) Create(ctx context.Context, params *model.Room) (room *model.Room, err error) {
	q := "INSERT INTO rooms (floor_id, name, x_pos,  "
	q += "y_pos, w_size, h_size, capacity, price, created_at) "
	q += "values ($1, $2, $3, $4, $5, $6, $7, $8, $9) RETURNING *"
	row := repo.Db.QueryRowContext(ctx, q,
		params.FloorID, params.Name, params.XPos,
		params.YPos, params.WSize, params.HSize,
		params.Capacity, params.Price, time.Now().Unix())
	room = &model.Room{}
	if err := row.Scan(&room.ID, &room.FloorID, &room.Name,
		&room.XPos, &room.YPos, &room.WSize,
		&room.HSize, &room.Capacity, &room.Price,
		&room.CreatedAt, &room.UpdatedAt,
	); err != nil {
		return nil, err
	}
	return room, nil
}

func (repo RoomSQLRepository) Update(ctx context.Context, params *model.Room) (room *model.Room, err error) {
	q := "UPDATE rooms SET "
	q += "floor_id = $1, name = $2, x_pos = $3, "
	q += "y_pos = $4, w_size = $5, h_size = $6, "
	q += "capacity= $7, price = $8, updated_at = $9 "
	q += "WHERE id = $10 RETURNING *"
	row := repo.Db.QueryRowContext(ctx, q,
		params.FloorID, params.Name, params.XPos,
		params.YPos, params.WSize, params.HSize,
		params.Capacity, params.Price, time.Now().Unix(), params.ID)
	room = &model.Room{}
	if err := row.Scan(&room.ID, &room.FloorID, &room.Name,
		&room.XPos, &room.YPos, &room.WSize,
		&room.HSize, &room.Capacity, &room.Price,
		&room.CreatedAt, &room.UpdatedAt,
	); err != nil {
		return nil, err
	}
	return room, nil
}

func (repo RoomSQLRepository) Delete(ctx context.Context, params *model.Room) error {
	q := "DELETE FROM rooms WHERE id = $1"
	_, err := repo.Db.ExecContext(ctx, q, params.ID)
	return err
}

func NewRoomSQLRepository() model.ICRUDAddOnRepository[model.Room] {
	return &RoomSQLRepository{Db: config.PostgresPool}
}
