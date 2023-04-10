package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Informe uma string.")
	scanner.Scan()
	numeros := scanner.Text()

	x, err := strconv.ParseFloat(numeros, 64)

	if err != nil {
		fmt.Printf("Sua string não contém apenas números.")
	} else {
		fmt.Printf("Sua string %f contém apenas números.", x)
	}
}
