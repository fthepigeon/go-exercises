package main

func reverseBits(num uint32) uint32 {
	var reversed uint32

	for i := 0; i < 32; i++ {
		reversed *= 2

		if num > 0 {
			reversed += num & uint32(1)
			num >>= 1
		}
	}
	return reversed
}
