package main

import (
	"fmt"
	"os"

	"github.com/mahldcat/adventlibgolang2024/day1"
	"github.com/mahldcat/adventlibgolang2024/inputfetcher"
)

func HandleDay1Solution(fetcher inputfetcher.DataFetcher) {
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