package input

import (
	"bufio"
	"fmt"
	"os"
)

func Int(text string) int {
	var word int
	fmt.Printf("%s", text)
	fmt.Scanln(&word)
	return word
}

func Float(text string) float64 {
	var word float64
	fmt.Printf("%s", text)
	fmt.Scanln(&word)
	return word
}

func Str(textF string) string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("%s", textF)
	text, _ := reader.ReadString('\n')
	return text
}

func Empty() {
	var a int
	fmt.Scanln(&a)
}
