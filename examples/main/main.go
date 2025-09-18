package main

import (
	"fmt"

	"github.com/ericorlovski/go-builder/examples"
)

func main() {
	user := examples.NewUserBuilder().
		WithID(1).
		WithName("Don").
		WithEmail("don@example.com").
		Build()

	fmt.Printf("User: %+v\n", user)
}
