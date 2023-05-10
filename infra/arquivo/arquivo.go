package arquivo

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func abrirArquivo() (os.File, error) {
	arquivo, err := os.Open("palavras.txt")

	if err != nil {
		fmt.Println("Não foi possível abrir o arquivo")
		fmt.Println(err)
		os.Exit(-1)
	}

	return *arquivo, err

}

func LerPalavraDoArquivo(numeroDeLinhasParaLer int) []string {

	arquivo, _ := abrirArquivo()
	leitor := bufio.NewReader(&arquivo)

	var palavras []string

	for i := 0; i < numeroDeLinhasParaLer; i++ {
		linha, err := leitor.ReadString('\n')
		linha = strings.TrimSpace(linha)

		palavras = append(palavras, linha)

		if err == io.EOF {
			break
		}
	}

	arquivo.Close()

	return palavras
}
