package dayhandlers

import (
	"fmt"
	"os"

	"github.com/mahldcat/adventlibgolang2024/day18"
	"github.com/mahldcat/adventlibgolang2024/inputfetcher"
)

func HandleDay18Solution(fetcher inputfetcher.DataFetcher) {
	rawData, err := fetcher.GetDataByDay(18)

	if err != nil {
		fmt.Println("Error on Data Fetch", err)
		os.Exit(-2)
	}

	sln := day18.SolveDay18Part1(rawData)
	fmt.Printf("Day 18.1 Solution: %d\n", sln)

	sln = day18.SolveDay18Part2(rawData)
	fmt.Printf("Day 18.2 Solution: %d\n", sln)
}
