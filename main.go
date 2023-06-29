package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	input := ""
	for !strings.Contains(input, "konec") {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Nasnímejte prosím kartu, nebo čip (zadejte konec pro ukončení programu): ")
		input, _ = reader.ReadString('\n')
		fmt.Println("Nasnímaná data: " + input)
	}
}
