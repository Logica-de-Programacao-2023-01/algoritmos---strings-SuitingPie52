package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Informe uma frase:")
	scanner.Scan()
	frase := scanner.Text()

	temNumero := false
	for x := 0; x <= 9; x++ {
		s := strconv.Itoa(x)
		if strings.ContainsAny(frase, s) == true {
			temNumero = true
			break
		}
	}

	if temNumero == true {
		fmt.Println("Sua frase tem um número.")
	} else {
		fmt.Println("Sua frase não tem um número.")
	}
}
