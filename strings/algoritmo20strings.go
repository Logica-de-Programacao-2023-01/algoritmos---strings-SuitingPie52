package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	camelCase := false
	contador := 1

	fmt.Println("Informe uma string:")
	scanner.Scan()
	frase := scanner.Text()

	fraseM := strings.ToUpper(frase)

	if fraseM[0] != frase[0] {
		camelCase = true
	}

	for _, ranFrase := range frase {
		ranFraseM := strings.ToUpper(string(ranFrase))
		if string(ranFrase) == ranFraseM {
			contador++
		}
	}

	if camelCase {
		fmt.Printf("Seu string é camel case e tem %d palavras.", contador)
	} else {
		fmt.Println("Seu string não é camel case.")
	}
}
