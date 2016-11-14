package main

import (
	"time"
	"os"
	"fmt"
	_ "strconv"
	"github.com/Pallinder/go-randomdata"
	"github.com/BurntSushi/toml"
	"log"
)

var location *Location

func main() {
	//read config
  location = initLocation("alaska")

	fmt.Println("Location is " + location.Name)
	//

}

func initLocation(locationName string) *Location {
	configfile := "config/"+locationName+".toml"
	if _, err := os.Stat(configfile); err != nil {
		log.Fatal("No valid config file found for ",configfile)
	}

	var location Location
	if _, err := toml.DecodeFile(configfile, &location); err != nil {
		log.Fatal(err)
	}

	return &location
}

func vote() {
	ballot := new(Ballot)
	if randomdata.Boolean() {
		ballot.Choice = MOVE
	} else {
		ballot.Choice = STAY
	}

	ballot.PollingLocation = location //this will be set by the config file
	ballot.Time = time.Now()
}

// An example of how you would implement an enum in Go
type Party string
const (
	VIM Party = "vim"
	EMACS Party = "emacs"
	SPACES Party = "spaces"
	TABS Party = "tabs"
)

type Race string
const (
	BLUE Race = "blue"
	GREEN Race = "green"
	PURPLE  Race = "purple"
	ORANGE Race = "orange"
)

type BallotOption int
const (
	MOVE BallotOption = iota
	STAY
)


type Location struct {
	Ballot
	Another
	Name string
	Population int
	Age int
	HouseholdIncome int
	PartyPreference Party 		//the most popular party in the location
	RacialBreakdown map[Race] float32
}

type Another struct {
	name string
}

type Ballot struct {
	Choice BallotOption
	PollingLocation *Location
	Time time.Time
}
