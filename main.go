package main

import (
	"fmt"
	"github.com/Pallinder/go-randomdata"
)


func main() {
	fmt.Println("Println is exported because it is capitalized")
 randomdata.seedAndReturnRandom(42) //not exported because it starts with a lower case letter
}
		