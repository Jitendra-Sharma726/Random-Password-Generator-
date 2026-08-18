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
    shuffle(passwordBytes)

    return string(passwordBytes)
  }

  func evaluateStrength(pwd string) int {
    score := 1
    if len(pwd) >= 12 {
        score++
    }
    if len(pwd) >= 16 {
      score++
    }

    var flagsMet int 
    if strings.ContainsAny(pwd, upperChars) {
      flagsMet++
    }
    if strings.ContainsAny(pwd, digitChars) {
      flagsMet++
    }
    if strings.ContainsAny(pwd, specialChars) {
      flagsMet++
    }

    if flagsMet >=2 {
      score++
    }
    if flagsMet == 3 {
      score++
    }

    if score > 5 {
      return 5
    }
    return score
  }

 //---  High-Level Logic ---
 //This function handles the looping, strength evaluation, and file writing.
 func generateAndSavePasswords(count, length int, upper, digits, symbols bool, outPath string) error {
   var outputData string

   for i := 0; i < count; i++ {
     password := generatePassword(length, upper, digits, symbols)
     score := evaluateStrength(password)

     outputData += fmt.Sprintf("Password:  %-25s | Strength: %d/5\n", password, score)
    }

   err := os.WriteFile(outPath, []byte(outputData), 0644)
   if err != nil {
     // Return the error to the caller rather than killing the program here 
     return fmt.Errorf("failed to write to file %s: %w", outPath, err)
    }

   fmt.Printf("Sucessfully saved %d password(s) to '%s'\n", count,outPath)
   return nil
  }

  // --- Prompt Helpers (DO NOT MODIFY THESE FUNCTIONS)  ---

  func readString(reader *bufio.Reader, prompt string, defaultVal string) string {
    fmt.Print(prompt)
    input, _ := reader.ReadString('\n')
    input = strings.TrimSpace(input)
    if input == "" {
            return defaultVal
    }
    return input
  }


 func readInt(reader *bufio.Reader, prompt string, defaultVal int) int {
   input := readString(reader, prompt, "")
   if input == "" {
     return defaultVal
  }

 Val, err := strconv.Atoi(input)
 if err != nil {
    fmt.Printf("Invalid number, using default (%d).\n", defaultVal)
    return defaultVal
  }
  return val
  }

func readBool(reader *bufio.Reader, prompt string, defaultVal bool) bool {
  input := strings.ToLower(readString(reader, prompt, ""))
  if input == "" {
       return defaultVal
  }
  return input == "y" || input == "yes"
}
                   












    
  



















  


  









       


  


  


                   
                   

                   
                   

                   
                   
                   










                   








                   















  
