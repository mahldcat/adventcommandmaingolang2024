package dayhandlers

import (
	"fmt"
	"os"

	"github.com/mahldcat/adventlibgolang2024/day16"
	"github.com/mahldcat/adventlibgolang2024/inputfetcher"
)

func HandleDay16Solution(fetcher inputfetcher.DataFetcher) {
	rawData, err := fetcher.GetDataByDay(16)

	if err != nil {
		fmt.Println("Error on Data Fetch", err)
		os.Exit(-2)
	}

	sln := day16.SolveDay16Part1(rawData)
	fmt.Printf("Day 16.1 Solution: %d\n", sln)

	sln = day16.SolveDay16Part2(rawData)
	fmt.Printf("Day 16.2 Solution: %d\n", sln)
}
