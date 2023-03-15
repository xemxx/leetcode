package main

import "fmt"

func main() {
	a := []int{2, 5, 16, 2, 2, 5, 5}

	fmt.Println(singleNumber(a))

}

// 按照二进制位每位相加最后除3求余
func singleNumber(nums []int) int {
	ans := int32(0)
	for i := 0; i < 32; i++ {
		total := int32(0)
		for _, num := range nums {
			total += int32(num) >> i & 1
		}
		if total%3 > 0 {
			ans |= 1 << i
		}
	}
	return int(ans)
}
