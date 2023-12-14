package solution

func Part1(data string) int {
	records, damagedCnts := parse(data)

	sum := 0
	for i, r := range records {
		ways := findWaysToFixRecord([]byte(r), damagedCnts[i], 0)
		sum += ways
	}
	return sum
}

func findWaysToFixRecord(record []byte, cnts []int, recordPos int) int {
	if recordPos == len(record) && isRecordValid(record, cnts) {
		return 1
	}
	if recordPos == len(record) {
		return 0
	}
	if record[recordPos] != '?' {
		return findWaysToFixRecord(record, cnts, recordPos+1)	
	} 
	
	record[recordPos] = '.'
	ways := findWaysToFixRecord(record, cnts, recordPos+1)
	record[recordPos] = '#'
	ways += findWaysToFixRecord(record, cnts, recordPos+1)
	record[recordPos] = '?'	// backtrack
	return ways
}

func isRecordValid(record []byte, expDamagedCnts []int) bool {
	cnts := 0
	var actulDamagedCnts []int

	for _, ch := range record {
		if ch == '#' {
			cnts++
			continue
		}
		if cnts > 0 {
			actulDamagedCnts = append(actulDamagedCnts, cnts)
			cnts = 0
		}
	}
	if cnts > 0 {
		actulDamagedCnts = append(actulDamagedCnts, cnts)
	}
	return areEqual(expDamagedCnts, actulDamagedCnts)
}

func areEqual(nums1, nums2 []int) bool {
	if len(nums1) != len(nums2) {
		return false
	}
	for i, n1 := range nums1 {
		if nums2[i] != n1 {
			return false
		}
	}
	return true
}