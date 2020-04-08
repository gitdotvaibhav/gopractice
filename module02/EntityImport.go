package main

import (
	"fmt"

	"./entities"
)

func main() {

	en := entities.Admin{
		Role: "Admin",
		Name: "Vaibhav",
	}

	fmt.Println(en)

}
