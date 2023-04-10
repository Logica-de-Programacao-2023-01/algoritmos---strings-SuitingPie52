package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Informe uma string.")
	scanner.Scan()
	x := scanner.Text()
	y := strings.ToUpper(x)
	fmt.Printf("Sua string é: %s\n", y)
}
