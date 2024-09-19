package main

import (
	"github.com/edokimok/golang-boilerplate/helloworld/hello"
	hello2 "github.com/edokimok/golang-boilerplate/helloworld_hello2"
)

func main() {
	println("hello: " + hello.HelloWorld())
	println("hello2: " + hello2.HelloWorld())
}
