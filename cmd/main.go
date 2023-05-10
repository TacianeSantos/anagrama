package main

import (
	"fmt"

	"anagrama/infra/arquivo"
)

func main() {
	fmt.Println("Vamos conhecer os anagramas das palavras?!")

	menuOrigemPalavra()

	opcao := opcaoDeEscolha()

	var palavras []string

	switch opcao {
	case 1:
		var palavra string
		fmt.Println("Você escolheu: Digitar uma nova palavra")
		fmt.Print("A palavra é:")
		fmt.Scan(&palavra)

		palavras = append(palavras, palavra)

	case 2:
		fmt.Println("Você escolheu: Palavra aleatória")
		qtdePalavra := 1

		palavras = arquivo.LerPalavraDoArquivo(qtdePalavra)

	default:
		fmt.Println("Desculpa, mas desconheço essa opção")
	}

	fmt.Println(palavras)
	// anagrama.CalculaAnagrama(palavra)
}

func menuOrigemPalavra() {
	fmt.Println("1 - Digitar palavra")
	fmt.Println("2 - Escolher palavra de forma aleatória")
}

func opcaoDeEscolha() int {
	fmt.Println("\nSua vez, escolha uma opção:")
	var opcao int

	fmt.Scan(&opcao)

	return opcao
}
