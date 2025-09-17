package airportrobot

import "fmt"

// Write your code here.
// This exercise does not have tests for each individual task.
// Try to solve all the tasks first before running the tests.

type Greeter interface {
	LanguageName() string
	Greet(visitor string) string
}

type Italian struct {}

type Portuguese struct {}

func(itln Italian) LanguageName() string {
	return "Italian"
}

func(itln Italian) Greet(v string) string {
	return fmt.Sprintf("Ciao %s!",v)
}

func(prtg Portuguese) LanguageName() string {
	return "Portuguese"
}

func(prtg Portuguese) Greet(v string) string {
	return fmt.Sprintf("Olá %s!",v)
}

func SayHello(name string, langSpecGreeting Greeter) string {
	return fmt.Sprintf("I can speak %s: %s", langSpecGreeting.LanguageName(), langSpecGreeting.Greet(name))
}
