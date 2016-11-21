package main

import (
	"fmt"
	"time"
	"github.com/Pallinder/go-randomdata"
	"strconv"
)


func main() {

	locations := initializeLocations()

	fmt.Println("num locations  " + strconv.Itoa(len(locations)))
}

func initializeLocations() map[string] Location {
	locations :=  make(map[string] Location)
	locations["AL"] = generateLocation(4817678,-10)
	locations["AK"] = generateLocation(728300,0)
	locations["AZ"] = generateLocation(6561516,-5)
	locations["AR"] = generateLocation(2947036,-5)
	locations["CA"] = generateLocation(38066920,15)
	locations["CO"] = generateLocation(5197580,5)
	locations["CT"] = generateLocation(3592053,15)
	locations["DE"] = generateLocation(917060,15)
	locations["FL"] = generateLocation(19361792,-5)
	locations["GA"] = generateLocation(9907756,-5)
	locations["HI"] = generateLocation(1392704,0)
	locations["ID"] = generateLocation(1599464,-10)
	locations["IL"] = generateLocation(12868747,5)
	locations["IN"] = generateLocation(6542411,-15)
	locations["IA"] = generateLocation(4383272,-15)
	locations["KS"] = generateLocation(2882946,-25)
	locations["KY"] = generateLocation(4383272,-15)
	locations["LA"] = generateLocation(4601049,0)
	locations["ME"] = generateLocation(1328535,10)
	locations["MD"] = generateLocation(5887776,45)
	locations["MA"] = generateLocation(6657291,35)
	locations["MI"] = generateLocation(9889024,5)
	locations["MN"] = generateLocation(5383661,-10)
	locations["MS"] = generateLocation(2984345,-5)
	locations["MO"] = generateLocation(6028076,-15)
	locations["MT"] = generateLocation(1006370,-10)
	locations["NE"] = generateLocation(1855617,-10)
	locations["NV"] = generateLocation(2761584,0)
	locations["NH"] = generateLocation(1321069,5)
	locations["NJ"] = generateLocation(8874374,25)
	locations["NM"] = generateLocation(2080085,0)
	locations["NY"] = generateLocation(19594330,30)
	locations["NC"] = generateLocation(9750405,15)
	locations["ND"] = generateLocation(704925,-10)
	locations["OH"] = generateLocation(11560380,-5)
	locations["OK"] = generateLocation(3818851,-10)
	locations["OR"] = generateLocation(3900343,-15)
	locations["PA"] = generateLocation(12758729,10)
	locations["RI"] = generateLocation(1053252,25)
	locations["SC"] = generateLocation(4727273,0)
	locations["SD"] = generateLocation(834708,-45)
	locations["TN"] = generateLocation(6451365,-20)
	locations["TX"] = generateLocation(26092033,-25)
	locations["UT"] = generateLocation(2858111,-35)
	locations["VT"] = generateLocation(626358,5)
	locations["VA"] = generateLocation(8185131, 5)
	locations["WA"] = generateLocation(6899123, 15)
	locations["WV"] = generateLocation(1853881, -25)
	locations["WI"] = generateLocation(5724692, 20)
	locations["WY"] = generateLocation(575251, -30)


	return locations
}

type Location struct {
	Population int
	VotesFor int
	VotesAgainst int
	votePool ChoicePool
}

func generateLocation(population int, weight int) Location {
	return Location {
		Population: population,
		votePool: generateVotingPool(weight),
	}
}

func generateVotingPool(weight int) ChoicePool {

	choicePool := ChoicePool{}

	distribution := 50 + weight
	for i := 0; i < distribution; i++ {
		idx := randomdata.Number(0,99)
		if choicePool[idx] != true {
			choicePool[idx] = true
		}
	}

	return choicePool;
}

type ChoicePool [100]bool


type Ballot struct {
	Choice bool
	Location string
	Time time.Time
}
