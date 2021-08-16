package leetcode

func twoSum(nums []int, target int) []int {
	var m map[int]int
	m = make(map[int]int)
	for i := 0; i < len(nums); i++ {
		m[nums[i]] = i
	}
	for i := 0; i < len(nums); i++ {
		complement := target - nums[i]
		val, exists := m[complement]
		if exists == true && val != i {
			j := m[complement]
			return []int{i, j}
		}
	}
	return []int{-1, -1}
}
