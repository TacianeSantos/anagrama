package arquivo

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func AbrirArquivo() (os.File, error) {
	arquivo, err := os.Open("palavras.txt")

	if err != nil {
		fmt.Println("Não foi possível abrir o arquivo")
		fmt.Println(err)
		os.Exit(-1)
	}

	numeroDeLinhasParaLer := 5

	LerArquivo(*arquivo, numeroDeLinhasParaLer)

	return *arquivo, err

}

func LerArquivo(arquivo os.File, numeroDeLinhasParaLer int) {

	leitor := bufio.NewReader(&arquivo)

	for {
		linha, err := leitor.ReadString('\n')
		linha = strings.TrimSpace(linha)

		fmt.Println(linha)

		if err == io.EOF {
			break
		}
	}

	arquivo.Close()
}
