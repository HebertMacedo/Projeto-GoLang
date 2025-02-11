package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {

	exibeIntroducao()
	exibeMenu()

	comando := leComando()

	switch comando {
	case 1:
		iniciaMonitoramento()
	case 2:
		fmt.Println("Exibindo Logs...")
	case 0:
		fmt.Println("Encerrando o programa")
		os.Exit(0)
	default:
		fmt.Println("Não conheço este comando")
		os.Exit(-1)
	}

}

func exibeIntroducao() {
	nome := "Hebert"
	version := 1.1
	fmt.Println("Olá, Sr.", nome)
	fmt.Println("Este programa está na versão", version)
}

func exibeMenu() {
	fmt.Println("1- Iniciar Monitoramento")
	fmt.Println("2- Exibir Logs")
	fmt.Println("0- Sair do Programa")
}

func leComando() int {
	var comandoLido int
	fmt.Scan(&comandoLido)
	fmt.Println("O comando escolhido foi", comandoLido)

	return comandoLido
}

func iniciaMonitoramento() {
	fmt.Println("Monitorando...")
	sites := []string{

		"https://random-status-code.herokuapp.com",
		"https://pontonet.srv.br/hsma/index.php?m=1",
		"https://rboconcursos.selecao.net.br/",
	}

	fmt.Println(sites)

	site := "https://random-status-code.herokuapp.com"
	resp, _ := http.Get(site)

	if resp.StatusCode == 200 {
		fmt.Println("Site:", site, "foi carregado com sucesso!")
	} else {
		fmt.Println("Site:", site, "esta com problema. Status Code:",
			resp.StatusCode)
	}
}
