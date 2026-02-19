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

func (u *User) UpdateLocation(ctx context.Context) error {
	query := `select location_id from users where index = $1`
	fields := []any{u.ID}
	err := database.PGClient.QueryRowContext(ctx, query, fields...).Scan(&u.Location.ID)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return err
	}
	query = `update locations set name=$1, street=$2, street_number=$3, country=$4, federal_state=$5, updated_at=$6 where index=$7`
	fields = []any{u.Location.Name, u.Location.Street, u.Location.StreetNumber, u.Location.Country, u.Location.FederalState, u.Location.UpdatedAt, u.Location.ID}
	_, err = database.PGClient.ExecContext(ctx, query, fields...)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return err
	}
	return nil
}

func (u *User) GetLocation(ctx context.Context) error {
	query := `select location_id from users where index = $1`
	fields := []any{u.ID}
	err := database.PGClient.QueryRowContext(ctx, query, fields...).Scan(&u.Location.ID)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return err
	}
	query = `select * from locations where index=$1`
	fields = []any{u.Location.ID}
	err = database.PGClient.QueryRowContext(ctx, query, fields...).Scan(
		&u.Location.ID, &u.Location.Name, &u.Location.CreatedAt, &u.Location.UpdatedAt, &u.Location.Street,
		&u.Location.StreetNumber, &u.Location.Country, &u.Location.FederalState,
	)
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

func (u *User) Update(ctx context.Context) error {
	err := u.UpdateLocation(ctx)
	if err != nil {
		slog.Error("updating", "location", u.Location, "error", err)
		return err
	}
	nameQuery := `select created_at from users where index=$1`
	nameFields := []any{u.ID}
	err = database.PGClient.QueryRowContext(ctx, nameQuery, nameFields...).Scan(&u.CreatedAt)
	if err != nil {
		slog.Error("selecting created_at", "error", err)
		return err
	}
	query := `update users set password=$1, email=$2, updated_at=$3 where index=$4`
	fields := []any{u.Password, u.EMail, u.UpdatedAt, u.ID}
	_, err = database.PGClient.ExecContext(ctx, query, fields...)
	return err
}

func (u *User) Get(ctx context.Context) error {
	err := u.GetLocation(ctx)
	if err != nil {
		slog.Error("getting location", "error", err)
		return err
	}
	query := `select * from users where index=$1`
	fields := []any{u.ID}
	return database.PGClient.QueryRowContext(ctx, query, fields...).Scan(
		&u.ID, &u.Username, &u.EMail, &u.Password, &u.Location.ID, &u.CreatedAt, &u.UpdatedAt,
	)
}
