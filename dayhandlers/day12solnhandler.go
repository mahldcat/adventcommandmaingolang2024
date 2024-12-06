package dayhandlers

import (
	"fmt"
	"os"

	"github.com/mahldcat/adventlibgolang2024/day12"
	"github.com/mahldcat/adventlibgolang2024/inputfetcher"
)

func HandleDay12Solution(fetcher inputfetcher.DataFetcher) {
	rawData, err := fetcher.GetDataByDay(12)

	if err != nil {
		fmt.Println("Error on Data Fetch", err)
		os.Exit(-2)
	}

	sln := day12.SolveDay12Part1(rawData)
	fmt.Printf("Day 12.1 Solution: %d\n", sln)

	sln = day12.SolveDay12Part2(rawData)
	fmt.Printf("Day 12.2 Solution: %d\n", sln)
}
