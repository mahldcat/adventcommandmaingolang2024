package dayhandlers

import (
	"fmt"
	"os"

	"github.com/mahldcat/adventlibgolang2024/day4"
	"github.com/mahldcat/adventlibgolang2024/inputfetcher"
)

func HandleDay6Solution(fetcher inputfetcher.DataFetcher) {
	rawData, err := fetcher.GetDataByDay(6)

	if err != nil {
		fmt.Println("Error on Data Fetch", err)
		os.Exit(-2)
	}

	sln := day4.SolveDay4Part1(rawData)
	fmt.Printf("Day 6.1 Solution: %d\n", sln)

	sln = day4.SolveDay4Part2(rawData)
	fmt.Printf("Day 6.2 Solution: %d\n", sln)
}
