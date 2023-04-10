package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Informe uma string:")
	scanner.Scan()
	frase1 := scanner.Text()

	fmt.Println("Informe outra string:")
	scanner.Scan()
	frase2 := scanner.Text()

	if strings.Contains(frase1, frase2) {
		fmt.Println("Sua segunda string é uma substring da primeira.")
	} else {
		fmt.Println("Sua segunda string NÃO é uma substring da primeira.")
	}
}
