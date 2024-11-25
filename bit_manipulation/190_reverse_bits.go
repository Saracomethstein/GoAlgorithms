package bit_manipulation

// link to the task on leetcode
// https://leetcode.com/problems/reverse-bits/description/

func reverseBits(num uint32) uint32 {
	var reversed uint32
	for i := 0; i < 32; i++ {
		reversed <<= 1
		reversed |= num & 1
		num >>= 1
	}
	return reversed
}
