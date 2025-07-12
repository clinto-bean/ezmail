package main

import (
	"net/http"
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID       uuid.UUID
	Username string
	EmailAddress string
	RegisteredAt time.Time
	DOB time.Time
}

type UserParams struct {
Username string `json:"username"`
		EmailAddress string `json:"email"`
		DOB time.Time `json:"dob"`
}

func validateUserAge(dob time.Time) bool {
	now := time.Now()
	ageThreshold := now.AddDate(-18,0,0)
	return !dob.After(ageThreshold)
}

// User functions
// createUser takes in a username and email string as well as an ISO 8601 formatted date for the date of birth
func createUser(username, emailAddress string, dob time.Time) (bool, User) {
	ok := validateUserAge(dob)
	if !ok {
		return false, User{}
	}
	return ok, User{
		ID: uuid.New(), 
		Username: username, 
		EmailAddress: emailAddress, 
		RegisteredAt: time.Now(),
		DOB: dob, 
	}
}

// User API Methods
// handlerCreateUser will attempt to call createUser with the request body's data 
func (a *API) handlerCreateUser(w http.ResponseWriter, r *http.Request) {

	params := UserParams{}

	if r.Method != "POST" {
		errorResponse(w, 405, "method not allowed")
	}

	err := decodeJSON(r.Body, &params)
	if err != nil {
		errorResponse(w, 400, "bad request parameters")
	}

	ok, u := createUser(params.Username, params.EmailAddress, params.DOB)
	if ok {
		jsonResponse(w, 200, u)
	} else {
		errorResponse(w, 400, "user must be at least 18 years old to register")
	}
}