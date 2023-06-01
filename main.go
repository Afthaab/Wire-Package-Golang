package main

import "fmt"

type Message string

type Greeter struct {
	Message Message
}

type Event struct {
	Greeter Greeter
}

func GetMessage() Message {
	return Message("Hello World")
}

func GetGreeter(m Message) Greeter {
	return Greeter{
		Message: m,
	}
}

func GetEvent(g Greeter) Event {
	return Event{
		Greeter: g,
	}
}
func (g Greeter) greet() Message {
	return g.Message
}

func (e Event) Start() {
	msg := e.Greeter.greet()
	fmt.Println(msg)
}

func main() {
	event := InitializeEvent()
	event.Start()
}
