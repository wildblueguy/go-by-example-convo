package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	Name  string
	Email string
}

func main() {
	user := User{
		Name:  "First Last",
		Email: "first-last@domain.tld",
	}
	userJson, _ := json.Marshal(user)
	fmt.Println(string(userJson))
}
