package main

import (
	"fmt"
	"pubsub"
	"strings"
	"time"
)

func main() {
	p := pubsub.NewPubliser(100 * time.Millisecond, 10)
	defer p.Close()

	all := p.Subscribe()
	golang := p.SubscribeTopic(func(v interface{}) bool {
		if s, ok := v.( string); ok {
			return strings.Contains(s, "golang")
		}
		return false
	})

	p.Publish("Hello, world!")
	p.Publish("hello, golang!")

	go func() {
		for msg := range all {
			fmt.Println("all: ", msg)
		}
	}()

	go func() {
		for msg := range golang {
			fmt.Println("golang: ", msg)
		}
	}()

	time.Sleep(3 * time.Second)
}
