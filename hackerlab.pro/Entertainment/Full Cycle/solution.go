// https://hackerlab.pro/en/categories/misc/f63aef59-cfa2-4351-9f6b-8ecbc16c5300
// CODEBY{Y0U'R3_C00L_47_S73G_BR0}

package main

import (
	"fmt"
	"os"
)

func main() {
	inputFilePath := "input.jpg"
	outputFilePath := "output.jpg"

	contentBytes, err := os.ReadFile(inputFilePath)
	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		os.Exit(1)
	}

	outputFile, err := os.Create(outputFilePath)
	if err != nil {
		fmt.Printf("Error creating file: %v\n", err)
		os.Exit(1)
	}
	defer outputFile.Close()

	lenBytes := len(contentBytes)
	for i := 0; i < lenBytes/2; i++ {
		tempBuffer := contentBytes[i]
		contentBytes[i] = contentBytes[lenBytes-1-i]
		contentBytes[lenBytes-1-i] = tempBuffer
	}

	_, err = outputFile.Write(contentBytes)
	if err != nil {
		fmt.Printf("Error writing to file: %v\n", err)
		os.Exit(1)
	}

	os.Exit(0)
}
