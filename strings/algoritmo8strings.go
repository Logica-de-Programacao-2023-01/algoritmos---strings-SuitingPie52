package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	novo := ""
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Informe um string:")
	scanner.Scan()
	x := scanner.Text()

	for _, y := range x {
		novo = string(y) + novo
	}

	fmt.Println("Seu string ao contrário é: ", novo)
}
