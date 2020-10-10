package main

import (
	"fmt"

	"github.com/NileshGule/golang-bootcamp/webservice/models"
)

func main() {
	u := models.User{
		ID:        2,
		FirstName: "Virat",
		LastName:  "Kohli",
	}

	fmt.Println(u)
}
