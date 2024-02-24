package main

import "fmt"

const french = "French"
const spanish = "Spanish" 

const englishHelloPrefix = "Hello, " 
const spanishHelloPrefix = "Hola, "
const frenchHelloPrefix = "Bonjour, "
func Hello(name string, language string) string {
  if name == "" {
    name = "World"
  }
  return greetingPrefix(language) + name
}

func greetingPrefix(language string) (prefix string) {

  switch language {
  case spanish: 
    prefix = spanishHelloPrefix
  case french: 
    prefix = frenchHelloPrefix
  default:
    prefix = englishHelloPrefix 
  }
  return 
}

func main()  {
  fmt.Println(Hello("Chris", "English"))
}