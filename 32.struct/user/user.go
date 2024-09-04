package user

import (
	"errors"
	"math/rand"
	"time"
)

type User struct {
	firstName string
	lastName  string
	UserId    int
	emailId   string
	createdAt time.Time
}

type Admin struct {
	User    // embading User struct in Admin struct
	isAdmin bool
}

// constructor function / method
func New(firstName, lastName, emailId string) (*User, error) {
	// ✅ validation
	if firstName == "" || lastName == "" || emailId == "" {
		return nil, errors.New("🚫 Invalid parameter")
	}
	return &User{
		firstName: firstName,
		lastName:  lastName,
		emailId:   emailId,
		UserId:    rand.Intn(5),
	}, nil
}
