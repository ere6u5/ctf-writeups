
// https://hackerlab.pro/en/categories/misc/fc833637-a2b3-4cb8-b18d-b0eda09cf274
//

package main

import (
  "archive/zip"
  "encoding/base64"
  "encoding/hex"
  "flag"
  "fmt"
  "io"
  "os"
  "path/filepath"
  "strconv"
  "strings"
)

func DecodeBase64(message string, depth int) (string, error) {
  if depth > 0 {
    decodedBytes, err := base64.StdEncoding.DecodeString(message)

    if err != nil {
      return message, fmt.Errorf("Fail layer %d: %v\n", depth, err)
    }

    return DecodeBase64(string(decodedBytes), depth-1)
  }
  return message, nil
}

func StringToByte(byteString string) (byte, error) {

  if len(byteString) != 8 {
    return 0, fmt.Errorf("The number of character not equal to 8")
  }

  var number byte = 0
  for i := 0; i < 8; i++ {
    if byteString[i] == '1' {
      number |= 1 << (7 - i)
    } else if byteString[i] != '0' {
      return 0, fmt.Errorf("Invalid character at position %d: %c", i, byteString[i])
    }
  }
  return number, nil
}

func OctalStringToByte(octalNumberString string) (byte, error) {
  lenInputString := len(octalNumberString)
  if lenInputString > 3 || lenInputString < 0 {
    return 0, fmt.Errorf("The number of character less 0 or more 3")
  }

  decimalVal, err := strconv.ParseInt(octalNumberString, 8, 16)
  if err != nil {
    return 0, fmt.Errorf("Fail to convert str to int %v", err)
  }

  return byte(decimalVal), nil
}

func Unzip(src string, dest string) error {
  r, err := zip.OpenReader(src)
  if err != nil {
    return fmt.Errorf("error opening zip file: %w", err)
  }
  defer r.Close()

  for _, f := range r.File {
    filePath := filepath.Join(dest, f.Name)

    if f.FileInfo().IsDir() {
      err := os.MkdirAll(filePath, os.ModePerm)
      if err != nil {
        return fmt.Errorf("error creating directory: %w", err)
      }
      continue
    }

    err := os.MkdirAll(filepath.Dir(filePath), os.ModePerm)
    if err != nil {
      return fmt.Errorf("error creating parent directories: %w", err)
    }

    outFile, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
    if err != nil {
      return fmt.Errorf("error creating file: %w", err)
    }

    rc, err := f.Open()
    if err != nil {
      outFile.Close()
      return fmt.Errorf("error opening zipped file content: %w", err)
    }

    _, err = io.Copy(outFile, rc)

    outFile.Close()
    rc.Close()

    if err != nil {
      return fmt.Errorf("error copying file content: %w", err)
    }
  }

  return nil
}

func main() {

  inputFilePath := flag.String("i", "task_pug.dat", "Input path to file")
  depth := flag.Int("d", 1, "Number of iteration iteration to decode")
  tmpArchive := flag.String("z", "output.zip", "Output path to file")
  outputPathToPng := flag.String("op", "dog.png", "Output png picture")
  outputPathToFirstPng := flag.String("opf", "dog1.png", "Out second png picture")
  outputPathToSecondPng := flag.String("ops", "dog2.png", "Out second png picture")

  flag.Parse()

  inputData, err := os.ReadFile(*inputFilePath)
  if err != nil {
    fmt.Printf("Error reading file: %v\n", err)
    os.Exit(1)
  }

  outputFile, err := os.Create(*tmpArchive)
  if err != nil {
    fmt.Printf("Error create file: %v\n", err)
    os.Exit(1)
  }
  defer outputFile.Close()

  outputPng, err := os.Create(*outputPathToPng)
  if err != nil {
    fmt.Printf("Error create file: %v\n", err)
    os.Exit(1)
  }
  defer outputPng.Close()

  outputPngFirst, err := os.Create(*outputPathToFirstPng)
  if err != nil {
    fmt.Printf("Error create file: %v\n", err)
    os.Exit(1)
  }
  defer outputPngFirst.Close()

  outputPngSecond, err := os.Create(*outputPathToSecondPng)
  if err != nil {
    fmt.Printf("Error create file: %v\n", err)
    os.Exit(1)
  }
  defer outputPngSecond.Close()

  clearedInput := strings.TrimSpace(string(inputData))
  encodedBytesString, err := DecodeBase64(clearedInput, *depth)
  if err != nil {
    fmt.Printf("Error decode base64: %v\n", err)
    os.Exit(1)
  }



  sliceEncodedBytes := []byte{}
  sliceBytesToString := strings.Split(encodedBytesString, " ")
  for i, seqBytes := range sliceBytesToString {
    tmpBuffer, err := StringToByte(seqBytes)
    if err != nil {
      fmt.Printf("Error at segment %d: %v\n", i, err)
      os.Exit(1)
    }
    sliceEncodedBytes = append(sliceEncodedBytes, tmpBuffer)
  }

  cleanedEncodedOctalToStrings := strings.Split(string(sliceEncodedBytes), " ")
  zipArchiveBytes := []byte{}
  for _, octalStr := range cleanedEncodedOctalToStrings {
    byteFromOctal, err := OctalStringToByte(octalStr)
    if err != nil {
      fmt.Printf("Error decode octal: %v\n", err)
      os.Exit(0)
    }
    zipArchiveBytes = append(zipArchiveBytes, byteFromOctal)
  }

  _, err = outputFile.Write(zipArchiveBytes)
  if err != nil {
    fmt.Printf("Error writing to file: %v\n", err)
  }

  err = Unzip(*tmpArchive, "./")
  if err != nil {
    fmt.Printf("Error unzip file:%v\n", err)
  }

  encodedPngInString, err := os.ReadFile("./file.txt")
  if err != nil {
    fmt.Printf("Error open file: %v\n", err)
  }

  bytesPngInString := strings.ReplaceAll(string(encodedPngInString), " ", "")

  imageDataBytes, err := hex.DecodeString(bytesPngInString)

  _, err = outputPng.Write(imageDataBytes)
  if err != nil {
    fmt.Printf("Error writing to file: %v\n", err)
  }

  firstPng := []byte{}
  secondPng := []byte{}
  flagSecondImage := false
  for i := 0; i < len(imageDataBytes); i++ {

    if flagSecondImage {
      secondPng = append(secondPng, imageDataBytes[i])
      continue
    }

    if i > 7 {
      if imageDataBytes[i-7] == 0x49 && imageDataBytes[i-6] == 0x45 && imageDataBytes[i-5] == 0x4e && imageDataBytes[i-4] == 0x44 {
        if imageDataBytes[i-3] == 0xae && imageDataBytes[i-2] == 0x42 && imageDataBytes[i-1] == 0x60 && imageDataBytes[i] == 0x82 {
          flagSecondImage = true
        }
      }
    }

    firstPng = append(firstPng, imageDataBytes[i])
  }

  _, err = outputPngFirst.Write(firstPng)
  if err != nil {
    fmt.Printf("Error writing to file: %v\n", err)
  }

  _, err = outputPngSecond.Write(secondPng)
  if err != nil {
    fmt.Printf("Error writing to file: %v\n", err)
  }

  os.Exit(0)
}
