package db

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Metric struct {
	UserID      string
	AnalyteCode string
	Value       float64
	Time        time.Time
}

type Trend struct {
	UserID      string
	AnalyteCode string
	Value       float64
}

type Repository struct {
	pool *pgxpool.Pool
}

func NewRepository(pool *pgxpool.Pool) *Repository {
	return &Repository{pool: pool}
}

func (r *Repository) InsertMetric(ctx context.Context, m Metric) error {
	_, err := r.pool.Exec(ctx, `
		INSERT INTO blood_test_metrics (time, user_ref, analyte_code, value)
		VALUES ($1, $2, $3, $4)
	`, m.Time, m.UserID, m.AnalyteCode, m.Value)
	return err
}

func (r *Repository) UpsertTrend(ctx context.Context, t Trend) error {
	_, err := r.pool.Exec(ctx, `
		INSERT INTO blood_test_trends (user_ref, analyte_code, last_value, updated_at)
		VALUES ($1, $2, $3, now())
		ON CONFLICT (user_ref, analyte_code) DO UPDATE
			SET prev_value = blood_test_trends.last_value,
				last_value = EXCLUDED.last_value,
				delta = EXCLUDED.last_value - blood_test_trends.last_value,
				updated_at = now()
	`, t.UserID, t.AnalyteCode, t.Value)
	return err
}
