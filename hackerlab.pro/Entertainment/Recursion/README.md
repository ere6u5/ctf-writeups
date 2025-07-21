**🔍 Lab analysis Recursion Challenge on HackerLab**
__([Entertainment: Recoursion](https://hackerlab.pro/en/categories/misc/37f1dca8-b855-4993-8673-160e5d1b8bfe))__

**🛠 Tools Used**

  - Go (Programming language)
  - `tail` (Linux utility)

**📌 Objective**

Decrypt a base64-encoded file with multiple layers of encoding.

**🔧 Solution steps**

  1. Used tail to inspect the file and identified = as the padding charecter, confirming base64 encoding.
```bash
tail -c 5 ./recursion.txt
```

  2. Wrote recursive Go program to decode 50 layers of base64.
```go
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
```

**🎯 Result**

The program successfully decrypted the message after processing all layers.

[View full solution on GitHub](https://github.com/ere6u5/ctf-writeups/blob/main/hackerlab.pro/Entertainment/Recursion/README.md)

**#Programming #Linux #HackerLab #Golang**
