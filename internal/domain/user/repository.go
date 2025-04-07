package user

import (
	"context"
	"database/sql"
	"errors"
)

type PostgresRepo struct {
	db *sql.DB
}

func NewPostgresRepo(db *sql.DB) *PostgresRepo {
	return &PostgresRepo{db: db}
}

func (r *PostgresRepo) GetByID(id string) (*User, error) {
	row := r.db.QueryRowContext(context.Background(), `
        SELECT id, name, email FROM users WHERE id = $1
    `, id)

	var user User
	err := row.Scan(&user.Id, &user.Name, &user.Email)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("user not found")
		}
		return nil, err
	}

	return &user, nil
}

func (r *PostgresRepo) Save(user *User) error {
	_, err := r.db.ExecContext(context.Background(), `
        INSERT INTO users (id, name, email)
        VALUES ($1, $2, $3)
    `, user.Id, user.Name, user.Email)
	return err
}
