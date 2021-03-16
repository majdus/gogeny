package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type User struct {
	Name string		`json:"name"`
	Address string	`json:"address"`
	Age int			`json:"age"`
}

var users []User

func sayHello(w http.ResponseWriter, _ *http.Request) {
	_, _ = fmt.Fprintf(w, "Hello geny\n")
}

func getUsers(w http.ResponseWriter, _ *http.Request) {
	_ = json.NewEncoder(w).Encode(users)
}

func addUser(w http.ResponseWriter, r *http.Request) {
	bUser, _ := ioutil.ReadAll(r.Body)
	var user User
	err := json.Unmarshal(bUser, &user)
	fmt.Println(string(bUser))
	fmt.Print(user)
	if err == nil {
		users = append(users, user)
		_, _ = fmt.Fprintf(w,"User successfully added\n")
	} else {
		_, _ = fmt.Fprintf(w,"Error while adding user\n")
	}
}

func main() {
	user1 := User{
		Name:    "user1",
		Address: "Paris",
		Age:     30,
	}

	user2 := User{
		Name:    "user2",
		Address: "Lyon",
		Age:     40,
	}

	users = append(users, user1)
	users = append(users, user2)

	http.HandleFunc("/", sayHello)
	http.HandleFunc("/users", getUsers)
	http.HandleFunc("/adduser", addUser)
	log.Fatal(http.ListenAndServe(":9999", nil))
}
