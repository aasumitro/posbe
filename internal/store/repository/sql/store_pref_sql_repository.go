package sql

import (
	"context"
	"database/sql"
	"github.com/aasumitro/posbe/configs"
	"github.com/aasumitro/posbe/domain"
	"time"
)

type StorePrefSQLRepository struct {
	Db *sql.DB
}

func (repo StorePrefSQLRepository) Find(ctx context.Context, key string) (pref *domain.StoreSetting, err error) {
	q := "SELECT * FROM store_prefs WHERE key = $1 LIMIT 1"
	row := repo.Db.QueryRowContext(ctx, q, key)

	var storePref domain.StorePref
	if err := row.Scan(
		&storePref.Key, &storePref.Value,
		&storePref.CreatedAt, &storePref.UpdatedAt,
	); err != nil {
		return nil, err
	}

	return &domain.StoreSetting{
		storePref.Key: storePref.Value,
	}, err
}

func (repo StorePrefSQLRepository) All(ctx context.Context) (prefs *domain.StoreSetting, err error) {
	q := "SELECT * FROM store_prefs"
	rows, err := repo.Db.QueryContext(ctx, q)
	if err != nil {
		return nil, err
	}
	defer func(rows *sql.Rows) { _ = rows.Close() }(rows)

	var storeSetting = make(domain.StoreSetting)
	for rows.Next() {
		var storePref domain.StorePref

		if err := rows.Scan(
			&storePref.Key, &storePref.Value,
			&storePref.CreatedAt, &storePref.UpdatedAt,
		); err != nil {
			return nil, err
		}

		storeSetting[storePref.Key] = storePref.Value
	}

	return &storeSetting, nil
}

func (repo StorePrefSQLRepository) Update(ctx context.Context, key, value string) (prefs *domain.StoreSetting, err error) {
	q := "UPDATE store_prefs SET value = $1, updated_at = $2 WHERE key = $3 RETURNING *"
	row := repo.Db.QueryRowContext(ctx, q, value, time.Now().Unix(), key)

	var storePref domain.StorePref
	if err := row.Scan(
		&storePref.Key, &storePref.Value,
		&storePref.CreatedAt, &storePref.UpdatedAt,
	); err != nil {
		return nil, err
	}

	return &domain.StoreSetting{
		storePref.Key: storePref.Value,
	}, err
}

func NewStorePrefSQLRepository() domain.IStorePrefRepository {
	return &StorePrefSQLRepository{Db: configs.DbPool}
}
