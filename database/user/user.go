package user

import (
	"github.com/monstrasitix/gopnik/database"
)

type User struct {
    Id int
    FirstName string
    LastName string
    Email string
    DateOfBirth string
}

func GetUsers() []User {
	var users []User

    db, err := database.Open()
    database.HandleError(err)

    rows, err := db.Query("call select_users()")
    database.HandleError(err)

    for rows.Next() {
	    var id int
	    var firstName string
	    var lastName string
	    var email string
	    var dateOfBirth string

	    rows.Scan(
		    &id,
		    &firstName,
		    &lastName,
		    &email,
		    &dateOfBirth)

	    users = append(users, User {
		    Id: id,
		    FirstName: firstName,
		    LastName: lastName,
		    Email: email,
		    DateOfBirth: dateOfBirth,
	    })
    }

    db.Close()
    rows.Close()

    return users
}
