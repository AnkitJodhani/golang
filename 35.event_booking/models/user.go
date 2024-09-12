package models

import (
	"ankit.com/event_booking/db"
	"ankit.com/event_booking/utils"
)

type User struct {
	ID       int64
	Email_ID string `binding:"required"`
	Password string `binding:"required"`
}

func (u *User) Save() (*int64, error) {
	query := `INSERT INTO users(email_id,password)
	VALUES (?,?)`

	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	hashedPassword, err := utils.HashPassword(u.Password)
	if err != nil {
		return nil, err
	}
	result, err := stmt.Exec(u.Email_ID, hashedPassword)

	if err != nil {
		return nil, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}
	u.ID = id
	return &u.ID, nil

}

func (u *User) Verify() bool {
	query := `SELECT id,password FROM users WHERE email_id=?`
	row := db.DB.QueryRow(query, u.Email_ID)
	var retrivePassword string
	err := row.Scan(&u.ID, &retrivePassword)
	if err != nil {
		return false
	}

	// compare the hashedpasowd with orignail passowrd
	PassOk := utils.ComparePassword(retrivePassword, u.Password)
	if !PassOk {
		// password is incorrect
		return PassOk // false
	}
	// password is correct
	return PassOk //true
}
