package main

import (
	"fmt"
	"os"
	"time"

	"example.com/hello/mocking"
)

const (
	spanish = "Spanish"
	french  = "French"
	polish  = "Polish"

	englishHelloPrefix = "Hello, "
	spanishHelloPrefix = "Hola, "
	frenchHelloPrefix  = "Bonjour, "
	polishHelloPrefix  = "Cześć, "
)

func Hello(name string, language string) string {
	if name == "" {
		name = "world"
	}

	return greetingPrefix(language) + name
}

func greetingPrefix(language string) (prefix string) {
	switch language {
	case spanish:
		prefix = spanishHelloPrefix
	case french:
		prefix = frenchHelloPrefix
	case polish:
		prefix = polishHelloPrefix
	default:
		prefix = englishHelloPrefix
	}

	return
}

func main() {
	fmt.Println(Hello("world", ""))
	sleeper := mocking.NewConfigurableSleeper(2 * time.Second, time.Sleep)
	mocking.Countdown(os.Stdout, &sleeper)
}
