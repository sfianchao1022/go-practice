package main

// echo $GOPATH $GO111MODULE
// 設置GOPATH: export GOPATH=$HOME/go
// 設置go mod:
// [linux/mac]vim ~/.zshrc export GO111MODULE=on source ~/.zshrc
// [windows]set GO111MODULE=on
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
