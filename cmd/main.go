package main

import (
	"fmt"

	model "core-backend/pkg/model"
)

func main() {
	var employee = model.User{
		ID:          1,
		FirstName:   "First name",
		LastName:    "Last Name",
		BadgeNumber: 1000,
	}
	fmt.Printf(employee.FirstName)
}
