package airportrobot

import "fmt"

// Write your code here.
// This exercise does not have tests for each individual task.
// Try to solve all the tasks first before running the tests.

type Greeter interface {
	LanguageName() string
	Greet(visitorName string) string
}

type Italian struct{}

func (italianGreeter Italian) LanguageName() string {
	return "Italian"
}

func (italianGreeter Italian) Greet(visitorName string) string {
	return fmt.Sprintf("Ciao %s!", visitorName)
}

type Portuguese struct{}

func (portugueseGreeter Portuguese) LanguageName() string {
	return "Portuguese"
}

func (italianGreeter Portuguese) Greet(visitorName string) string {
	return fmt.Sprintf("Olá %s!", visitorName)
}

func SayHello(visitorName string, greeter Greeter) string {
	return fmt.Sprintf("I can speak %s: %s", greeter.LanguageName(), greeter.Greet(visitorName))
}
