package main

import (
	"fmt"

	"github.com/matheusbuldrini/gowebservice/models"
)

func main() {
	u := models.User{
		Id:        2,
		FirstName: "Joao",
		LastName:  "Silva",
	}
	fmt.Println(u)
}
