package dayhandlers

import (
	"fmt"
	"os"

	"github.com/mahldcat/adventlibgolang2024/day10"
	"github.com/mahldcat/adventlibgolang2024/inputfetcher"
)

func HandleDay10Solution(fetcher inputfetcher.DataFetcher) {
	rawData, err := fetcher.GetDataByDay(10)

	if err != nil {
		fmt.Println("Error on Data Fetch", err)
		os.Exit(-2)
	}

	sln := day10.SolveDay10Part1(rawData)
	fmt.Printf("Day 10.1 Solution: %d\n", sln)

	sln = day10.SolveDay10Part2(rawData)
	fmt.Printf("Day 10.2 Solution: %d\n", sln)
}
