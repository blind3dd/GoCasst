package users

import "github.com/gocql/gocql"

type User struct {
	ID        gocql.UUID `json."id"`
	Name      string     `json:"name"`
	LastName  string     `json:"lastname"`
	Email     string     `json:"email"`
	BirthYear int        `json:"birthyear"`
	City      string     `json:"city"`
}

// error array
type ErrorResponse struct {
	Errors []string `json:"errors"`
}

// Response for payload creation returning user struct
type GetUser struct {
	User User `json:"user"`
}

// all users to form array of users struct
type GetAllUsers struct {
	Users []User `json:"users"`
}

// new user ID resource to build payload
type GetNewUser struct {
	ID gocql.UUID `json:"id"`
}

// lowercase !

type ResponseStatus struct {
	Status string `json:"status"`
	Code   int    `json:'code'`
}
