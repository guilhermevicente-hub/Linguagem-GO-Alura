package main

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

const monitoramentos = 3

const delay = 5

func main() {

	exibeIntroducao()

	for {

		exibeMenu()
		comando := lerComando()

		switch comando {
		case 1:
			iniciarMonitoramento()
		case 2:
			fmt.Println("Exibindo Logs...")
		case 0:
			fmt.Println("Saindo do programa")
			os.Exit(0)
		default:
			fmt.Println("Não conheço este comando")
			os.Exit(-1)
		}
	}

}

func exibeIntroducao() {
	nome := "Guilherme"
	versao := 1.1
	fmt.Println("Olá, sr", nome)
	fmt.Println("Este programa está na versão", versao)
}

func lerComando() int {
	var comandoLido int
	fmt.Scan(&comandoLido)
	fmt.Println("O comando escolhido foi", comandoLido)
	fmt.Println("")

	return comandoLido
}

func exibeMenu() {
	fmt.Println("1 - Iniciar Monitoramento")
	fmt.Println("2 - Exibir Logs")
	fmt.Println("0 - Sair do Programa")
}

func iniciarMonitoramento() {
	fmt.Println("Monitorando...")
	sites := [] string {"https://httpbin.org/status/200", "https://alura.com.br", "https://httpbin.org/status/404"}

	for i := 0; i < monitoramentos ; i++{
		for i, site := range sites{
			fmt.Println("Testando site", i, ":", site)
			testaSites(site)
		}
		time.Sleep(delay * time.Second)
		fmt.Println("")
	}

	fmt.Println("")

}

func testaSites(site string) {
	
	resp, _ := http.Get(site)

	if resp.StatusCode == 200 {
		fmt.Println("Site:", site, "foi carregado com sucesso!")
	} else {
		fmt.Println("Site:", site, "esta com problemas. Status Code:", resp.StatusCode)
	}
}
