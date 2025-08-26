package store

import (
	"context"
	"fmt"
	"strings"

	"github.com/aasumitro/posbe/internal/model"
	"github.com/jackc/pgx/v5"
)

type settingRepository struct {
	db model.IPgxPool
}

func (repository settingRepository) FindByKey(ctx context.Context, key string) (*model.StoreSetting, error) {
	q := "SELECT * FROM store_prefs WHERE key = $1 LIMIT 1"

	var storePref model.StorePref
	if err := repository.db.QueryRow(ctx, q, key).Scan(
		&storePref.Key, &storePref.Value,
		&storePref.CreatedAt, &storePref.UpdatedAt,
	); err != nil {
		return nil, err
	}

	return &model.StoreSetting{storePref.Key: storePref.Value}, nil
}

func (repository settingRepository) GetAll(ctx context.Context) (*model.StoreSetting, error) {
	q := "SELECT * FROM store_prefs"
	rows, err := repository.db.Query(ctx, q)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var storePref model.StorePref
	storeSetting := make(model.StoreSetting)
	if _, err = pgx.ForEachRow(rows, []any{
		&storePref.Key, &storePref.Value,
		&storePref.CreatedAt, &storePref.UpdatedAt,
	}, func() error {
		storeSetting[storePref.Key] = storePref.Value
		return nil
	}); err != nil {
		return nil, err
	}

	return &storeSetting, nil
}

func (repository settingRepository) Update(ctx context.Context, setting *model.StoreSetting) error {
	if len(*setting) == 0 {
		return nil
	}

	// prepare placeholders ($1,$2), ($3,$4), ...
	var args []interface{}
	var values []string
	i := 1
	for key, value := range *setting {
		f := "($%d, $%d, extract(epoch from now()), extract(epoch from now()))"
		values = append(values, fmt.Sprintf(f, i, i+1))
		args = append(args, key, value)
		i += 2
	}

	q := fmt.Sprintf(`
		INSERT INTO store_prefs (key, value, created_at, updated_at)
		VALUES %s
		ON CONFLICT (key) DO UPDATE
		SET value = EXCLUDED.value,
		    updated_at = extract(epoch from now())
	`, strings.Join(values, ", "))

	_, err := repository.db.Exec(ctx, q, args...)
	return err
}

func NewSettingRepository(db model.IPgxPool) IStoreSettingRepository {
	return &settingRepository{db: db}
}
