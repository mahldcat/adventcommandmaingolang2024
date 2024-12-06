package dayhandlers

import (
	"fmt"
	"os"

	"github.com/mahldcat/adventlibgolang2024/day6"
	"github.com/mahldcat/adventlibgolang2024/inputfetcher"
)

func HandleDay6Solution(fetcher inputfetcher.DataFetcher) {
	rawData, err := fetcher.GetDataByDay(6)

	if err != nil {
		fmt.Println("Error on Data Fetch", err)
		os.Exit(-2)
	}

	sln := day6.SolveDay6Part1(rawData)
	fmt.Printf("Day 6.1 Solution: %d\n", sln)

	sln = day6.SolveDay6Part2(rawData)
	fmt.Printf("Day 6.2 Solution: %d\n", sln)
}
