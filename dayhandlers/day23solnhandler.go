package dayhandlers

import (
	"fmt"
	"os"

	"github.com/mahldcat/adventlibgolang2024/day23"
	"github.com/mahldcat/adventlibgolang2024/inputfetcher"
)

func HandleDay23Solution(fetcher inputfetcher.DataFetcher) {
	rawData, err := fetcher.GetDataByDay(23)

	if err != nil {
		fmt.Println("Error on Data Fetch", err)
		os.Exit(-2)
	}

	sln := day23.SolveDay23Part1(rawData)
	fmt.Printf("Day 23.1 Solution: %d\n", sln)

	sln = day23.SolveDay23Part2(rawData)
	fmt.Printf("Day 23.2 Solution: %d\n", sln)
}
