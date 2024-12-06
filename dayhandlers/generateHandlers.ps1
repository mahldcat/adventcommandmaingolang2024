# Generates a handler template for each day

$template=
@'
package dayhandlers

import (
	"fmt"
	"os"

	"github.com/mahldcat/adventlibgolang2024/day##DAY##"
	"github.com/mahldcat/adventlibgolang2024/inputfetcher"
)

func HandleDay##DAY##Solution(fetcher inputfetcher.DataFetcher) {
	rawData, err := fetcher.GetDataByDay(##DAY##)

	if err != nil {
		fmt.Println("Error on Data Fetch", err)
		os.Exit(-2)
	}

	sln := day##DAY##.SolveDay##DAY##Part1(rawData)
	fmt.Printf("Day ##DAY##.1 Solution: %d\n", sln)

	sln = day##DAY##.SolveDay##DAY##Part2(rawData)
	fmt.Printf("Day ##DAY##.2 Solution: %d\n", sln)
}
'@

foreach($d in 6..24){
    echo $template.Replace("##DAY##",$d) > "day$($d)solnhandler.go"
}