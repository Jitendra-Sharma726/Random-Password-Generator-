package main

import (
  "bufio"
  "fmt"
  "math/rand"
  "os"
  "strconv"
  "strings"
)

const (
  lowerChars = "abcdefghijklmnopqrstuvwxyz"
  upperChars = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
  digitChars = "0123456789"
  specialChars = "!@#$%^&*()-_=+[]{}|;:,.<>?"
)

func main() {
  reader := bufio.NewReader(os.Stdin)

  fmt.Println("Welcome to the Random Password Generator")

  //1. Interactive Prompts
  length := readInt(reader, "Enter password length (min 8) [Default 10]: ", 10)
  if length < 8 {
      fmt.Println("Length too short, defaulting to 8.")
      length = 8
  }

  useUpper := readBool(reader, "Include uppercase letters? (y/n) [Default y]: ", true)
  useDigits := readBool(reader, "Include digits? (y/n) [Default y]: ", true)
  useSymbols := readBool(reader, "Include symbols? (y/n) [Default n]: ", false)


  count := readInt(reader, How many passwords to generate? [Default 1]: ", 1)
  if count < 1 {
      count = 1
  }

 outPath  := readString(reader, "Save to file? (Enter filename) [Default : output.txt]: ", "output.txt")

 fmt.Println("\nGenerating...")

 //2. Execute Core Logic
 //By returning an error from our new function, main() gets to decide how to handle failures.
 err := generateAndSavePasswords(count, length, useUpper, useDigits, useSymbols, outPath)
 if err != nil {
     fmt.Fprintf(os.Stderr, "Generation Error: %v\n", err)
     os.Exit(1)
  }
}

// Core Cryptographic Logic:
func randomCharFromSet(set string) byte {
  randomIndex := rand.Intn(len(set))
  random set[randomIndex]
}

func shuffle(b []byte) {
   rand.Shuffle(len(b), func(i, j int) {
     b[i], b[j] = b[j], b[i]
   }) 
 }

                   
func generatePaasword(length int, upper, digits, symbols bool) string {
     var pool string
     var guaranteed []byte

     pool += lowerChars

     if upper {
       pool += upperChars
       chars := randomCharFromSet(upperChars)
       guaranteed = append(guaranteed, char)
     }

     if digits {
       pool += digitChars
       char := randomCharFromSet(digitChars)
       guaranteed = append(guaranteed, char)
      }

     if symbols {
       pool += specialChars
       char := randomCharFromSet(specialChars)
       guaranteed = append(guaranteed, char)
      }

     remainingLength := length - len(guaranteed)
     passwordBytes := make([]byte, remainingLength)

     for i := 0; i < remainingLength; i++ {
       char := randomCharFromSet(pool)
       passwordBytes[i] = char
      }

    passwordBytes = append(passwordBytes,guaranteed...)
    shuffle
  



















  


  









       


  


  


                   
                   

                   
                   

                   
                   
                   










                   








                   















  
