package utils

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func MsgInput() string {
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
	return sb.String()
}
