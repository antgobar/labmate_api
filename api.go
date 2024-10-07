package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

type api struct {
	addr string
}

var users = []User{}

func (s *api) getUsersHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	err := json.NewEncoder(w).Encode(users)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (s *api) createUsersHandler(w http.ResponseWriter, r *http.Request) {
	var payload User
	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	u := User{
		Name: payload.Name,
		Age:  payload.Age,
	}

	if err = insertUser(u); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func insertUser(u User) error {
	if u.Name == "" {
		return errors.New("Name required")
	}

	if u.Age < 0 {
		return errors.New("age must be positive")
	}

	for _, user := range users {
		if user.Name == u.Name {
			return fmt.Errorf("user %s already exists", u.Name)
		}
	}

	users = append(users, u)
	return nil
}
