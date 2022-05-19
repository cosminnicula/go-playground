package main

import "fmt"

// // first way of doing things
// type enBot struct{}

// type esBot struct {
// }

// func main() {
// 	en := enBot{}
// 	// es := esBot{}

// 	printGreeting(en)
// 	// printGreeting(es)
// }

// func printGreeting(en enBot) {
// 	fmt.Println(en.getGreeting())
// }

// // func printGreeting(es esBot) {
// // 	fmt.Println(es.getGreeting())
// // }

// func (enBot) getGreeting() string {
// 	return "Hi!"
// }

// // func (esBot) getGreeting() string {
// // 	return "Hola!"
// // }

// second, better way of doing things
type bot interface {
	getGreeting() string
}

type enBot struct{}

type esBot struct{}

func main() {
	en := enBot{}
	es := esBot{}

	printGreeting(en)
	printGreeting(es)
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

func (enBot) getGreeting() string {
	return "Hi!"
}

func (esBot) getGreeting() string {
	return "Hola!"
}
