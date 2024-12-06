package dayhandlers

import (
	"fmt"
	"os"

	"github.com/mahldcat/adventlibgolang2024/day15"
	"github.com/mahldcat/adventlibgolang2024/inputfetcher"
)

func HandleDay15Solution(fetcher inputfetcher.DataFetcher) {
	rawData, err := fetcher.GetDataByDay(15)

	if err != nil {
		fmt.Println("Error on Data Fetch", err)
		os.Exit(-2)
	}

	sln := day15.SolveDay15Part1(rawData)
	fmt.Printf("Day 15.1 Solution: %d\n", sln)

	sln = day15.SolveDay15Part2(rawData)
	fmt.Printf("Day 15.2 Solution: %d\n", sln)
}
