package main

import (
	"fmt"

	"github.com/ericorlovski/go-builder/examples"
)

func main() {
	user := examples.NewUserBuilder().
		ID(1).
		Name("Don").
		Email("don@example.com").
		Build()

	fmt.Printf("User: %+v\n", user)
}
