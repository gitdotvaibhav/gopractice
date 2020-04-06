package main

import (
	"fmt"
)

type notifier interface {
	notify()
}

type user struct {
	name  string
	email string
}

func (u *user) notify() {
	fmt.Printf("Notify user with username %s and  mail %s\n", u.name, u.email)

}

func main() {

	u := user{name: "vaibhav", email: "vabhi42"}

	u.notify()
}
