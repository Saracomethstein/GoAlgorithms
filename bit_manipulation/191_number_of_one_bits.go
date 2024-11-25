package bit_manipulation

// link to the task on leetcode
// https://leetcode.com/problems/number-of-1-bits/description/

func hammingWeight(n int) int {
	weight := 0
	for i := 1; i <= 32; i++ {
		if n&1 == 1 {
			weight += 1
		}
		n >>= 1
	}
	return weight
}
