package main

import (
	"fmt"

	"github.com/chazsmi/mapsexcerise/people"
)

func main() {
	channel := make(chan []string)
	go people.CharliesTeam(channel)
	go people.PeopleOnFloorFive(channel)
	go people.PeopleWorkingInProduct(channel)

	for i := 0; i < 3; i++ {
		fmt.Printf("Result %s\n", <-channel)
	}
}
