package main

import "fmt"

type gasEngine struct {
	mpg     uint8
	gallons uint8
	owner
}

type electricEngine struct {
	mpkwh uint8
	kwh   uint8
}

type owner struct {
	name string
}

type engine interface {
	milesLeft() uint8
}

func (e electricEngine) milesLeft() uint8 {
	return e.kwh * e.mpkwh
}

func (e gasEngine) milesLeft() uint8 {
	return e.gallons * e.mpg
}

func canMakeIt(e engine, miles uint8) {
	if miles <= e.milesLeft() {
		fmt.Println("all good")
	} else {
		fmt.Println("no good")
	}
}
func main() {
	var myEngine gasEngine = gasEngine{25, 15, owner{"alex"}}
	myEngine.mpg = 20
	fmt.Println(myEngine.gallons, myEngine.name, myEngine.mpg)
	fmt.Printf("Total miles left in tank: %v\n", myEngine.milesLeft())
	canMakeIt(myEngine, 50)
	var myEV electricEngine = electricEngine{20, 20}
	canMakeIt(myEV, 30)
}
