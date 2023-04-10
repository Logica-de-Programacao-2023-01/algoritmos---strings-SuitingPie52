package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Digite uma frase: ")
	scanner.Scan()
	frase := scanner.Text()
	palavras := strings.Fields(frase)

	fmt.Printf("Suas palavras são: %s, e são %d palavras.", palavras, len(palavras))
}
