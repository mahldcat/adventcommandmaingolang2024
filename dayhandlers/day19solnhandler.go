package dayhandlers

import (
	"fmt"
	"os"

	"github.com/mahldcat/adventlibgolang2024/day19"
	"github.com/mahldcat/adventlibgolang2024/inputfetcher"
)

func HandleDay19Solution(fetcher inputfetcher.DataFetcher) {
	rawData, err := fetcher.GetDataByDay(19)

	if err != nil {
		fmt.Println("Error on Data Fetch", err)
		os.Exit(-2)
	}

	sln := day19.SolveDay19Part1(rawData)
	fmt.Printf("Day 19.1 Solution: %d\n", sln)

	sln = day19.SolveDay19Part2(rawData)
	fmt.Printf("Day 19.2 Solution: %d\n", sln)
}
