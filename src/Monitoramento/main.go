package main

import (
	"fmt"
	"net/http"
	"os"
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
			fmt.Println("Exibindo Logs...")
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

	sites := []string{"https://random-status-code.herokuapp.com/",
		"https://www.alura.com.br", "https://www.caelum.com.br"}

	// var resp *http.Response
	for _, site := range sites {
		resp, _ := http.Get(site)

		if resp.StatusCode == 200 {
			fmt.Println("Site:", site, "-> Carregado com sucesso! StatusCode:", resp.StatusCode)
		} else {
			fmt.Println("Site:", site, "-> Fora do ar! StatusCode:", resp.StatusCode)
		}
	}

}
