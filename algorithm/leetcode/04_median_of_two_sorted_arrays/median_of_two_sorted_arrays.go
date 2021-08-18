package leetcode

/**
* time complexity: O(log(m+n)), where m is nums1's length, n is nums2's length.
* space complexity: O(log(m+n)) 
*/
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	var (
		m = len(nums1)
		n = len(nums2)
	)

	if 0 == m && 0 == n {
		panic("nums1 and nums2 cannot be both empty.")
	}

	if 0 == m {
		return float64(nums2[n>>1]+nums2[n>>1-(n+1)%2]) / 2
	}

	if 0 == n {
		return float64(nums1[m>>1]+nums1[m>>1-(m+1)%2]) / 2
	}

	var (
		left  = (m + n + 1) >> 1
		right = (m + n + 2) >> 1
	)
	return float64(findKth(nums1, 0, nums2, 0, left)+findKth(nums1, 0, nums2, 0, right)) / 2
}

func findKth(nums1 []int, i int, nums2 []int, j int, k int) int {
	if i >= len(nums1) {
		return nums2[j+k-1]
	}
	if j >= len(nums2) {
		return nums1[i+k-1]
	}

	if 1 == k {
		if nums1[i] < nums2[j] {
			return nums1[i]
		}
		return nums2[j]
	}
	var (
		midVal1, midVal2 int
	)

	if i+k/2-1 < len(nums1) {
		midVal1 = nums1[i+k/2-1]
	} else {
		midVal1 = int(^uint(0) >> 1)
	}

	if j+k/2-1 < len(nums2) {
		midVal2 = nums2[j+k/2-1]
	} else {
		midVal2 = int(^uint(0) >> 1)
	}

	if midVal1 < midVal2 {
		return findKth(nums1, i+k/2, nums2, j, k-k/2)
	}
	return findKth(nums1, i, nums2, j+k/2, k-k/2)
}
