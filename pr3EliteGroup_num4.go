package main

import (
 "bufio"
 "fmt"
 "os"
 "strings"
)

func longestCommonPrefix(strs []string) string {
 if len(strs) == 0 {
  return ""
 }
 prefix := strs[0] 
 for _, str := range strs[1:] {
  for prefix != "" && !strings.HasPrefix(str, prefix) {
   prefix = prefix[:len(prefix)-1] 
  }
 }
 return prefix
}

func main() {
 reader := bufio.NewReader(os.Stdin)
 fmt.Println("Введите слова: ")

 var words []string
 for {
  input, _ := reader.ReadString('\n')
  input = strings.TrimSpace(input)
  if input == "" {
   break
  }
  words = append(words, input)
 }

 prefix := longestCommonPrefix(words)
 if prefix == "" {
  fmt.Println("Нет общего префикса.")
 } else {
  fmt.Printf("Самый длинный общий префикс: %s\n", prefix)
 }
}
