package main

import (
	"fmt"
	"net/http" //biblioteca responsável por fzer
	"os"       //essa biblioteca comunica a aplicação com o sistema operacional
)

func main() {

	showIntro()

	for {
		showMenu()
		command := readOption()

		switch command {
		case 1:
			startMonitoring()
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

func startMonitoring() {
	fmt.Println("Start monitoring")
	site := "http://qacademico.ifce.edu.br/"
	response, _ := http.Get(site) //o _ significa que o segundo valor retornado pela função será ignorado

	if response.StatusCode == 200 {
		fmt.Println("Site:", site, "successfuly loaded!")
	} else {
		fmt.Println("Site:", site, "not reachable. Code:", response.StatusCode)
	}
}
