package dayhandlers

import (
	"fmt"
	"os"

	"github.com/mahldcat/adventlibgolang2024/day21"
	"github.com/mahldcat/adventlibgolang2024/inputfetcher"
)

func HandleDay21Solution(fetcher inputfetcher.DataFetcher) {
	rawData, err := fetcher.GetDataByDay(21)

	if err != nil {
		fmt.Println("Error on Data Fetch", err)
		os.Exit(-2)
	}

	sln := day21.SolveDay21Part1(rawData)
	fmt.Printf("Day 21.1 Solution: %d\n", sln)

	sln = day21.SolveDay21Part2(rawData)
	fmt.Printf("Day 21.2 Solution: %d\n", sln)
}
