package main

func hammingWeight(num uint32) int {
	var ones uint32

	for num > 0 {
		ones += num & uint32(1)
		num >>= 1
	}
	return int(ones)
}
