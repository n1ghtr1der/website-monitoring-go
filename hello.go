package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http" //biblioteca responsável por fzer
	"os"       //essa biblioteca comunica a aplicação com o sistema operacional
	"strconv"
	"strings"
	"time"
)

const monitoring = 5
const delaySeconds = 3

func main() {

	showIntro()

	for {
		showMenu()
		command := readOption()

		switch command {
		case 1:
			startMonitoring()
		case 2:
			printLogs()
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
	domains := readWebsitesFileList()

	for i := 0; i < monitoring; i++ {
		for i, domain := range domains {
			websiteNum := i + 1
			fmt.Println("Testing website", websiteNum, ":", domain)
			testSite(domain)
		}
		time.Sleep(delaySeconds * time.Second)
	}
}

func testSite(domain string) {
	response, error := http.Get(domain)

	if error != nil {
		fmt.Println("An error has ocurred:", error)
	}

	if response.StatusCode == 200 {
		fmt.Println("Website:", domain, "successfuly loaded!")
		registerLogs(domain, true)
	} else {
		fmt.Println("Website:", domain, "not reachable. Code:", response.StatusCode)
		registerLogs(domain, false)
	}
}

func readWebsitesFileList() []string {
	var websites []string

	file, error := os.Open("websites.txt")

	if error != nil {
		fmt.Println("An error has ocurred:", error)
	}

	reader := bufio.NewReader(file)

	for {
		line, error := reader.ReadString('\n')
		line = strings.TrimSpace(line)
		fmt.Println(line)

		websites = append(websites, line)

		if error == io.EOF {
			break
		}
	}

	file.Close()

	return websites

}

func registerLogs(domain string, status bool) {
	file, error := os.OpenFile("logs.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	if error != nil {
		fmt.Println("An error has ocurred:", error)
	}

	file.WriteString(time.Now().Format("02/01/2006 15:04:05") + " - " + domain + " - online:" + strconv.FormatBool(status) + "\n") //golang time format

	file.Close()
}

func printLogs() {
	fmt.Println("Retrieving Data...")
	time.Sleep(1 * time.Second)
	file, error := os.ReadFile("logs.txt")

	if error != nil {
		fmt.Println(error)
	}

	fmt.Println(string(file))

}
