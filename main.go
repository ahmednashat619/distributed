package main

import (
	"fmt"
	"log"

	"github.com/ahmednashat619/go-docker-app/models"
)

func main() {
	u, err := models.AddUser(models.User{Name: "Alice", Age: 30, Address: models.Address{City: "Cairo", Country: "Egypt"}})
	if err != nil {
		log.Fatalf("error adding user: %v", err)
	}

	fmt.Printf("✅ User added: %+v\n", u)

	user, err := models.GetUserByID(u.ID)
	if err != nil {
		log.Fatalf("error retrieving user: %v", err)
	}

	fmt.Printf("👤 Retrieved user: %+v\n", user)
}
