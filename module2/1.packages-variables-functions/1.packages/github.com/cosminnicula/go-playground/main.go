package main

import (
	"fmt"

	"github.com/cosminnicula/go-playground/abc"
	xyz "github.com/cosminnicula/go-playground/xyz" // lib.go from xyz folder is declare as package xyz_something, so here we're using an alias
)

// The root of a project is always known as the main package => main.go is assigned to the main package
func main() {

	fmt.Println(abc.PublicVar)
	fmt.Println(abc.PublicFunction())
	fmt.Println(xyz.Something())
}
