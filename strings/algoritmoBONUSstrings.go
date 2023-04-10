package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	indexes := []int{}
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Informe um string:")
	scanner.Scan()
	frase := scanner.Text()

	fmt.Println("Informe um padrão:")
	scanner.Scan()
	padrao := scanner.Text()

	lastIdx := strings.LastIndex(frase, padrao)
	for lastIdx != -1 {
		indexes = append([]int{lastIdx}, indexes...)
		frase = frase[:lastIdx]
		lastIdx = strings.LastIndex(frase, padrao)
	}

	fmt.Println("O padrão ocorre nos índices:", indexes)
}
