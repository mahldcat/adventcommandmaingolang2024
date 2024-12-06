package dayhandlers

import (
	"fmt"
	"os"

	"github.com/mahldcat/adventlibgolang2024/day20"
	"github.com/mahldcat/adventlibgolang2024/inputfetcher"
)

func HandleDay20Solution(fetcher inputfetcher.DataFetcher) {
	rawData, err := fetcher.GetDataByDay(20)

	if err != nil {
		fmt.Println("Error on Data Fetch", err)
		os.Exit(-2)
	}

	sln := day20.SolveDay20Part1(rawData)
	fmt.Printf("Day 20.1 Solution: %d\n", sln)

	sln = day20.SolveDay20Part2(rawData)
	fmt.Printf("Day 20.2 Solution: %d\n", sln)
}
