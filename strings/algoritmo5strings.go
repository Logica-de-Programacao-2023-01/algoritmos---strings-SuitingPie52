package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Informe uma string:")
	scanner.Scan()
	x := scanner.Text()

	y, err := strconv.ParseFloat(x, 64)

	if err != nil {
		fmt.Printf("Seu string %f não é float.", y)
	} else {
		fmt.Printf("Seu string %f é float.", y)
	}
}
