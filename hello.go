package main

import (
	"fmt"
	"os" //essa biblioteca comunica a aplicação com o sistema operacional
)

func main() {

	showIntro()
	showMenu()
	command := readOption()

	switch command {
	case 1:
		fmt.Println("Start monitoring")
	case 2:
		fmt.Println("Retrieving data...")
	case 3:
		fmt.Println("Goodbye")
		os.Exit(0) // sai do programa sem erros
	default:
		fmt.Println("Unknown command... exiting application")
		os.Exit(-1) //sai do programa retornando um erro
	}
}

func showIntro() {
	version := 1.0
	fmt.Printf("This program is running on version %v\n\n", version)
}

func showMenu() {
	fmt.Println("1 - Start monitoring")
	fmt.Println("2 - Show logs")
	fmt.Println("3 - Exit")
}

func readOption() int {
	var option int
	fmt.Scan(&option)

	return option
}
