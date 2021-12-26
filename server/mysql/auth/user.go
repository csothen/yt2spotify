package auth

import (
	"database/sql"
	"fmt"

	"github.com/csothen/yt2spotify/data"
)

type MySQLRepository struct {
	db *sql.DB
}

func NewMySQLRepository(db *sql.DB) *MySQLRepository {
	return &MySQLRepository{
		db: db,
	}
}

func (r *MySQLRepository) GetAll() []*data.User {
	users := []*data.User{}

	results, err := r.db.Query("SELECT id, access_token, refresh_token, token_type, expires_in FROM users")
	if err != nil {
		return users
	}

	for results.Next() {
		var u data.User

		err := results.Scan(&u.ID, &u.AccessToken, &u.RefreshToken, &u.TokenType, &u.ExpiresIn)
		if err != nil {
			return users
		}

		users = append(users, &u)
	}
	return users
}

func (r *MySQLRepository) GetByID(id string) (*data.User, error) {
	result := r.db.QueryRow("SELECT id, access_token, refresh_token, token_type, expires_in FROM users WHERE id = ?", id)

	var u data.User

	err := result.Scan(&u.ID, &u.AccessToken, &u.RefreshToken, &u.TokenType, &u.ExpiresIn)
	if err != nil {
		return nil, err
	}

	return &u, nil
}

func (r *MySQLRepository) Upsert(u *data.User) (*data.User, error) {
	id := u.ID
	_, err := r.GetByID(id)
	if err != nil {
		if err != sql.ErrNoRows {
			return nil, err
		}
		// User does not exist so we create it
		return r.create(u)
	}
	// User already exists so we update it
	return r.update(u)
}

func (r *MySQLRepository) Delete(id string) error {
	res, err := r.db.Exec("DELETE FROM users WHERE id = ?", id)
	if err != nil {
		return err
	}

	n, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if n == 0 {
		return fmt.Errorf("no users were deleted")
	}

	return nil
}

func (r *MySQLRepository) create(u *data.User) (*data.User, error) {
	res, err := r.db.Exec("INSERT INTO users(id, access_token, refresh_token, token_type, expires_in) VALUES (?, ?, ?, ?, ?)", u.ID, u.AccessToken, u.RefreshToken, u.TokenType, u.ExpiresIn)
	if err != nil {
		return nil, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return nil, err
	}

	return r.GetByID(fmt.Sprintf("%d", id))
}

func (r *MySQLRepository) update(u *data.User) (*data.User, error) {
	stmt, err := r.db.Prepare("UPDATE users SET access_token = ?, refresh_token = ?, token_type = ?, expires_in = ? WHERE id = ?")
	if err != nil {
		return nil, err
	}

	_, err = stmt.Exec(u.AccessToken, u.RefreshToken, u.TokenType, u.ExpiresIn, u.ID)
	if err != nil {
		return nil, err
	}

	return r.GetByID(u.ID)
}
