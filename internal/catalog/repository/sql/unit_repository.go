package sql

import (
	"context"
	"database/sql"
	"github.com/aasumitro/posbe/configs"
	"github.com/aasumitro/posbe/domain"
)

type UnitSQLRepository struct {
	Db *sql.DB
}

func (repo UnitSQLRepository) All(ctx context.Context) (data []*domain.Unit, err error) {
	q := "SELECT * FROM units"
	rows, err := repo.Db.QueryContext(ctx, q)
	if err != nil {
		return nil, err
	}
	defer func(rows *sql.Rows) { _ = rows.Close() }(rows)

	for rows.Next() {
		var unit domain.Unit

		if err := rows.Scan(
			&unit.ID, &unit.Magnitude,
			&unit.Name, &unit.Symbol,
		); err != nil {
			return nil, err
		}

		data = append(data, &unit)
	}

	return data, nil
}

func (repo UnitSQLRepository) Find(ctx context.Context, _ domain.FindWith, val any) (data *domain.Unit, err error) {
	q := "SELECT * FROM units WHERE id = $1 LIMIT 1"
	row := repo.Db.QueryRowContext(ctx, q, val)

	data = &domain.Unit{}
	if err := row.Scan(
		&data.ID, &data.Magnitude,
		&data.Name, &data.Symbol,
	); err != nil {
		return nil, err
	}

	return data, nil
}

func (repo UnitSQLRepository) Create(ctx context.Context, params *domain.Unit) (data *domain.Unit, err error) {
	q := "INSERT INTO units (magnitude, name, symbol) VALUES ($1, $2, $3) RETURNING *"
	row := repo.Db.QueryRowContext(ctx, q, params.Magnitude, params.Name, params.Symbol)

	data = &domain.Unit{}
	if err := row.Scan(
		&data.ID, &data.Magnitude,
		&data.Name, &data.Symbol,
	); err != nil {
		return nil, err
	}

	return data, nil
}

func (repo UnitSQLRepository) Update(ctx context.Context, params *domain.Unit) (data *domain.Unit, err error) {
	q := "UPDATE units SET magnitude = $1, name = $2, symbol = $3 WHERE id = $4 RETURNING *"
	row := repo.Db.QueryRowContext(ctx, q, params.Magnitude, params.Name, params.Symbol, params.ID)

	data = &domain.Unit{}
	if err := row.Scan(
		&data.ID, &data.Magnitude,
		&data.Name, &data.Symbol,
	); err != nil {
		return nil, err
	}

	return data, nil
}

func (repo UnitSQLRepository) Delete(ctx context.Context, params *domain.Unit) error {
	q := "DELETE FROM units WHERE id = $1"
	_, err := repo.Db.ExecContext(ctx, q, params.ID)
	return err
}

func NewUnitSQLRepository() domain.ICRUDRepository[domain.Unit] {
	return &UnitSQLRepository{Db: configs.DbPool}
}
