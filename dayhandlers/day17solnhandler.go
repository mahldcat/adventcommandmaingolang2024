package dayhandlers

import (
	"fmt"
	"os"

	"github.com/mahldcat/adventlibgolang2024/day17"
	"github.com/mahldcat/adventlibgolang2024/inputfetcher"
)

func HandleDay17Solution(fetcher inputfetcher.DataFetcher) {
	rawData, err := fetcher.GetDataByDay(17)

	if err != nil {
		fmt.Println("Error on Data Fetch", err)
		os.Exit(-2)
	}

	sln := day17.SolveDay17Part1(rawData)
	fmt.Printf("Day 17.1 Solution: %d\n", sln)

	sln = day17.SolveDay17Part2(rawData)
	fmt.Printf("Day 17.2 Solution: %d\n", sln)
}
