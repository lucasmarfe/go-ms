package user

import (
	"database/sql"
	"errors"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestPostgresRepo_GetByID_Success(t *testing.T) {
	db, mock, _ := sqlmock.New()
	defer db.Close()

	expected := &User{
		Id:    "1",
		Name:  "user1",
		Email: "user1@test.com",
	}

	rows := sqlmock.NewRows([]string{"id", "name", "email"}).
		AddRow(expected.Id, expected.Name, expected.Email)

	mock.ExpectQuery(`SELECT id, name, email FROM users WHERE id `).
		WithArgs("1").
		WillReturnRows(rows)

	repo := NewPostgresRepo(db)
	user, err := repo.GetByID("1")

	assert.NoError(t, err)
	assert.Equal(t, expected, user)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestPostgresRepo_GetByID_NotFound(t *testing.T) {
	db, mock, _ := sqlmock.New()
	defer db.Close()

	mock.ExpectQuery(`SELECT id, name, email FROM users WHERE id`).
		WithArgs("2").
		WillReturnError(sql.ErrNoRows)

	repo := NewPostgresRepo(db)
	user, err := repo.GetByID("2")

	assert.Nil(t, user)
	assert.EqualError(t, err, "user not found")
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestPostgresRepo_GetByID_DBError(t *testing.T) {
	db, mock, _ := sqlmock.New()
	defer db.Close()

	mock.ExpectQuery(`SELECT id, name, email FROM users WHERE id`).
		WithArgs("3").
		WillReturnError(errors.New("db error"))

	repo := NewPostgresRepo(db)
	user, err := repo.GetByID("3")

	assert.Nil(t, user)
	assert.EqualError(t, err, "db error")
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestPostgresRepo_Save_Success(t *testing.T) {
	db, mock, _ := sqlmock.New()
	defer db.Close()

	user := &User{
		Id:    "1",
		Name:  "user3",
		Email: "user3@test.com",
	}

	mock.ExpectExec(`INSERT INTO users \(id, name, email\)`).
		WithArgs(user.Id, user.Name, user.Email).
		WillReturnResult(sqlmock.NewResult(1, 1))

	repo := NewPostgresRepo(db)
	err := repo.Save(user)

	assert.NoError(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestPostgresRepo_Save_Error(t *testing.T) {
	db, mock, _ := sqlmock.New()
	defer db.Close()

	user := &User{
		Id:    "2",
		Name:  "user2",
		Email: "user2@test.com",
	}

	mock.ExpectExec(`INSERT INTO users \(id, name, email\)`).
		WithArgs(user.Id, user.Name, user.Email).
		WillReturnError(errors.New("insert error"))

	repo := NewPostgresRepo(db)
	err := repo.Save(user)

	assert.EqualError(t, err, "insert error")
	assert.NoError(t, mock.ExpectationsWereMet())
}
