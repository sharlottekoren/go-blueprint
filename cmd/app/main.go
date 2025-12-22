package main

import (
	"fmt"
	"context"

	"github.com/sharlottekoren/go-blueprint/internal/datastore/inmem"
	"github.com/sharlottekoren/go-blueprint/internal/service"
)

// This is just to test that the main function is working.
// In order to run this I had to run `go mod init` and create a go.mod file.
// A go.mod file tells Go what your project is called and which outside packages it depends on.
func main() {
	newMemStore := inmem.NewInMemUserStore()
	fmt.Println("In-memory user store created:", newMemStore)

	newService := service.NewService(newMemStore)
	fmt.Println("Service created with in-memory user store:", newService)

	ctx := context.Background()

	newUser, err := newService.CreateUser(ctx, service.CreateUserRequest{
		Name:  "Alice Smith",
		Email: "abc@def.com",
	})
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("New user created:", newUser)

	getUserID, err := newService.GetUserByID(ctx, newUser.GetID())
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("User retrieved by ID:", getUserID)

	_, err = newService.GetUserByID(ctx, "non-existent-id")
	if err != nil {
		fmt.Println("Expected error retrieving non-existent user:", err)
	}
}