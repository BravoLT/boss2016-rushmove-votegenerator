package main

import (
	"fmt"
	"time"
	rd "github.com/Pallinder/go-randomdata"
	"strconv"
	"errors"
	"os"
	"encoding/json"
	"sync"
	"net/http"
	"bytes"
)

func main() {
	Test()

	var wg sync.WaitGroup

	wg.Add(50)

	fmt.Println("Opening Eastern Timezone polls at ", time.Now())
	//Eastern
	go openPoll("Connecticut", 3592053, 15, &wg)
	go openPoll("Delaware", 917060, 25, &wg)
	go openPoll("Florida", 19361792, -5, &wg)
	go openPoll("Georgia", 9907756, -5, &wg)
	go openPoll("Indiana", 6542411, -15, &wg)
	go openPoll("Maine", 1328535, 10, &wg)
	go openPoll("Maryland", 5887776, 45, &wg)
	go openPoll("Massachusetts", 6657291, 35, &wg)
	go openPoll("Michigan", 9889024, 5, &wg)
	go openPoll("NewHampshire", 1321069, 5, &wg)
	go openPoll("NewJersey", 8874374, 25, &wg)
	go openPoll("NewYork", 19594330, 30, &wg)
	go openPoll("NorthCarolina", 9750405, 15, &wg)
	go openPoll("Ohio", 11560380, -5, &wg)
	go openPoll("Pennsylvania", 12758729, 10, &wg)
	go openPoll("RhodeIsland", 1053252, 25, &wg)
	go openPoll("SouthCarolina", 4727273, 0, &wg)
	go openPoll("Vermont", 626358, 5, &wg)
	go openPoll("Virginia", 8185131, 5, &wg)
	go openPoll("WestVirginia", 1853881, -25, &wg)

	time.Sleep(time.Second * 120) //sleep, seconds to minutes

	fmt.Println("Opening Central Timezone polls at ", time.Now())
	//Central
	go openPoll("Alabama", 4817678, -10, &wg)
	go openPoll("Arkansas", 2947036, -5, &wg)
	go openPoll("Illinois", 12868747, 5, &wg)
	go openPoll("Iowa", 3078116, -15, &wg)
	go openPoll("Kansas", 2882946, -25, &wg)
	go openPoll("Kentucky", 4383272, -15, &wg)
	go openPoll("Louisiana", 4601049, 0, &wg)
	go openPoll("Minnesota", 5383661, -10, &wg)
	go openPoll("Mississippi", 2984345, -5, &wg)
	go openPoll("Missouri", 6028076, -15, &wg)
	go openPoll("Nebraska", 1855617, -10, &wg)
	go openPoll("NorthDakota", 704925, -10, &wg)
	go openPoll("Oklahoma", 3818851, -10, &wg)
	go openPoll("SouthDakota", 834708, -45, &wg)
	go openPoll("Tennessee", 6451365, -20, &wg)
	go openPoll("Texas", 26092033, -25, &wg)
	go openPoll("Wisconsin", 5724692, 20, &wg)

	time.Sleep(time.Second * 120) //sleep, seconds to minutes

	fmt.Println("Opening Mountain Timezone polls at ", time.Now())
	//Mountain
	go openPoll("Arizona", 6561516, -5, &wg)
	go openPoll("Colorado", 5197580, 5, &wg)
	go openPoll("Idaho", 1599464, -10, &wg)
	go openPoll("Montana", 1006370, -10, &wg)
	go openPoll("NewMexico", 2080085, 0, &wg)
	go openPoll("Utah", 2858111, -35, &wg)

	time.Sleep(time.Second * 120) //sleep, seconds to minutes

	fmt.Println("Opening Pacific Timezone polls at ", time.Now())
	//Pacific
	go openPoll("California", 38066920, 15, &wg)
	go openPoll("Nevada", 2761584, 0, &wg)
	go openPoll("Oregon", 3900343, -15, &wg)
	go openPoll("Washington", 6899123, 15, &wg)
	go openPoll("Wyoming", 575251, -30, &wg)

	time.Sleep(time.Second * 120) //sleep, seconds to minutes

	fmt.Println("Opening Alaska polls at ", time.Now())
	go openPoll("Alaska", 728300, 0, &wg)

	time.Sleep(time.Second * 120) //sleep, seconds to minutes
	fmt.Println("Opening Hawaii polls at ", time.Now())
	go openPoll("Hawaii", 1392704, 0, &wg)

	wg.Wait()
}

