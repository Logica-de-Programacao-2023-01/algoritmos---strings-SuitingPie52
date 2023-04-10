package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Informe uma string:")
	scanner.Scan()
	x := scanner.Text()
	fmt.Println("Informe outra string:")
	scanner.Scan()
	y := scanner.Text()

	if x == y {
		fmt.Println("Suas strings são iguais.")
	} else {
		fmt.Println("Suas strings são diferentes.")
	}
}
