package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Informe uma string:")
	scanner.Scan()
	x := scanner.Text()
	y := strings.ReplaceAll(x, " ", "")
	fmt.Println("Sua nova string é: ", y)
}
