package main

import (
	"fmt"

	"anagrama/domain/anagrama"
)

func main() {
	fmt.Println("Vamos conhecer os anagramas das palavras?!")

	menuOrigemPalavra()

	opcao := opcaoDeEscolha()

	var palavra string

	switch opcao {
	case 1:
		fmt.Println("Você escolheu: Digitar uma nova palavra")
		fmt.Print("A palavra é:")
		fmt.Scan(&palavra)

	case 2:
		fmt.Println("Você escolheu: Palavra aleatória")
		// chamar rotina do arquivo

	default:
		fmt.Println("Desculpa, mas desconheço essa opção")
	}

	anagrama.CalculaAnagrama(palavra)
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
