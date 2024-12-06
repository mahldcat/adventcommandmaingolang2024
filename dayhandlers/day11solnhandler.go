package dayhandlers

import (
	"fmt"
	"os"

	"github.com/mahldcat/adventlibgolang2024/day11"
	"github.com/mahldcat/adventlibgolang2024/inputfetcher"
)

func HandleDay11Solution(fetcher inputfetcher.DataFetcher) {
	rawData, err := fetcher.GetDataByDay(11)

	if err != nil {
		fmt.Println("Error on Data Fetch", err)
		os.Exit(-2)
	}

	sln := day11.SolveDay11Part1(rawData)
	fmt.Printf("Day 11.1 Solution: %d\n", sln)

	sln = day11.SolveDay11Part2(rawData)
	fmt.Printf("Day 11.2 Solution: %d\n", sln)
}
