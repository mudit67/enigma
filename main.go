package main

import (
	"bufio"
	"enigma/core"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("Intializing Engima Machine")
	engima := core.InitEnigma()

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter text: ")
	text, _ := reader.ReadString('\n')
	fmt.Println(text)
	var sb strings.Builder
	for _, c := range text {
		if c >= 'A' && c <= 'Z' {
			sb.WriteRune(c)
		}
	}
	fmt.Println(engima.Cipher(sb.String()))

	// fmt.Println(engima.Cipher("ABC"))
}
