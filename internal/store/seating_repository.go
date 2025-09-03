package store

import (
	"context"
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/aasumitro/posbe/internal/model"
	"github.com/jackc/pgx/v5"
)

type seatingRepository struct {
	db model.IPgxPool
}

func (repository seatingRepository) GetAllFloors(ctx context.Context) ([]*model.Floor, error) {
	q := `
		SELECT f.id, f.name,
		       (SELECT COUNT(*) FROM tables AS t WHERE t.floor_id = f.id) AS total_tables
		FROM floors AS f
	`
	rows, err := repository.db.Query(ctx, q)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var floors []*model.Floor
	var floor model.Floor
	if _, err = pgx.ForEachRow(rows, []any{
		&floor.ID, &floor.Name, &floor.TotalTables,
	}, func() error {
		f := floor
		floors = append(floors, &f)
		return nil
	}); err != nil {
		return nil, err
	}

	return floors, nil
}

// GetFloorByID
// Time and space complexity analysis:
//   - Query 1 (fetch floor): O(1), as it fetches a single row by primary key.
//   - Query 2 (fetch tables for floor): O(n), where n is the number of tables for the floor.
//   - Scanning rows into Go structs: O(n), proportional to number of tables.
//   - Total time complexity: O(n).
//   - Space complexity: O(n), as all tables are loaded into memory in a slice.
//   - Using SendBatch reduces network round trips (single round trip for both queries)
//     and avoids duplication/deduplication overhead that would occur with a SQL join.
func (repository seatingRepository) GetFloorByID(ctx context.Context, id int64) (*model.Floor, error) {
	batch := &pgx.Batch{}
	batch.Queue(`SELECT id, name FROM floors WHERE id = $1`, id)
	batch.Queue(`SELECT id, floor_id, name, 
		x_pos, y_pos, w_size, h_size, d_size, capacity, type		
		FROM tables WHERE floor_id = $1 ORDER BY created_at DESC
	`, id)

	br := repository.db.SendBatch(ctx, batch)
	defer func() { _ = br.Close() }()

	// Fetch floor
	var floor model.Floor
	if err := br.QueryRow().Scan(&floor.ID, &floor.Name); err != nil {
		return nil, err
	}

	// Fetch active shifts
	rows, err := br.Query()
	if err != nil {
		return nil, err
	}
	var tables []*model.Table
	for rows.Next() {
		var t model.Table
		if err := rows.Scan(
			&t.ID, &t.FloorID, &t.Name,
			&t.XPos, &t.YPos, &t.WSize, &t.HSize,
			&t.DSize, &t.Capacity, &t.Type,
		); err != nil {
			rows.Close()
			return nil, err
		}
		tables = append(tables, &t)
	}
	rows.Close()
	floor.Tables = tables
	floor.TotalTables = len(tables)

	return &floor, nil
}

func (repository seatingRepository) CreateFloor(ctx context.Context, form *SeatingFloorRequest) error {
	q := "INSERT INTO floors (name, created_at) VALUES ($1, $2)"
	_, err := repository.db.Exec(ctx, q, form.Name, time.Now().Unix())
	return err
}

func (repository seatingRepository) UpdateFloor(ctx context.Context, form *SeatingFloorRequest) error {
	q := "UPDATE floors SET name = $1, updated_at = $2 WHERE id = $3"
	_, err := repository.db.Exec(ctx, q, form.Name, time.Now().Unix(), form.ID)
	return err
}

func (repository seatingRepository) DeleteFloor(ctx context.Context, id int64) error {
	tx, err := repository.db.Begin(ctx)
	if err != nil {
		return err
	}
	defer func() {
		if tx != nil {
			_ = tx.Rollback(ctx)
		}
	}()

	// First delete tables under this floor
	if _, err := tx.Exec(ctx, "DELETE FROM tables WHERE floor_id = $1", id); err != nil {
		return err
	}

	// Then delete the floor itself
	if _, err := tx.Exec(ctx, "DELETE FROM floors WHERE id = $1", id); err != nil {
		return err
	}

	// Commit transaction
	if err := tx.Commit(ctx); err != nil {
		return err
	}

	// set tx = nil so rollback is skipped
	tx = nil
	return nil
}

