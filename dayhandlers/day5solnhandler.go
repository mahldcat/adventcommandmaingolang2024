package dayhandlers

import (
	"fmt"
	"os"

	"github.com/mahldcat/adventlibgolang2024/day5"
	"github.com/mahldcat/adventlibgolang2024/inputfetcher"
)

func HandleDay5olution(fetcher inputfetcher.DataFetcher) {
	rawData, err := fetcher.GetDataByDay(5)

	if err != nil {
		fmt.Println("Error on Data Fetch", err)
		os.Exit(-2)
	}

	sln := day5.SolveDay5Part1(rawData)
	fmt.Printf("Day 5.1 Solution: %d\n", sln)

	sln = day5.SolveDay5Part2(rawData)
	fmt.Printf("Day 5.2 Solution: %d\n", sln)
}
