package dayhandlers

import (
	"fmt"
	"os"

	"github.com/mahldcat/adventlibgolang2024/day1"
	"github.com/mahldcat/adventlibgolang2024/inputfetcher"
)

func HandleDay1Solution(fetcher inputfetcher.DataFetcher) {
	rawData, err := fetcher.GetDataByDay(1)

	if err != nil {
		fmt.Println("Error on Data Fetch", err)
		os.Exit(-2)
	}

	sln, err := day1.SolveDay1Part1(rawData)
	if err != nil {
		fmt.Println("Error on Day1.1 solution", err)
		os.Exit(-3)
	}

	fmt.Printf("Day 1.1 Solution: %d\n", sln)

	sln, err = day1.SolveDay1Part2(rawData)
	if err != nil {
		fmt.Println("Error on Day1.2 solution", err)
		os.Exit(-4)
	}

	fmt.Printf("Day 1.2 Solution: %d\n", sln)
}
