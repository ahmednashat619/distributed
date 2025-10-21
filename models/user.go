package models

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
)

type User struct {
	ID      int
	Name    string // must be Name not name :D, won’t work
	Age     int
	Address Address
}

var (
	// Initialize nextID based on number of files in users_saved directory
	nextID = func() int {
		const dir = "../users_saved"
		files, err := ioutil.ReadDir(dir)
		if err != nil {
			// If directory doesn’t exist, create it
			_ = os.MkdirAll(dir, os.ModePerm)
			return 1
		}
		return len(files) + 1
	}()
)

func AddUser(u User) (User, error) {
	if u.ID != 0 {
		return User{}, errors.New("new user must not include id or it must be set to zero")
	}
	u.ID = nextID
	nextID++

	// Ensure directory exists
	dir := "../users_saved"
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		if err := os.MkdirAll(dir, os.ModePerm); err != nil {
			return User{}, fmt.Errorf("failed to create users_saved dir: %v", err)
		}
	}

	// Create file and write JSON data
	filePath := fmt.Sprintf("%s/%d.txt", dir, u.ID)
	data, err := json.MarshalIndent(u, "", "  ")
	if err != nil {
		return User{}, fmt.Errorf("failed to marshal user: %v", err)
	}

	if err := ioutil.WriteFile(filePath, data, 0644); err != nil {
		return User{}, fmt.Errorf("failed to write user file: %v", err)
	}

	return u, nil
}

func GetUserByID(id int) (User, error) {
	dir := "../users_saved"
	filePath := fmt.Sprintf("%s/%d.txt", dir, id)

	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		return User{}, fmt.Errorf("user with ID '%v' not found", id)
	}

	var u User
	if err := json.Unmarshal(data, &u); err != nil {
		return User{}, fmt.Errorf("failed to unmarshal user data: %v", err)
	}

	return u, nil
}

type Address struct {
	City    string
	Country string
}
