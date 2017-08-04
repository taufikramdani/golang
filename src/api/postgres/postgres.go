package postgres

import (
	"api"
	"database/sql"

	_ "github.com/lib/pq" //database library
)

// UserService represents a PostgreSQL implementation
type UserService struct {
	DB *sql.DB
}

// user function returns a user for a given id
func (s *UserService) User(id int) (*api.User, error) {
	var u api.User
	row := db.QueryRow(`Select id, name FROM users WHERE id= $1`, id)
	if row.Scan(&u.ID, $u.Name); err != nil {
		return nil, err
	}
	return &u, nil
}
