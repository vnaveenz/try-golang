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
	return "I can speak Italian:"
}

func(itln Italian) Greet(v string) string {
	return fmt.Sprintf("%s Ciao %s!", itln.LanguageName() ,v)
}

func(prtg Portuguese) LanguageName() string {
	return "I can speak Portuguese:"
}

func(prtg Portuguese) Greet(v string) string {
	return fmt.Sprintf("%s Olá %s!", prtg.LanguageName() ,v)
}

func SayHello(name string, langSpecGreeting Greeter) string {
	return langSpecGreeting.Greet(name)
}
