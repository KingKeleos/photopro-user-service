package user

import (
	"context"
	"log/slog"
	"time"

	"github.com/KingKeleos/photopro-user-service/database"
)

type Role struct {
	ID        uint64
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (r *Role) Create(ctx context.Context) error {
	query := `insert into roles (name, created_at, updated_at) values ($1, $2, $3) returning index`
	fields := []any{r.Name, time.Now(), time.Now()}
	return database.PGClient.QueryRowContext(ctx, query, fields...).Scan(&r.ID)
}

func (r *Role) List(ctx context.Context) ([]Role, error) {
	var roles []Role
	query := `select * from roles`
	rows, err := database.PGClient.QueryContext(ctx, query)
	if err != nil {
		slog.Error("fetching roles", "error", err)
		return nil, err
	}
	for rows.Next() {
		role := Role{}
		err := rows.Scan(&role.ID, &role.Name, &role.CreatedAt, &role.UpdatedAt)
		if err != nil {
			slog.Error("fetching role", "error", err)
		}
		roles = append(roles, role)
	}
	return roles, err
}

func (r *Role) Update(ctx context.Context) error {
	query := `update roles set name=$1 where index=$2`
	fields := []any{r.Name, r.ID}
	_, err := database.PGClient.ExecContext(ctx, query, fields...)
	return err
}
