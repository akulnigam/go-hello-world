package main

import (
	"github.com/go-hello-world/pkg/util"
	"log"
)

func main() {

	//var message string = "Hello World"
	//message := "Hello World"
	if m, v, error := util.CreateMessage(-2); error != nil{
		log.Fatalf("%v", error)
	} else {
		log.Printf("%v,%v\n", m, v)
	}

	//fmt.Printf("%v\n", util.HelloWorldMessage)

	//log.Printf("%v,%v", m, v)
	//println("Hello World!!!!")
}

