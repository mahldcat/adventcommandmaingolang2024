package dayhandlers

import (
	"fmt"
	"os"

	"github.com/mahldcat/adventlibgolang2024/day24"
	"github.com/mahldcat/adventlibgolang2024/inputfetcher"
)

func HandleDay24Solution(fetcher inputfetcher.DataFetcher) {
	rawData, err := fetcher.GetDataByDay(24)

	if err != nil {
		fmt.Println("Error on Data Fetch", err)
		os.Exit(-2)
	}

	sln := day24.SolveDay24Part1(rawData)
	fmt.Printf("Day 24.1 Solution: %d\n", sln)

	sln = day24.SolveDay24Part2(rawData)
	fmt.Printf("Day 24.2 Solution: %d\n", sln)
}
