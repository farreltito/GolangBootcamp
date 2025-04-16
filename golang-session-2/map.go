package main

import (
	"fmt"
)

/*
	This is a Map project.
*/

func main() {
	//Map
	//Method 1 (Without assign data)
	var user1 map[string]string //Declare
	user1 = map[string]string{} //Initialization

	user1["username"] = "farreltito"
	user1["password"] = "root"
	user1["email"] = "tito@example.com"

	fmt.Println("Username:", user1["username"])
	fmt.Println("Password:", user1["password"])
	fmt.Println("E-Mail:", user1["email"])
	fmt.Println("")

	//Method 2 (instant assign data)
	var user2 = map[string]string{
		"username": "yunasinaga",
		"password": "root",
		"email":    "yuna@example.com",
	}

	fmt.Println("Username:", user2["username"])
	fmt.Println("Password:", user2["password"])
	fmt.Println("E-Mail:", user2["email"])
	fmt.Println("")

	//Map (Looping with map)
	for key2, value2 := range user1 {
		fmt.Println(key2, ":", value2)
	}
	fmt.Println("")

	//Map (Deleting value)
	fmt.Println("Checking data:", user1)
	delete(user1, "email")
	fmt.Println("After delete data:", user1)

	//Map (Detecting a value)
	value3, exist3 := user2["email"]
	if exist3 {
		fmt.Println(value3)
	} else {
		fmt.Println("Value doesn't exist")
	}

	delete(user2, "email")

	value3, exist3 = user2["email"]
	if exist3 {
		fmt.Println(value3)
	} else {
		fmt.Println("Value has been deleted")
	}

	//Map (Combining slice with map)
	var user3 = []map[string]string{
		{"username": "ibebebe", "password": "root"},
		{"username": "rudididi", "password": "root"},
		{"username": "gunawawan", "password": "root"},
	}

	for i, fuser3 := range user3 {
		fmt.Printf("Index number: %d, username: %s, password: %s\n", i, fuser3["username"], fuser3["password"])
	}
}

/*
	Note:

	Result:
	Username: farreltito
	Password: root
	E-Mail: tito@example.com

	Username: yunasinaga
	Password: root
	E-Mail: yuna@example.com

	username : farreltito
	password : root
	email : tito@example.com

	Checking data: map[email:tito@example.com password:root username:farreltito]
	After delete data: map[password:root username:farreltito]
	yuna@example.com
	Value has been deleted
	Index number: 0, username: ibebebe, password: root
	Index number: 1, username: rudididi, password: root
	Index number: 2, username: gunawawan, password: root
*/
