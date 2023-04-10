package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Informe uma sequência de números:")
	scanner.Scan()
	sequencia := scanner.Text()

	crescente := false
	for i := 1; i < len(sequencia); i++ {
		if sequencia[i-1] == sequencia[i]-1 {
			crescente = true
		} else {
			crescente = false
			break
		}
	}

	if crescente {
		fmt.Println("Sua sequência é uma sequência crescente.")
	} else {
		fmt.Println("Sua sequência NÃO é uma sequência crescente.")
	}
}
