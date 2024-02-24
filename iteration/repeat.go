package iteration

import "strings"

func Repeat(letter string, repeatCount int) string {
  // str := ""
  // for i:=0; i < repeatCount; i++ {
  //   str += letter 
  // }  
  //return str
  return strings.Repeat(letter, repeatCount)
}
