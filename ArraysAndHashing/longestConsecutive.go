package arrays_and_hashing

func LongestConsecutive(nums []int) int {
	hashMap := make(map[int]bool)
	var long int

	for _, v := range nums {
		hashMap[v] = true
	}

	for _, v := range nums {
		if hashMap[v-1] {
			continue
		}

		seq := 1
		tmp := v + 1

		for hashMap[tmp] {
			seq++
			tmp++
		}

		if seq > long {
			long = seq
		}
	}
	return long
}
