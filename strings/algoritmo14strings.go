package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Informe uma sequencia de números.")
	scanner.Scan()
	sequencia := scanner.Text()

	decrescente := false
	for i := 1; len(sequencia) > i; i++ {
		if sequencia[i-1] == sequencia[i]+1 {
			decrescente = true
		} else {
			decrescente = false
			break
		}
	}

	if decrescente {
		fmt.Println("Sua sequência é uma sequência decrescente.")
	} else {
		fmt.Println("Sua sequência NÃO é uma sequência decrescente.")
	}
}
