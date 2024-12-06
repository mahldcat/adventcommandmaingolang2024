package dayhandlers

import (
	"fmt"
	"os"

	"github.com/mahldcat/adventlibgolang2024/day14"
	"github.com/mahldcat/adventlibgolang2024/inputfetcher"
)

func HandleDay14Solution(fetcher inputfetcher.DataFetcher) {
	rawData, err := fetcher.GetDataByDay(14)

	if err != nil {
		fmt.Println("Error on Data Fetch", err)
		os.Exit(-2)
	}

	sln := day14.SolveDay14Part1(rawData)
	fmt.Printf("Day 14.1 Solution: %d\n", sln)

	sln = day14.SolveDay14Part2(rawData)
	fmt.Printf("Day 14.2 Solution: %d\n", sln)
}
