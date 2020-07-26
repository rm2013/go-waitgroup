package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type AdditionUnit struct {
	addends []int
	sum     int
}

func main() {

	for i := 0; i < 10; i++ {
		runAdditionUnits(i)
	}

}

func runAdditionUnits(numOneK int) {
	var waitGroup sync.WaitGroup

	start := time.Now()
	additionUnits := buildAdditionUnits(numOneK)
	//fmt.Println("Total AddUnits: ", len(additionUnits))
	waitGroup.Add(len(additionUnits))
	end := time.Now()
	elapsed := end.Sub(start)
	fmt.Println("Filled ", len(additionUnits), " in ", elapsed)
	start = time.Now()
	for index := range additionUnits {
		go addOneSetWithFixedDelay(additionUnits, index, &waitGroup, 3)
		//go addOneSetWithRandomDelay(additionUnits, index, &waitGroup, 6)
	}
	waitGroup.Wait()
	end = time.Now()
	elapsed = end.Sub(start)
	fmt.Println("computing time:", elapsed)
	/*for _, additionUnit := range additionUnits {
		fmt.Println(additionUnit)
	}*/
}

func build1KAdditionUnits() []AdditionUnit {
	additionUnits := []AdditionUnit{}
	for i := 0; i < 1000; i++ {
		additionUnit := AdditionUnit{}
		additionUnit.addends = []int{i, i + 1, i + 2}
		additionUnits = append(additionUnits, additionUnit)
	}
	return additionUnits
}

func buildAdditionUnits(numberOfKUnits int) []AdditionUnit {
	additionUnits := []AdditionUnit{}

	oneK := build1KAdditionUnits()
	additionUnits = oneK
	for i := 0; i < numberOfKUnits; i++ {
		//length := len(additionUnits)
		//fmt.Println("Length of array: ", length, " round: ", i)
		newSlice := additionUnits
		additionUnits = append(additionUnits, newSlice...)
	}

	return additionUnits
}

func addOneSetWithFixedDelay(additionUnits []AdditionUnit, index int, waitGroup *sync.WaitGroup, delayedSecs time.Duration) {
	time.Sleep(delayedSecs * time.Second)
	for _, i := range additionUnits[index].addends {
		additionUnits[index].sum += i
	}
	waitGroup.Done()
}

func addOneSetWithRandomDelay(additionUnits []AdditionUnit, index int, waitGroup *sync.WaitGroup, maxDelayedSecs int) {
	randomDelay := rand.Intn(maxDelayedSecs)
	time.Sleep(time.Duration(randomDelay) * time.Second)
	for _, i := range additionUnits[index].addends {
		additionUnits[index].sum += i
	}
	waitGroup.Done()
}
