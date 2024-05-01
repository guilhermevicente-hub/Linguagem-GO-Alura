package main

import "fmt"
import "os"

func main() { 

	exibeIntroducao()
	exibeMenu()
	comando := lerComando()

	switch comando {
	case 1:
		fmt.Println("Monitorando...")
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

func exibeIntroducao(){
	nome := "Guilherme"
	versao := 1.1
	fmt.Println("Olá, sr", nome)
	fmt.Println("Este programa está na versão", versao)
}

func lerComando() int {
	var comandoLido int
	fmt.Scan(&comandoLido)
	fmt.Println("O comando escolhido foi", comandoLido)
	
	return comandoLido
}

func exibeMenu(){
	fmt.Println("1 - Iniciar Monitoramento")
	fmt.Println("2 - Exibir Logs")
	fmt.Println("0 - Sair do Programa")
}




	// if comando == 1 {
	// 	fmt.Println("Monitorando...")
	// }else if comando == 2 {
	// 	fmt.Println("Exibindo Logs...")
	// }else if comando == 0 {
	// 	fmt.Println("Saindo do programa")
	// }else{
	// 	fmt.Println("Não conheço este comando")
	// }