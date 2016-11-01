package main

import (
	"time"
	"fmt"
	"strconv"
	"github.com/Pallinder/go-randomdata"
)

func main() {

	//read config
	//

}

// An example of how you would implement an enum in Go
type Party string
const (
	VIM Party = "vim"
	EMACS Party = "emacs"
	SPACES Party = "spaces"
	TABS Party = "tabs"
)

type Race int
const (
	BLUE Race = "blue"
	GREEN Race = "green"
	PURPLE  Race = "purple"
	ORANGE Race = "orange"
)

type BallotOption int
const (
	MOVE BallotOption = iota
	REMAIN
)


type Location struct {
	name string
	population int
	age int
	householdIncome int
	party Party 		//the most popular party in the location
	racialBreakdown map[Race] float32
}



type Ballot struct {
	choice BallotOption
	location *Location
	time time.Time
}
