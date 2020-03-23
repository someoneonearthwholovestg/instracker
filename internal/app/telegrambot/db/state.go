package db

import (
	"github.com/jackc/pgx"
	"time"
)

type User struct {
	ID               int64
	Username         string
	FirstName        string
	LastName         string
	State            int
	RegistrationDate time.Time
}

func GetUserState(conn *pgx.ConnPool, userID int64) (state int, err error) {
	sql := `SELECT "state" FROM users WHERE "id" = $1`
	err = conn.QueryRow(sql, userID).Scan(&state)
	return
}

func CreateOrUpdateUser(conn *pgx.ConnPool, user *User) {
	_, err := GetUserState(conn, user.ID)
	if err != nil {
		CreateUser(conn, user)
		return
	}

	UpdateUser(conn, user)
}

func CreateUser(conn *pgx.ConnPool, user *User) {
	sqlInsert := `
		INSERT INTO "users"
			("id", "username", "firstname", "lastname", "state", "registration_date")
		VALUES 
			($1, $2, $3, $4, $5, $6)
	`
	_, _ = conn.Exec(sqlInsert,
		user.ID,
		user.Username,
		user.FirstName,
		user.LastName,
		user.State,
		time.Now(),
	)
}

func UpdateUser(conn *pgx.ConnPool, user *User) {
	sqlUpdate := `
		UPDATE 
			users
		SET 
			"username" = $2,
			"firstname" = $3,
			"lastname" = $4,
			"state" = $5
		WHERE 
			"id" = $1
	`
	_, _ = conn.Exec(sqlUpdate,
		user.ID,
		user.Username,
		user.FirstName,
		user.LastName,
		user.State,
	)
}
