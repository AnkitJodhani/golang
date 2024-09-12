package models

import (
	"fmt"
	"strconv"

	"ankit.com/event_booking/db"
)

type Registration struct {
	ID       int64
	Event_id int64
	User_id  int64
}

func CreateRegistration(event_id, user_id int64) (*int64, error) {
	query := `INSERT INTO registrations(event_id,user_id) VALUES(?,?)`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	result, err := stmt.Exec(event_id, user_id)
	if err != nil {
		return nil, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}
	return &id, nil
}

func GetRegistrationByID(registrationID int64) (*int64, error) {
	query := `SELECT * FROM registrations WHERE id = ?`
	row := db.DB.QueryRow(query, registrationID)
	var id, event_id, user_id int64
	err := row.Scan(&id, &event_id, &user_id)
	if err != nil {
		return nil, err
	}
	fmt.Println(id, event_id, user_id)
	return &user_id, nil
}

func DeleteRegistration(registrationID int64) error {
	query := `DELETE FROM registrations WHERE id = ?`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}

	_, err = stmt.Exec(registrationID)
	if err != nil {
		return err
	}
	return nil
}

func ListRegistrations(userID int64) ([]map[string]string, error) {

	query := `SELECT * FROM registrations WHERE user_id = ?`
	rows, err := db.DB.Query(query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var listOfRegistrations []map[string]string
	for rows.Next() {
		var registration Registration
		mapRegistrations := make(map[string]string)
		err := rows.Scan(&registration.ID, &registration.Event_id, &registration.User_id)
		if err != nil {
			fmt.Println("here")
			return nil, err
		}
		mapRegistrations["ID"] = strconv.FormatInt(registration.ID, 10)
		mapRegistrations["EventID"] = strconv.FormatInt(registration.Event_id, 10)
		mapRegistrations["UserID"] = strconv.FormatInt(registration.User_id, 10)
		fmt.Println(mapRegistrations)
		listOfRegistrations = append(listOfRegistrations, mapRegistrations)
	}
	fmt.Println(listOfRegistrations)
	return listOfRegistrations, nil
}
