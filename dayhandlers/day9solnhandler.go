package dayhandlers

import (
	"fmt"
	"os"

	"github.com/mahldcat/adventlibgolang2024/day9"
	"github.com/mahldcat/adventlibgolang2024/inputfetcher"
)

func HandleDay9Solution(fetcher inputfetcher.DataFetcher) {
	rawData, err := fetcher.GetDataByDay(9)

	if err != nil {
		fmt.Println("Error on Data Fetch", err)
		os.Exit(-2)
	}

	sln := day9.SolveDay9Part1(rawData)
	fmt.Printf("Day 9.1 Solution: %d\n", sln)

	sln = day9.SolveDay9Part2(rawData)
	fmt.Printf("Day 9.2 Solution: %d\n", sln)
}
