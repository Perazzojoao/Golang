package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

const (
	monitoramentos = 3
	delay          = 5
)

func main() {
	exibirIntro()

	for {
		exibirOpcoes()

		opcao := lerComando()

		switch opcao {
		case 1:
			iniciarMonitoramento()
		case 2:
			exibeLogs()
		case 0:
			fmt.Println("Saindo do Programa...")
			os.Exit(0)
		default:
			fmt.Println("Comando inválido!")
		}
		println()
	}

}

func exibirIntro() {
	var nome = "João Victor"
	var versao float32 = 1.2

	fmt.Println("Olá, sr.", nome)
	fmt.Println("Este programa está na versão", versao)
	println()
}

func exibirOpcoes() {
	fmt.Println("1- Iniciar Monitoramento.")
	fmt.Println("2- Exibir Logs.")
	fmt.Println("0- Sair do Programa.")
	println()
}

func lerComando() int8 {
	var opcao int8
	fmt.Scan(&opcao)
	println()
	println()
	return opcao
}

func iniciarMonitoramento() {
	fmt.Println("Monitorando...")
	println()

	sites := lerSites()

	for i := 0; i < monitoramentos; i++ {
		for i, site := range sites {
			testaSite(i, site)
		}

		if i < monitoramentos-1 {
			time.Sleep(delay * time.Second)
		}
		println()
	}
}

func lerSites() []string {
	file, err := os.Open("sites.txt")
	// file, err := os.ReadFile("sites.txt")

	if err != nil {
		fmt.Println("Ocorreu um erro ->", err)
	}

	leitor := bufio.NewReader(file)

	var lista []string
	for {
		linha, err := leitor.ReadString('\n')

		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Println("Ocorreu um erro ->", err)
		}
		linha = strings.TrimSpace(linha)
		lista = append(lista, linha)
	}
	file.Close()
	return lista
}

func testaSite(i int, site string) {
	resp, err := http.Get(site)

	if err != nil {
		fmt.Println("Ocorreu um erro ->", err)
	}

	fmt.Print("Site ", i+1, ": ")
	if resp.StatusCode == 200 {
		registraLogs(i, site, true)
		fmt.Println(site, "-> Carregado com sucesso! StatusCode:", resp.StatusCode)
	} else {
		registraLogs(i, site, false)
		fmt.Println(site, "-> Fora do ar! StatusCode:", resp.StatusCode)
	}
}

func registraLogs(i int, site string, status bool) {
	file, err := os.OpenFile("logs.txt", os.O_APPEND|os.O_CREATE, 0666)

	if err != nil {
		fmt.Println("Erro ao escrever em arquivo -> Erro:", err)
	}

	file.WriteString(time.Now().Format("02/01/2006 15:04:05 - "))
	file.WriteString("Site " + strconv.FormatInt(int64(i+1), 10) + ": ")

	if status == true {
		file.WriteString(site + " -> Site online\n")
	} else {
		file.WriteString(site + " -> Site offline\n")
	}
	if i == monitoramentos+1 {
		file.WriteString("\n")
	}
	file.Close()
}

func exibeLogs() {
	fmt.Println("Exibindo Logs...")
	println()

	file, err := os.Open("logs.txt")

	if err != nil {
		fmt.Println("Incapaz de abrir o arquivo: logs.txt -> Erro:", err)
	}

	leitor := bufio.NewReader(file)

	for {
		linha, err := leitor.ReadString('\n')

		if err == io.EOF {
			break
		}
		linha = strings.TrimSpace(linha)
		fmt.Println(linha)
	}
	file.Close()
}
