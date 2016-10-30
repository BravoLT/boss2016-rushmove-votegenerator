package main

import (
	"fmt"
	"strconv"
)


func main() {
	fmt.Println("The sum is " + strconv.Itoa(add( 2, 2)))

	out, err := NewNumberPairPtr(1,2).Add()
	if err != nil { //error handling, no exceptions
		fmt.Println(err)
	}
	fmt.Println("The sum is " + strconv.Itoa(out))

	//error handling while tightly scoping variables
	if out, err1 := NewNumberPairPtr(2,2).Add(); err1 != nil {
		fmt.Println(err1)
	} else {
		fmt.Println("The sum is " + strconv.Itoa(out))
	}


	// this call fails because the Add() function was added to the pointer and not a value.
	_, err = NewNumberPairValue(1,2).Add()

}

//functions can be stand alone
// when consecutive arguments have the same type, you can just put it on the last one
func add(x, y int) int {
	return x + y
}

//A constructor example returning a value
func NewNumberPairValue(a int, b int) NumberPair {
	pair := new(NumberPair)
	pair.a = a
	pair.b = b
	return pair
}

//A constructor example returning a pointer
func NewNumberPairPtr(a int, b int) *NumberPair {
	pair := new(NumberPair)
	pair.a = a
	pair.b = b
	return pair
}

//Structs are similar to objects in Java but do not have behavior included in the struct definition
type NumberPair struct {
	a int
  b int
}

//functions can also be added to a struct or pointer, does not automatically dereference to the struct, so this method
// must be called on a pointer to a NumberPair and not a NumberPair value.
//Also, functions can have multiple return values
func (n *NumberPair) Add() (int, error) {
	//member functions have access to unexported variables and functions as you would expect.
	return n.a + n.b
}

