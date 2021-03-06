package hello

import (
	"fmt"
	"tutorial/greeting"

	"rsc.io/quote"
)

func SayHello() {
	fmt.Println(Hello())
	fmt.Println("hello world")
	greeting.Greeting()
}

func Hello() string {
	return quote.Hello()
}
