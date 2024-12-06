package dayhandlers

import (
	"fmt"
	"os"

	"github.com/mahldcat/adventlibgolang2024/day8"
	"github.com/mahldcat/adventlibgolang2024/inputfetcher"
)

func HandleDay8Solution(fetcher inputfetcher.DataFetcher) {
	rawData, err := fetcher.GetDataByDay(8)

	if err != nil {
		fmt.Println("Error on Data Fetch", err)
		os.Exit(-2)
	}

	sln := day8.SolveDay8Part1(rawData)
	fmt.Printf("Day 8.1 Solution: %d\n", sln)

	sln = day8.SolveDay8Part2(rawData)
	fmt.Printf("Day 8.2 Solution: %d\n", sln)
}
