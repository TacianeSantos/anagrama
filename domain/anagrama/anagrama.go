package anagrama

import (
	"fmt"
)

func CalculaAnagrama(palavra string) []string {

	var umSliceVazio []string

	return umSliceVazio
}

func ImprimeAnagrama(resultadoAnagramas []string) {
	if resultadoAnagramas == nil {
		fmt.Println("Nenhum anagrama encontrado")
	} else {
		fmt.Println("Quantidade de anagramas é: ", len(resultadoAnagramas))
	}

	for anagrama := range resultadoAnagramas {
		fmt.Println(anagrama)

	}
}
