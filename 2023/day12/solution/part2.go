package solution

import (
	"fmt"
	"strings"
)

const COPIES = 5

var record string
var cnts []int
var memo [][][]int
func Part2(data string) int {
	records, damagedCnts := parse(data)
	
	newRecords, newCnts := unfoldRecords(records), unforldCnts(damagedCnts)

	sum := 0
	for i, r := range newRecords {
		record, cnts = r, newCnts[i]
		fmt.Println(record, cnts)
		memo = initMemo(len(record), len(cnts), 1000)
		ways := findWays(0, 0, 0)
		fmt.Println(ways)
		sum += ways
	}
	return sum
}

func unfoldRecords(records []string) []string {
	var nrecords []string

	for _, r := range records {
		var copies []string
		for i := 0; i < COPIES; i++ {
			copies = append(copies, r)
		}
		nrecords = append(nrecords, strings.Join(copies, "?"))
	}
	
	return nrecords
}

func unforldCnts(cnts [][]int) [][]int {
	ncnts := make([][]int, len(cnts))
	
	for i, cs := range cnts {
		for k := 0; k < COPIES; k++ {
			ncnts[i] = append(ncnts[i], cs...)
		}	
	}
	
	return ncnts
}

func initMemo(x, y, z int) [][][]int {
	memo := make([][][]int, x)

	for i, _ := range memo {
		memo[i] = make([][]int, y+1)
		for j, _ := range memo[i] {
			memo[i][j] = make([]int, z)
			for k, _ := range memo[i][j] {
				memo[i][j][k] = -1
			}
		}
	}
	return memo
}

func findWays(recordPos, cntPos, damagedCnt int) int {
	if recordPos >= len(record) {
		if cntPos >= len(cnts) && damagedCnt == 0 {
			return 1
		}
		if cntPos == len(cnts)-1 && damagedCnt == cnts[cntPos] {
			return 1
		}
		return 0
	}

	ans := memo[recordPos][cntPos][damagedCnt]
	if ans != -1 {
		return ans
	}
	
	ans = 0
	for _, curr := range []byte{'.', '#'} {
		if curr != record[recordPos] && record[recordPos] != '?' {
			continue
		}
		if curr == '#' {
			ans += findWays(recordPos+1, cntPos, damagedCnt+1)
		} else if curr == '.' && cntPos < len(cnts) && damagedCnt == cnts[cntPos] {
			ans += findWays(recordPos+1, cntPos+1, 0)
		} else if curr == '.' && damagedCnt == 0 {
			ans += findWays(recordPos+1, cntPos, damagedCnt)
		}
	}
	
	memo[recordPos][cntPos][damagedCnt] = ans
	return ans 
}

func findWaysForDamaged(recordPos, cntPos, damagedCnt int) int {
	return findWays(recordPos+1, cntPos, damagedCnt+1)
}

func findWaysForUnDamaged(recordPos, cntPos, damagedCnt int) int {
	if damagedCnt >= cnts[cntPos] {
		damagedCnt = 0
		cntPos++
	}
	return findWays(recordPos+1, cntPos, damagedCnt)
}

