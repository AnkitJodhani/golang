package models

import (
	"time"

	"ankit.com/temp/db"
)

type Event struct {
	ID          int64
	UserID      int
	Name        string    `binding:"required"`
	Description string    `binding:"required"`
	Location    string    `binding:"required"`
	DateTime    time.Time `binding:"required"`
}

func (e *Event) Save() (*Event, error) {
	query := `INSERT INTO events(user_id,name,description,location,datetime)
	VALUES (?,?,?,?,?)`
	stmt, err := db.DB.Prepare(query)
	// fmt.Println("11") // 🧪 Just for Testing
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	result, err := stmt.Exec(e.UserID, e.Name, e.Description, e.Location, e.DateTime)
	// fmt.Println("12") // 🧪 Just for Testing

	if err != nil {
		return nil, err
	}
	id, err := result.LastInsertId()
	// fmt.Println("13") // 🧪 Just for Testing

	if err != nil {
		return nil, err
	}
	e.ID = id
	return e, nil

}

func GetAllEvents() ([]Event, error) {
	query := `SELECT * FROM events`
	rows, err := db.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var events []Event
	for rows.Next() {
		var event Event
		err := rows.Scan(&event.ID, &event.UserID, &event.Name, &event.Description, &event.Location, &event.DateTime)
		if err != nil {
			return nil, err
		}
		events = append(events, event)
	}
	return events, nil
}
