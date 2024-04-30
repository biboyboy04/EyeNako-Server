package user

import (
	"database/sql"
	"fmt"

	"github.com/biboyboy04/EyeNako-Server/types"
)

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{db:db}
}

func (s *Store) GetUserByEmail(email string) (*types.User, error){
	rows, err := s.db.Query("SELECT * from user WHERE email = ?", email)
	if err != nil {
		return nil, err
	}

	u := new(types.User)
	for rows.Next() {
		u, err = scanRowsIntoUser(rows)

		if err != nil {
			return nil, err
		}
	}	

	if u.ID == 0 {
		return nil, fmt.Errorf("user not found")
	}

	return u, nil
}

func (s *Store) GetUserByID(id int) (*types.User, error){
	rows, err := s.db.Query("SELECT * FROM user WHERE id = ?", id)
	if err != nil {
		return nil, err
	}

	u := new(types.User)

	for rows.Next() {
		u, err = scanRowsIntoUser(rows)

		if err != nil {
			return nil, err
		}
	}

	if u.ID == 0 {
		return nil, fmt.Errorf("user not found")
	}

	return u, nil
}

func (s *Store) CreateUser(user types.User) error{
_, err := s.db.Exec("INSERT INTO user (username, password, email) VALUES (?, ?, ?)", user.Username, user.Password, user.Email)
	if err != nil {
		return err
	}
	return nil
}

func scanRowsIntoUser(rows *sql.Rows) (*types.User, error) {
	user := new(types.User)

	err := rows.Scan(
		&user.ID,
		&user.Username,
		&user.Email,
		&user.Password,
		&user.Xp,
		&user.LevelID,
		&user.Health,
		&user.CreatedAt,
		&user.UpdatedAt,
		&user.IsArchived,
	)

	if err != nil {
		return nil, err
	}

	return user, nil
}