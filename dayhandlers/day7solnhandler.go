package dayhandlers

import (
	"fmt"
	"os"

	"github.com/mahldcat/adventlibgolang2024/day7"
	"github.com/mahldcat/adventlibgolang2024/inputfetcher"
)

func HandleDay7Solution(fetcher inputfetcher.DataFetcher) {
	rawData, err := fetcher.GetDataByDay(7)

	if err != nil {
		fmt.Println("Error on Data Fetch", err)
		os.Exit(-2)
	}

	sln := day7.SolveDay7Part1(rawData)
	fmt.Printf("Day 7.1 Solution: %d\n", sln)

	sln = day7.SolveDay7Part2(rawData)
	fmt.Printf("Day 7.2 Solution: %d\n", sln)
}
