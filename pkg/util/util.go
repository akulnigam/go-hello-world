package util

import "errors"

const HelloWorldMessage = "Hello World from Constant"

func GetMessage() (string, int) {
	message := "Hello World"
	number := 1
	return message, number
}

func CreateMessage(x int) (string, int, error) {
	message := "Hello World"

	if x<0{
		return "",0, errors.New("x cannot be less than 0")
	}

	number := x
	return message, number, nil
}