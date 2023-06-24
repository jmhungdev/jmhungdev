package main

import "fmt"

func textToHex(text string) {
	for _, char := range text {
		fmt.Printf("%X ", int(char))
	}
	fmt.Println()
}

func textToASCII(text string) {
	for _, char := range text {
		fmt.Printf("%d ", int(char))
	}
	fmt.Println()
}

// convert hex -> decimal and then decimal -> text

func hexToText(hex []string) {
	for _, code := range hex {
		var asciiCode int
		fmt.Sscanf(code, "%X", &asciiCode)
		fmt.Printf("%c", asciiCode)
	}
	fmt.Println()
}

func asciiToText(ascii []int) {
	for _, code := range ascii {
		fmt.Printf("%c", code)
	}
	fmt.Println()
}

func main() {
	// Convert text to hexadecimal
	text := "Hello, World! 你好"
	fmt.Println("Text to Hexadecimal:")
	textToHex(text)

	fmt.Println("Text to decimal:")
	textToASCII(text)

	// Convert hexadecimal to text
	hex := []string{"48", "65", "6C", "6C", "6F", "2C", "20", "57", "6F", "72", "6C", "64", "21"}
	fmt.Println("Hexadecimal to Text:")
	hexToText(hex)
}

// output:
// Text to Hexadecimal:
// 48 65 6C 6C 6F 2C 20 57 6F 72 6C 64 21 20 4F60 597D
// Text to decimal:
// 72 101 108 108 111 44 32 87 111 114 108 100 33 32 20320 22909
// Hexadecimal to Text:
// Hello, World!