//run this as a goroutine
func openPoll(location string, population int, weight int, wg *sync.WaitGroup) {
	defer wg.Done()

	pool := generateVotingPool(weight)

	// set participation at 65%
	voters := int(float32(population) * .65)
	var move, stay int

	for i := 0; i < voters; i++ {
		randomIdx := rd.Number(0, 99)
		choice := pool[randomIdx]

		if choice {
			move++
		} else {
			stay++
		}

		ballot := Ballot{ Choice: choice, Location: location, Time: time.Now().Unix() }

		sendToFlume(ballot)
	}

	movePct := (float32(move) / float32(voters)) * float32(100)
	stayPct := (float32(stay) / float32(voters)) * float32(100)
	fmt.Println("Time: ", time.Now(),"[", location, "]\n\t",movePct, "% to MOVE\n\t",stayPct, "% to STAY")

}

func sendToFlume(ballot Ballot) error {
	url := "http://45.31.163.169:4444/"

	ballotJson, _ := json.Marshal(ballot)
	jsonEnvelope := JsonBody{Body: string(ballotJson)}
	newPayload := [1]JsonBody {jsonEnvelope}

	json, err := json.Marshal(newPayload)
	if err != nil {
		fmt.Println("ERROR generating a ballot")
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(json))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	if rsp, err := client.Do(req); err != nil {
		return err
	} else {
		defer rsp.Body.Close()
		if rsp.StatusCode != 200 {
			fmt.Println("Failure on request ",rsp.StatusCode, string(json))
		}
	}

	return nil
}


type PollTime struct {
	start time.Time
}

func (pt *PollTime) mark() {
	current := (time.Now().Sub(pt.start)).Nanoseconds() + 8
	fmt.Println("Current time: ", current)
}

func generateVotingPool(weight int) ChoicePool {
	choicePool := ChoicePool{}
	setTo := weight < 0

	if !setTo {
		//if we are setting to false, flip all the intial values to true
		for i := range choicePool {
			choicePool[i] = true
		}

		// need to reverse sign since we are reversing initial values
		weight = weight * -1
	}

	distribution := 50 + weight
	if distribution >= 100 || distribution <= 0 {
		return choicePool
	}

	for i := 0; i < distribution; i++ {
		idx := rd.Number(0, 99)
		if choicePool[idx] != setTo {
			choicePool[idx] = setTo
		} else {
			//decrement and try again if the value has already been set
			i--
		}
	}

	return choicePool;
}

type ChoicePool [100]bool


type JsonBody struct {
	Body string `json:"body"`
}

type Ballot struct {
	Choice   bool      `json:"choice"`
	Location string    `json:"location"`
	Time     int64     `json:"time"`
}

// Unit Tests

func Test() {
	if err := testWeights(0, 50, true); err != nil {
		reportFailure(err)
	} else {
		fmt.Println("testWeights 0,50 .... SUCCESS")
	}

	if err := testWeights(-10, 40, true); err != nil {
		reportFailure(err)
	} else {
		fmt.Println("testWeights -10,40 .... SUCCESS")
	}

	if err := testWeights(10, 60, true); err != nil {
		reportFailure(err)
	} else {
		fmt.Println("testWeights 10,60 .... SUCCESS")
	}

	if err := testWeights(-50, 0, true); err != nil {
		reportFailure(err)
	} else {
		fmt.Println("testWeights -50,0 .... SUCCESS")
	}

	if err := testWeights(50, 100, true); err != nil {
		reportFailure(err)
	} else {
		fmt.Println("testWeights 50,100 .... SUCCESS")
	}

}

func reportFailure(err error) {
	fmt.Println("Unit Tests failed, can not continue execution")
	fmt.Println(err)

	os.Exit(1)
}

func testWeights(weight int, expectedCount int, expectedValue bool) error {
	pool := generateVotingPool(weight)
	if l := len(pool); l != 100 {
		return errors.New("Voting pool was the wrong size, should be 100, but was " + strconv.Itoa(len(pool)))
	}

	actualCount := count(pool[:], expectedValue)
	if actualCount != expectedCount {
		return errors.New("Wrong number of " + strconv.FormatBool(expectedValue) + " values, should be " + strconv.Itoa(expectedCount) + " but was " + strconv.Itoa(actualCount))
	}

	return nil;
}

func count(vs []bool, value bool) int {
	count := 0
	for _, v := range vs {
		if v == value {
			count++
		}
	}

	return count
}
