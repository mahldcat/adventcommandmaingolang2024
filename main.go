package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/mahldcat/adventlibgolang2024/day1"
	"github.com/mahldcat/adventlibgolang2024/inputfetcher"
)

func getBearerToken() (string, error) {
	// os.Args[0] is the program name, and subsequent elements are the arguments
	if len(os.Args) < 2 {
		return "", errors.New("Need to pass the bearer token in as arg 1")
	}
	return os.Args[1], nil
}

func main() {

	bearerToken, err := getBearerToken()
	if err != nil {
		fmt.Println("Error getting bearer token", err)
		os.Exit(-1)
	}

	fetcher := inputfetcher.NewDataFetcher(bearerToken)

	left, right, err := fetcher.FetchDay1Data()

	if err != nil {
		fmt.Println("Error on Data Fetch", err)
		os.Exit(-2)
	}

	sln, err := day1.SolveDay1Part1(left, right)
	if err != nil {
		fmt.Println("Error on Day1.1 solution", err)
		os.Exit(-3)
	}

	fmt.Printf("Day 1.1 Solution: %d\n", sln)

	sln, err = day1.SolveDay1Part2(left, right)
	if err != nil {
		fmt.Println("Error on Day1.2 solution", err)
		os.Exit(-4)
	}

	fmt.Printf("Day 1.2 Solution: %d\n", sln)

}
