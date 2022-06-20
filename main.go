package main

import (
	"bufio"
	"fmt"
	"os"
)

func StringReduction(str string) int {

	fmt.Println("output", len(str))
	return len(str)
}

func main() {
	fmt.Printf("Output: %d", StringReduction(readline()))
}

func readline() string {
	fmt.Printf("Intput: ")
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	return text
}
