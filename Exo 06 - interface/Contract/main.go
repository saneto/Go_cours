package main

import (
	"fmt"
	"strings"
)

type Caller interface {
	Call(string) string
}

type User struct {
	Name string
}

// TODO: add a method Call attached to User so this type can implement the Caller interface.
// the implemented function shoud uppercase the given message in it's param function.

func main() {
	u := User{"Michelle"}
	Do(u, "love me!")
}

func Do(c Caller, msg string) {
	fmt.Println(c.Call(msg))
}

func (u User) Call(msg string) string{
	return strings.ToUpper(msg)
}