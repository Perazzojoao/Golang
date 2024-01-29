package main

import (
	"fmt"
	"os"

)

func main() {
	exibirIntro()

	exibirOpcoes()

	opcao := lerComando()

	switch opcao {
	case 1:
		fmt.Println("Monitorando...")
	case 2:
		fmt.Println("Exibindo Logs...")
	case 0:
		fmt.Println("Saindo do Programa.")
	default:
		fmt.Println("Comando inválido!")
		os.Exit(0)
	}
	println()
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
	return opcao
}
