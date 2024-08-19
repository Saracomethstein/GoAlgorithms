func containsDuplicate(nums []int) bool {
	var noneDup = make(map[int]bool)

	for _, v := range nums {
		if _, ok := noneDup[v]; ok {
			return true
		}
		noneDup[v] = true
	}

	return false
}
