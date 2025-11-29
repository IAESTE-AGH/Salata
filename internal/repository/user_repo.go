package repository

import (
	"database/sql"
	"go_server/internal/models"
)

func GetAllUsers(db *sql.DB) ([]models.User, error) {
	rows, err := db.Query(`SELECT id, first_name, last_name, email, "group" FROM users`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	users := []models.User{}

	for rows.Next() {
		var u models.User
		var first, last, group string

		if err := rows.Scan(
			&u.ID,
			&first,
			&last,
			&u.Email,
			&group,
		); err != nil {
			return nil, err
		}

		u.Name.First = first
		u.Name.Last = last
		u.Group = models.Group(group)

		users = append(users, u)
	}
	return users, nil
}
