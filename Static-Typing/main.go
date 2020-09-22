package main

import "fmt"

type car struct {
	model string
	year  string
	stock int
}

func main() {
	myCars()

}

func myCars() *car {
	cars := car{}
	cars.model = "Ford"

	fmt.Println(cars.model)

	return nil
}
