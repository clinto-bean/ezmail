package main

import (
	"errors"
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
		ID uuid.UUID
}

// function validates that user is over the age of 18 and will return an error if not
func validateUserAge(dob time.Time) error {
	now := time.Now()
	ageThreshold := now.AddDate(-18,0,0)
	if dob.After(ageThreshold) {
		return errors.New("user is too young")
	}
	return nil
}

// User functions
// createUser takes in a username and email string as well as an ISO 8601 formatted date for the date of birth
func createUser(username, emailAddress string, dob time.Time) (User, error) {
	err := validateUserAge(dob)
	if err != nil {
		return User{}, err
	}
	return  User{
		ID: uuid.New(), 
		Username: username, 
		EmailAddress: emailAddress, 
		RegisteredAt: time.Now(),
		DOB: dob, 
	}, nil
}

func modifyUser(username, emailAddress string, id uuid.UUID) (error, User) {
	return nil, User{}
}

// User API Methods
// handleUser will determine the correct handler based on the request's http method
func (a *API) handleUser(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		// query db for user
	}
	if r.Method == "POST" {
		a.handlerCreateUser(w, r)
	}
	if r.Method == "PUT" {
		a.handlerModifyUser(w, r)
	}
}

// handlerCreateUser will attempt to call createUser with the request body's data 
func (a *API) handlerCreateUser(w http.ResponseWriter, r *http.Request) {
	params := UserParams{}
	err := decodeJSON(r.Body, &params)
	if err != nil {
		errorResponse(w, 400, "bad request parameters")
	}
	u, err := createUser(params.Username, params.EmailAddress, params.DOB)
	if err != nil {
		errorResponse(w, 400, "user must be at least 18 years old to register")
	} else {
		jsonResponse(w, 200, u)
	}
}
// handlerModifyUser will attempt to modify the user's profile and return the user if successful
func (a *API) handlerModifyUser(w http.ResponseWriter, r *http.Request) {
	params := UserParams{}
	err := decodeJSON(r.Body, &params)
	if err != nil {
		errorResponse(w, 400, "bad request parameters")
	}
	err, u := modifyUser(params.Username, params.EmailAddress, params.ID)
	if err != nil {
		errorResponse(w, 500, "failed to modify user")
	} else {
		jsonResponse(w, 200, u)
	}
}