package user

import (
	"context"
	"database/sql"
	"errors"
	"log/slog"
	"time"

	"github.com/KingKeleos/photopro-user-service/database"
)

type User struct {
	ID        uint64
	Username  string
	EMail     string
	Password  string
	Location  Location
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Location struct {
	ID           int
	Name         string
	Street       string
	StreetNumber uint32
	Country      string
	FederalState string
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

func (u *User) CreateLocation(ctx context.Context) error {
	query := `insert into locations (name, street, street_number, country, federal_state, created_at, updated_at)
			values ($1, $2, $3, $4, $5, $6, $7) returning index`
	fields := []any{u.Location.Name, u.Location.Street, u.Location.StreetNumber, u.Location.Country, u.Location.FederalState, time.Now(), time.Now()}
	return database.PGClient.QueryRowContext(ctx, query, fields...).Scan(&u.Location.ID)
}

func (u *User) DeleteLocation(ctx context.Context) error {
	query := `select location_id from users where index = $1`
	fields := []any{u.ID}
	err := database.PGClient.QueryRowContext(ctx, query, fields...).Scan(&u.Location.ID)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return err
	}
	query = `delete from locations where index = $1`
	fields = []any{u.Location.ID}
	_, err = database.PGClient.ExecContext(ctx, query, fields...)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return err
	}
	return nil
}

// Create will write the data of the User into the Database
func (u *User) Create(ctx context.Context) error {
	err := u.CreateLocation(ctx)
	if err != nil {
		slog.Error("creating user location", "error", err)
	}
	query := `insert into users (username, email, password, location_id, created_at, updated_at)
			values ($1, $2, $3, $4, $5, $6) returning index`
	fields := []any{u.Username, u.EMail, u.Password, u.Location.ID, time.Now(), time.Now()}
	return database.PGClient.QueryRowContext(ctx, query, fields...).Scan(&u.ID)
}

func (u *User) Delete(ctx context.Context) error {
	err := u.DeleteLocation(ctx)
	if err != nil {
		slog.Error("deleting", "location", u.Location, "error", err)
		return err
	}
	nameQuery := `select username from users where index=$1`
	nameFields := []any{u.ID}
	err = database.PGClient.QueryRowContext(ctx, nameQuery, nameFields...).Scan(&u.Username)
	if err != nil {
		return nil
	}
	query := `delete from users where index = $1`
	fields := []any{u.ID}
	_, err = database.PGClient.ExecContext(ctx, query, fields...)
	return err
}
