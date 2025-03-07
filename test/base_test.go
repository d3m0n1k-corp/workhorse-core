package base_64_test

import (
	"fmt"
	"testing"
	"workhorse-core/base_64"
)

func TestBase64(t *testing.T) {
	testStr := "Man"
	bitrepresentation := base_64.BitsAccumulation(testStr)

	groupedBits := base_64.ConvertBitsTo6Bytes(bitrepresentation)

	fmt.Println("Grouped into 6 bit chunks:")
	for _, chunks := range groupedBits {
		for _, bit := range chunks {
			fmt.Printf("%b ", bit)
		}
		fmt.Println()
	}
	if groupedBits == nil{
		t.Errorf("Error: unable to print grouped bits")
	}
}
