package arrays_and_hashing

// link to the task on leetcode 
// https://leetcode.com/problems/two-sum/description/

func TwoSum(nums []int, target int) []int {
	var hash = make(map[int]int)
	for i, num := range nums {
		if j, ok := hash[target-num]; ok {
			return []int{i, j}
		}
		hash[num] = i
	}
	return nil	
}
