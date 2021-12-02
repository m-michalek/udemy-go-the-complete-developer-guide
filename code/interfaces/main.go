package main

import "fmt"

type bot interface {
	getGreeting() string
}

type englishBot struct{}

func (englishBot) getGreeting() string {
	return "Hello!"
}

type spanishBot struct{}

func (spanishBot) getGreeting() string {
	return "Hola!"
}

func printGreeting(b bot) {
	fmt.Printf("greeing: %v\n", b.getGreeting())
}

type dog struct{}

func (dog) getGreeting() string {
	return "Woof!"
}

func main() {
	eb := englishBot{}
	sb := spanishBot{}
	d := dog{}

	printGreeting(eb)
	printGreeting(sb)
	printGreeting(d)
}
