package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Informe uma string: ")
	scanner.Scan()
	frase := scanner.Text()
	fmt.Println("Informe uma letra para remover: ")
	scanner.Scan()
	x := scanner.Text()
	fmt.Println("Informe uma letra para por no lugar.")
	scanner.Scan()
	y := scanner.Text()

	novafrase := strings.ReplaceAll(frase, x, y)
	fmt.Printf("Seu novo string é: %s", novafrase)
}
