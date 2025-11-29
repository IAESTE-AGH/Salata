package repository

import (
	"database/sql"
	"go_server/internal/models"
)

func GetAllUsers(db *sql.DB) ([]models.User, error) {
	rows, err := db.Query("SELECT * FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	users := []models.User{}
	for rows.Next() {
		var u models.User
		//rows.Scan(&u.ID, &u.Name, &u.Email, &u.Group)
		rows.Scan(&u.ID, &u.Name)
	}
	return users, nil
}
