package leetcode

import (
	"fmt"
	"math"
)

func maxSubArray(nums []int) int {
	var (
		max           = math.MinInt32
		sum           int
		start, end, f int
	)
	for i, j := range nums {
		sum += j
		if sum > max {
			if f != 0 {
				start = f + 1
				f = 0
			}
			max = sum
			end = i
		}
		if sum < 0 {
			f = i
			sum = 0
		}
	}
	fmt.Printf("start index is %d, end index is %d\n", start, end)
	return max
}
