package base_64

// convertToBits converts a character to a 6-bit binary representation
func Convert_to_bits(c byte) []int {
	bits := make([]int, 6)
	i := int(c)
	index := 5

	for i != 0 && index >= 0 {
		bits[index] = i % 2
		i /= 2
		index--
	}
	return bits
}

// bitsAccumulation converts a string to a sequence of bits
func BitsAccumulation(s string) []int {
	var bitstream []int
	for i := 0; i < len(s); i++ {
		bits := Convert_to_bits(s[i])
		bitstream = append(bitstream, bits...)
	}
	return bitstream
}

// convertBitsTo6Bytes groups bits into chunks of 6 bits
func ConvertBitsTo6Bytes(bits []int) [][]int {
	var byteStream [][]int
	for i := 0; i < len(bits); i += 6 {
		byteStream = append(byteStream, bits[i:i+6])
	}
	return byteStream
}
