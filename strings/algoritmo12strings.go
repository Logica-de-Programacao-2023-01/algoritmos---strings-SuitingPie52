package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Informe um string:")
	scanner.Scan()
	frase := scanner.Text()

	novafrase := ""
	for _, ranFrase := range frase {
		novafrase = string(ranFrase) + novafrase
	}

	fraseM := strings.ToUpper(frase)
	novafraseM := strings.ToUpper(novafrase)

	if strings.ReplaceAll(fraseM, " ", "") == strings.ReplaceAll(novafraseM, " ", "") {
		fmt.Printf("Sua frase ''%s'' é um palíndromo.", frase)
	} else {
		fmt.Printf("Sua frase ''%s'' NÃO é um palíndromo.", frase)
	}
}
