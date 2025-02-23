package main

import (
	"fmt"
	"math/rand"
	"time"

	"rsc.io/quote"
)

var defaultGreetings = []string{
	quote.Glass(),
	quote.Go(),
	quote.Hello(),
	quote.Opt(),
}

func main() {
	fmt.Println(Greeting())
}

func Greeting() string {
	return formattedTime() + ": " + randomGreeting()
}

func formattedTime() string {
	return time.Now().Format(time.RFC822)
}

func randomGreeting() string {
	return defaultGreetings[rand.Intn(len(defaultGreetings))]
}
