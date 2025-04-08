package twoSum

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	var results []int

	for index := range len(nums) {

		requiredNumber := target - nums[index]
		if val, ok := m[requiredNumber]; ok {
			// match found we can break loop here
			results = append(results, val, index)
			break
		}

		// no match currently found - check for duplicate in map
		if _, ok := m[nums[index]]; ok {
			// key already exists - skip
			continue
		}

		m[nums[index]] = index

	}

	return results
}
