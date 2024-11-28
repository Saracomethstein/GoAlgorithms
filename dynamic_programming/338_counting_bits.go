package dynamic_programming

// link to the task on leetcode
// https://leetcode.com/problems/counting-bits/description/

func countBits(n int) []int {
	ans := make([]int, n+1)

	for i := 0; i <= n; i++ {
		ans[i] = ans[i/2] + i%2
	}

	return ans
}
