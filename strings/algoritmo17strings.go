package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	count := 0

	fmt.Println("Informe uma string.")
	scanner.Scan()
	frase := scanner.Text()

	fmt.Print("Suas letras únicas são: ")
	for _, ranFrase := range frase {
		count = strings.Count(frase, string(ranFrase))
		if count == 1 {
			fmt.Print(string(ranFrase))
		}
	}
}
