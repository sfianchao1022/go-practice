package main

// go env -w GO111MODULE=off, on, auto
// go mod 和 GOPATH 不能同時使用
// go mod 要使用絕對位置到專案最上層 eg. import "tutorial/hello"
// GOPATH 才能使用絕對位置 eg. import "./hello"

import (
	"tutorial/hello"
)

func main() {
	hello.SayHello()
}
