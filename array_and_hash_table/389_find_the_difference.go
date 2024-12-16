package arrays_and_hashing

func FindTheDifference(s string, t string) byte {
	var sumS, sumT byte

	for _, v := range t {
		sumT += byte(v)
	}

	for _, v := range s {
		sumS += byte(v)
	}

	return sumT - sumS
}
