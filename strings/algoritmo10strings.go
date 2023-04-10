package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	anagrama := false

	fmt.Println("Informe uma string:")
	scanner.Scan()
	frase1 := scanner.Text()

	fmt.Println("Informe outra string:")
	scanner.Scan()
	frase2 := scanner.Text()

	if len(frase2) == len(frase1) {
		for _, ranFrase2 := range frase2 {
			if strings.Contains(frase1, string(ranFrase2)) {
				anagrama = true
			} else {
				anagrama = false
				break
			}
		}
	}

	if frase1 == frase2 {
		fmt.Println("As frases são iguais.")
	} else if anagrama == true {
		fmt.Println("Os strings são anagramas.")
	} else {
		fmt.Println("Os strings não são anagramas.")
	}
}
