package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Informe um string:")
	scanner.Scan()
	frase := scanner.Text()
	fmt.Println("Informe um substring para remover:")
	scanner.Scan()
	del := scanner.Text()
	fmt.Println("Informe uma frase para por no lugar:")
	scanner.Scan()
	novo := scanner.Text()

	frase_nova := strings.ReplaceAll(frase, del, novo)

	fmt.Println("Sua nova frase é:", frase_nova)
}
