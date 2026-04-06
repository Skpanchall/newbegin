package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
)

type User struct {
	ID    int
	Name  string
	Email string
}

func saveUser(users map[int]User) {
	marshalJson, err := json.MarshalIndent(users, "", "  ")
	if err != nil {
		fmt.Println("Error converting to JSON:", err)
		return
	}
	err = os.WriteFile("users.json", marshalJson, 0644)
	if err != nil {
		fmt.Println("Error writing to users.json:", err)
		return
	}

}
func UserCli() {
	users := make(map[int]User)
	id := flag.Int("id", 0, "user id")
	name := flag.String("name", "", "user name")
	email := flag.String("email", "", "user email")
	add := flag.Bool("add", false, "add a new user")
	update := flag.Bool("update", false, "update an existing user")
	deleteUser := flag.Bool("delete", false, "delete user")
	showUsers := flag.Bool("show", false, "show all users")
	flag.Parse()

	jsonData, err := os.ReadFile("users.json")
	if err != nil {
		fmt.Println("users.json file not found, creating a new one.")
		err = os.WriteFile("users.json", []byte("{}"), 0644)
		if err != nil {
			fmt.Println("Error creating users.json:", err)
			return
		}
		jsonData = []byte("{}")
	}
	if len(jsonData) == 0 {
		users = make(map[int]User)
	} else {
		err = json.Unmarshal(jsonData, &users)
		if err != nil {
			fmt.Println("Error parsing JSON:", err)
			return
		}
	}

	if *add {
		var user User
		if _, exist := users[*id]; exist {
			fmt.Println("User with this ID already exists!")
			return
		}
		user.ID = *id
		user.Name = *name
		user.Email = *email
		users[user.ID] = user
		saveUser(users)
	} else if *update {
		if user, exists := users[*id]; exists {
			if *name != "" {
				user.Name = *name
			}
			if *email != "" {
				user.Email = *email
			}
			users[*id] = user
		} else {
			fmt.Println("User not found!")
		}
		saveUser(users)
		fmt.Println("User updated successfully!")
	} else if *deleteUser {
		if _, exists := users[*id]; exists {
			delete(users, *id)
		} else {
			fmt.Println("User not found!")
		}
		saveUser(users)
	} else if *showUsers {
		if len(users) == 0 {
			fmt.Println("No users found!")
		} else {
			fmt.Println("Users:")
			for _, user := range users {
				fmt.Printf("ID: %d, Name: %s, Email: %s\n", user.ID, user.Name, user.Email)
			}
		}
	} else {
		fmt.Println("Please provide a valid flag. Use -h for help.")
	}
}
