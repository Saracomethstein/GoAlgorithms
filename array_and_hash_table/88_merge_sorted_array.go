package arrays_and_hashing

import (
	"sort"
)

// link to the task on leetcode
// https://leetcode.com/problems/merge-sorted-array/description/

func merge(nums1 []int, m int, nums2 []int, n int) {
	j := 0
	for i := m; i < m+n; i++ {
		nums1[i] = nums2[j]
		j++
	}
	sort.Ints(nums1)
}
