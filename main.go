package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/mahldcat/adventcommandmaingolang2024/dayhandlers"
	"github.com/mahldcat/adventlibgolang2024/inputfetcher"
)

func getBearerToken() string {
	// os.Args[0] is the program name, and subsequent elements are the arguments
	if len(os.Args) < 3 {
		fmt.Println("pass the bearer token in as the second arg on the command line")
		os.Exit(-1)
	}
	return os.Args[2]
}

func getHandlerId() int {
	// os.Args[0] is the program name, and subsequent elements are the arguments
	if len(os.Args) < 2 {
		fmt.Println("pass the bearer token in as the second arg on the command line")
		os.Exit(-1)
	}

	hId, err := strconv.Atoi(os.Args[1])

	if err != nil {
		fmt.Printf("Error parsing '%s'\n", os.Args[1])
		os.Exit(-1)
	}

	return hId
}

func getDataFetcher() inputfetcher.DataFetcher {
	bearerToken := getBearerToken()

	return inputfetcher.NewDataFetcher(bearerToken)
}

var dailyHandlers = map[int]func(inputfetcher.DataFetcher){
	1: dayhandlers.HandleDay1Solution, //in day1solnhandler
	4: dayhandlers.HandleDay4Solution,
	5: dayhandlers.HandleDay5Solution,
	6: dayhandlers.HandleDay6Solution,
	7: dayhandlers.HandleDay7Solution,

	8:  dayhandlers.HandleDay8Solution,
	9:  dayhandlers.HandleDay9Solution,
	10: dayhandlers.HandleDay10Solution,
	11: dayhandlers.HandleDay11Solution,
	12: dayhandlers.HandleDay12Solution,
	13: dayhandlers.HandleDay13Solution,
	14: dayhandlers.HandleDay14Solution,
	15: dayhandlers.HandleDay15Solution,
	16: dayhandlers.HandleDay16Solution,
	17: dayhandlers.HandleDay17Solution,
	18: dayhandlers.HandleDay18Solution,
	19: dayhandlers.HandleDay19Solution,
	20: dayhandlers.HandleDay20Solution,
	21: dayhandlers.HandleDay21Solution,
	22: dayhandlers.HandleDay22Solution,
	23: dayhandlers.HandleDay23Solution,
	24: dayhandlers.HandleDay24Solution,
}

func main() {

	dataFetcher := getDataFetcher()

	handlerId := getHandlerId()

	handler, exists := dailyHandlers[handlerId]
	if !exists {
		fmt.Printf("Handler for Id %d not found", handlerId)
		os.Exit(-99)
	}

	handler(dataFetcher)

}