func (repository seatingRepository) GetAllTable(ctx context.Context, floorID int64) ([]*model.Table, error) {
	q := "SELECT id, floor_id, name, x_pos, y_pos, w_size, h_size, d_size, capacity, type FROM tables WHERE floor_id = $1 ORDER BY created_at DESC"

	rows, err := repository.db.Query(ctx, q, floorID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tables []*model.Table
	var table model.Table
	if _, err = pgx.ForEachRow(rows, []any{
		&table.ID, &table.FloorID, &table.Name,
		&table.XPos, &table.YPos, &table.WSize, &table.HSize,
		&table.DSize, &table.Capacity, &table.Type,
	}, func() error {
		t := table
		tables = append(tables, &t)
		return nil
	}); err != nil {
		return nil, err
	}

	return tables, nil
}

func (repository seatingRepository) GetTableByID(ctx context.Context, id int64) (*model.Table, error) {
	q := `SELECT id, floor_id, name, x_pos, 
		y_pos, w_size, h_size, d_size, capacity, 
		type FROM tables WHERE id = $1
	`
	var table model.Table
	if err := repository.db.QueryRow(ctx, q, id).Scan(
		&table.ID, &table.FloorID, &table.Name,
		&table.XPos, &table.YPos, &table.WSize, &table.HSize,
		&table.DSize, &table.Capacity, &table.Type,
	); err != nil {
		return nil, err
	}

	return &table, nil
}

func (repository seatingRepository) CreateTable(ctx context.Context, form *SeatingTableRequest) error {
	q := `
		INSERT INTO tables (
			floor_id, name, x_pos, y_pos,
			w_size, h_size, d_size,
			capacity, type, created_at
		) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
	`
	_, err := repository.db.Exec(ctx, q, form.FloorID, form.Name, form.XPos, form.YPos, form.WSize, form.HSize, form.DSize, form.Capacity, form.Type, time.Now().Unix())
	return err
}

func (repository seatingRepository) UpdateTable(ctx context.Context, form *SeatingTableRequest) error {
	if form.ID == 0 {
		return errors.New("missing table ID for update")
	}

	var setClauses []string
	var args []any
	argPos := 1
	if form.FloorID > 0 {
		setClauses = append(setClauses, "floor_id = $"+strconv.Itoa(argPos))
		args = append(args, form.FloorID)
		argPos++
	}
	if form.Name != "" {
		setClauses = append(setClauses, "name = $"+strconv.Itoa(argPos))
		args = append(args, form.Name)
		argPos++
	}
	if form.XPos != nil {
		setClauses = append(setClauses, "x_pos = $"+strconv.Itoa(argPos))
		args = append(args, form.XPos)
		argPos++
	}
	if form.YPos != nil {
		setClauses = append(setClauses, "y_pos = $"+strconv.Itoa(argPos))
		args = append(args, form.YPos)
		argPos++
	}
	if form.WSize != nil {
		setClauses = append(setClauses, "w_size = $"+strconv.Itoa(argPos))
		args = append(args, form.WSize)
		argPos++
	}
	if form.HSize != nil {
		setClauses = append(setClauses, "h_size = $"+strconv.Itoa(argPos))
		args = append(args, form.HSize)
		argPos++
	}
	if form.DSize != nil {
		setClauses = append(setClauses, "d_size = $"+strconv.Itoa(argPos))
		args = append(args, form.DSize)
		argPos++
	}
	if form.Capacity > 0 {
		setClauses = append(setClauses, "capacity = $"+strconv.Itoa(argPos))
		args = append(args, form.Capacity)
		argPos++
	}
	if form.Type != "" {
		setClauses = append(setClauses, "type = $"+strconv.Itoa(argPos))
		args = append(args, form.Type)
		argPos++
	}
	if len(setClauses) == 0 {
		return errors.New("no fields to update")
	}
	setClauses = append(setClauses, "updated_at = $"+strconv.Itoa(argPos))
	args = append(args, time.Now().Unix())
	argPos++
	args = append(args, form.ID) // Add WHERE clause

	q := fmt.Sprintf(`UPDATE tables SET %s WHERE id = $%d`,
		strings.Join(setClauses, ", "), argPos)
	_, err := repository.db.Exec(ctx, q, args...)
	return err
}

func (repository seatingRepository) DeleteTable(ctx context.Context, id int64) error {
	q := "DELETE FROM tables WHERE id = $1"
	_, err := repository.db.Exec(ctx, q, id)
	return err
}

func NewSeatingRepository(db model.IPgxPool) IStoreSeatingRepository {
	return &seatingRepository{db: db}
}
