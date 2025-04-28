package main

import (
	"fmt"
)

func main() {
	profileuser("Farrel Tito", "Jl. Ketintang", 25)
}

func profileuser(username, address string, age int) {
	fmt.Printf("Hello World! My name %s and I'm live in %s, and my age around %d years old.", username, address, age)
}
