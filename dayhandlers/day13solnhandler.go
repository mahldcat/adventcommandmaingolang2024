package dayhandlers

import (
	"fmt"
	"os"

	"github.com/mahldcat/adventlibgolang2024/day13"
	"github.com/mahldcat/adventlibgolang2024/inputfetcher"
)

func HandleDay13Solution(fetcher inputfetcher.DataFetcher) {
	rawData, err := fetcher.GetDataByDay(13)

	if err != nil {
		fmt.Println("Error on Data Fetch", err)
		os.Exit(-2)
	}

	sln := day13.SolveDay13Part1(rawData)
	fmt.Printf("Day 13.1 Solution: %d\n", sln)

	sln = day13.SolveDay13Part2(rawData)
	fmt.Printf("Day 13.2 Solution: %d\n", sln)
}
