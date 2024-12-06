package dayhandlers

import (
	"fmt"
	"os"

	"github.com/mahldcat/adventlibgolang2024/day22"
	"github.com/mahldcat/adventlibgolang2024/inputfetcher"
)

func HandleDay22Solution(fetcher inputfetcher.DataFetcher) {
	rawData, err := fetcher.GetDataByDay(22)

	if err != nil {
		fmt.Println("Error on Data Fetch", err)
		os.Exit(-2)
	}

	sln := day22.SolveDay22Part1(rawData)
	fmt.Printf("Day 22.1 Solution: %d\n", sln)

	sln = day22.SolveDay22Part2(rawData)
	fmt.Printf("Day 22.2 Solution: %d\n", sln)
}
