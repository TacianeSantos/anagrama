package arquivo

import (
	"bufio"
	"fmt"
	"os"
)

func AbrirArquivo() (os.File, error) {
	arquivo, err := os.Open("palavras.txt")

	if err != nil {
		fmt.Println("Não foi possível abrir o arquivo")
		fmt.Println(err)
		os.Exit(-1)
	}

	LerArquivo(*arquivo)

	return *arquivo, err

}

func LerArquivo(arquivo os.File) {
	leitura := bufio.NewReader(&arquivo)

	fmt.Println(leitura.ReadString('\n'))
}
