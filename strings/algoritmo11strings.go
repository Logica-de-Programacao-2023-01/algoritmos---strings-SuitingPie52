package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Informe uma string.")
	scanner.Scan()
	frase := scanner.Text()

	frase = strings.ReplaceAll(frase, "a", "")
	frase = strings.ReplaceAll(frase, "e", "")
	frase = strings.ReplaceAll(frase, "i", "")
	frase = strings.ReplaceAll(frase, "o", "")
	frase = strings.ReplaceAll(frase, "u", "")
	frase = strings.ReplaceAll(frase, "A", "")
	frase = strings.ReplaceAll(frase, "E", "")
	frase = strings.ReplaceAll(frase, "I", "")
	frase = strings.ReplaceAll(frase, "O", "")
	frase = strings.ReplaceAll(frase, "U", "")
	frase = strings.ReplaceAll(frase, "á", "")
	frase = strings.ReplaceAll(frase, "é", "")
	frase = strings.ReplaceAll(frase, "í", "")
	frase = strings.ReplaceAll(frase, "ó", "")
	frase = strings.ReplaceAll(frase, "ú", "")
	frase = strings.ReplaceAll(frase, "Á", "")
	frase = strings.ReplaceAll(frase, "É", "")
	frase = strings.ReplaceAll(frase, "Í", "")
	frase = strings.ReplaceAll(frase, "Ó", "")
	frase = strings.ReplaceAll(frase, "Ú", "")
	frase = strings.ReplaceAll(frase, "â", "")
	frase = strings.ReplaceAll(frase, "ê", "")
	frase = strings.ReplaceAll(frase, "î", "")
	frase = strings.ReplaceAll(frase, "ô", "")
	frase = strings.ReplaceAll(frase, "û", "")
	frase = strings.ReplaceAll(frase, "Â", "")
	frase = strings.ReplaceAll(frase, "Ê", "")
	frase = strings.ReplaceAll(frase, "Î", "")
	frase = strings.ReplaceAll(frase, "Ô", "")
	frase = strings.ReplaceAll(frase, "Û", "")
	frase = strings.ReplaceAll(frase, "ã", "")
	frase = strings.ReplaceAll(frase, "õ", "")
	frase = strings.ReplaceAll(frase, "Ã", "")
	frase = strings.ReplaceAll(frase, "Õ", "")

	fmt.Println("Sua string sem as vogais é: ", frase)
}
