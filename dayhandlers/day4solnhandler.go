package dayhandlers

import (
	"fmt"
	"os"

	"github.com/mahldcat/adventlibgolang2024/day4"
	"github.com/mahldcat/adventlibgolang2024/inputfetcher"
)

func HandleDay4Solution(fetcher inputfetcher.DataFetcher) {
	wordSearch, err := fetcher.GetDataByDay(4)

	if err != nil {
		fmt.Println("Error on Data Fetch", err)
		os.Exit(-2)
	}

	sln := day4.SolveDay4Part1(wordSearch)
	fmt.Printf("Day 4.1 Solution: %d\n", sln)

	sln = day4.SolveDay4Part2(wordSearch)
	fmt.Printf("Day 4.2 Solution: %d\n", sln)
}
