package postgres

import (
	"database/sql"
	"errors"

	"github.com/assaidy/expenses-tracking-api/internals/models"
)

// user ===========================================================================
func (pg *PgStorage) CheckUsernameAndEmailConflict(username string, email string) (bool, error) {
	query := `SELECT 1 FROM users WHERE username = $1 OR email = $2 LIMIT 1;`
	if err := pg.db.QueryRow(query, username, email).Scan(new(int)); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return false, nil
		}
		return false, err
	}
	return true, nil
}

func (pg *PgStorage) CheckUsernameConflict(username string) (bool, error) {
	query := `SELECT 1 FROM users WHERE username = $1 LIMIT 1;`
	if err := pg.db.QueryRow(query, username).Scan(new(int)); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return false, nil
		}
		return false, err
	}
	return true, nil
}

func (pg *PgStorage) CheckEmailConflict(email string) (bool, error) {
	query := `SELECT 1 FROM users WHERE email = $1 LIMIT 1;`
	if err := pg.db.QueryRow(query, email).Scan(new(int)); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return false, nil
		}
		return false, err
	}
	return true, nil
}

func (pg *PgStorage) CheckIfUserExists(id int) (bool, error) {
	query := `SELECT 1 FROM users WHERE id = $1 LIMIT 1;`
	if err := pg.db.QueryRow(query, id).Scan(new(int)); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return false, nil
		}
		return false, err
	}
	return true, nil
}

func (pg *PgStorage) CreateUser(user *models.User) error {
	query := `
    INSERT INTO users(name, username, password, email, joined_at)
    VALUES($1, $2, $3, $4, $5);
    `
	if _, err := pg.db.Exec(query,
		user.Name, user.Username, user.Password, user.Email, user.JoinedAt); err != nil {
		return err
	}
	return nil
}

func (pg *PgStorage) GetUserById(id int) (*models.User, error) {
	query := `
    SELECT
        name,
        email,
        username,
        password,
        joined_at
    FROM users
    WHERE id = $1;
    `
	user := models.User{Id: id}
	if err := pg.db.QueryRow(query, id).Scan(
		&user.Name, &user.Email, &user.Username, &user.Password, &user.JoinedAt); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

func (pg *PgStorage) GetUserByUsername(username string) (*models.User, error) {
	query := `
    SELECT
        id,
        name,
        email,
        password,
        joined_at
    FROM users
    WHERE username = $1;
    `
	user := models.User{Username: username}
	if err := pg.db.QueryRow(query, username).Scan(
		&user.Id, &user.Name, &user.Email, &user.Password, &user.JoinedAt); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

func (pg *PgStorage) UpdateUser(newUser *models.User) error {
	query := `
    UPDATE users
    SET 
        name = $1,
        username = $2,
        email = $3,
        password = $4
    WHERE id = $5;
    `
	if _, err := pg.db.Exec(query,
		newUser.Name, newUser.Username, newUser.Email, newUser.Password, newUser.Id); err != nil {
		return err
	}
	return nil
}

func (pg *PgStorage) DeleteUserById(id int) error {
	query := `DELETE FROM users WHERE id = $1;`
	if _, err := pg.db.Exec(query, id); err != nil {
		return err
	}
	return nil
}

// xxx ===========================================================================
