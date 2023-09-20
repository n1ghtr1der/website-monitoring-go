package main

import (
	"fmt"
)

func main() {
	version := 1.0
	fmt.Printf("This program is running on version %v\n", version)

	fmt.Println("1 - Start monitoring")
	fmt.Println("2 - Show logs")
	fmt.Println("3 - Exit")

	var option int
	fmt.Scan(&option)
	fmt.Println(option)
}
