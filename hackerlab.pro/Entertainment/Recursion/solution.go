// https://hackerlab.pro/en/categories/misc/37f1dca8-b855-4993-8673-160e5d1b8bfe
// CODEBY{l0ng_r3curs1on}

package main

import (
  "encoding/base64"
  "flag"
  "fmt"
  "os"
  "strings"
)

func DecodeBase64(message string, depth int) (string, error) {
  if depth > 0 {
    fmt.Printf("[%5d] Layer: %.20s...\n", depth, message)
    decodedBytes, err := base64.StdEncoding.DecodeString(message)

    if err != nil {
      fmt.Printf("Fail layer %d: %v\n", depth, err)
      fmt.Printf("End layer:\n%s\n", message)
      return message, fmt.Errorf("Fail layer %d: %v\n", depth, err)
    }

    return DecodeBase64(string(decodedBytes), depth-1)
  }
  fmt.Printf("[%5d] Last layer :]\n%s\n\n", depth, message)
  return message, nil
}

func main() {

  inputFilePath := flag.String("i", "", "Input path to file")
  depth := flag.Int("d", 100, "Number of iteration iteration to decode")

  flag.Parse()

  inputData, err := os.ReadFile(*inputFilePath)
  if err != nil {
    fmt.Println("Error reading file:", err)
    os.Exit(1)
  }

  clearedInput := strings.TrimSpace(string(inputData))
  _, err = DecodeBase64(clearedInput, *depth)

  if err != nil {
    os.Exit(1)
  }

  os.Exit(0)
}
