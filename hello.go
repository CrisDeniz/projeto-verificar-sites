package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"
)

const monitoringNumber = 3
const monitoringDelay = 5

func main() {

	intro()

	for {

		menu()

		command := readCommand()

		switch command {

		case 1:
			monitor()
		case 2:
			fmt.Println("Exibindo Logs...")
			displayLog()
		case 0:
			fmt.Println("Saindo do programa...")
			os.Exit(0)
		default:
			fmt.Println("Não conheço este comando")
			os.Exit(-1)
		}
	}
}

func monitor() {

	fmt.Println("Monitorando...")

	sites := readArchive()

	for i := 0; i < monitoringNumber; i++ {

		for i, site := range sites {

			fmt.Println("Testando site", i, ":", site)
			siteTest(site)

		}
		time.Sleep(monitoringDelay * time.Second)
	}

	fmt.Println("")

}

func intro() {

	name := "Cristian"
	version := "1.2"

	fmt.Println("Olá sr.", name)
	fmt.Println("Este programa esta na versão", version)
}

func menu() {

	fmt.Println("1- Iniciar Monitoramento")
	fmt.Println("2- Exibir Logs")
	fmt.Println("0- Sair do Programa")
}

func readCommand() int {

	var readCommand int
	fmt.Scan(&readCommand)

	return readCommand
}

func siteTest(site string) {
	resp, err := http.Get(site)

	if err != nil {
		fmt.Println("Ocorreu o seguinte erro:", err)
	}

	if resp.StatusCode == 200 {
		fmt.Println("O site", site, "foi carregado com sucesso")
		logRegister(site, true)
	} else {
		fmt.Println("O site", site, "esta com problemas")
		logRegister(site, false)
	}
}

func readArchive() []string {

	var sites []string

	archive, err := os.Open("sites.txt")

	if err != nil {
		fmt.Println("Ocorreu o seguinte erro:", err)
	}

	reader := bufio.NewReader(archive)

	for {
		line, err := reader.ReadString('\n')
		line = strings.TrimSpace(line)
		fmt.Println(line)

		sites = append(sites, line)

		if err == io.EOF {
			break
		}

	}

	archive.Close()

	return sites
}

func logRegister(site string, status bool) {

	archive, err := os.OpenFile("log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {
		fmt.Println("Ocorreu o seguinte erro:", err)
	}

	if status {
		archive.WriteString(time.Now().Format("02/01/2006 15:04:05") + " - " + "O site: " + site + " esta online" + "\n")
	} else {
		archive.WriteString(time.Now().Format("02/01/2006 15:04:05") + " - " + "O site: " + site + " esta offline" + "\n")
	}

	archive.Close()

}

func displayLog() {
	archive, err := ioutil.ReadFile("log.txt")

	if err != nil {
		println("Ocooreu o seguinte erro:", err)
	}

	println(string(archive))
}
