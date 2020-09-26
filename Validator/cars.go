package main

import (
	"fmt"
	"log"

	"github.com/go-playground/validator"
)

type car struct {
	Make  string
	Model string
}

// The function below is otherwise known as an annonymous function, or functional literal
// It contains two key componenents - The parameter values and the return type
// The return type (Val()) is a method signature, which sort of "attaches" to the Car type, which
// you see on line 36

func (c *car) Val() {
	fmt.Println("validating...")

	val := validator.New()
	validate := val.Struct(c)

	if validate != nil {
		log.Println("error on validate of type: car")
	}
}

func main() {
	v := new(car)
	v.Val()
}
