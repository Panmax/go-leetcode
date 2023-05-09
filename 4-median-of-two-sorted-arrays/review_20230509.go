package main

func findMedianSortedArrays_20230509(nums1 []int, nums2 []int) float64 {
	n := len(nums1) + len(nums2)
	if n%2 == 1 {
		// n/2+1
		return float64(getKthElement_20230509(nums1, nums2, n/2+1))
	} else {
		return float64(getKthElement_20230509(nums1, nums2, n/2)+getKthElement_20230509(nums1, nums2, n/2+1)) / 2
	}
}

func getKthElement_20230509(nums1 []int, nums2 []int, k int) int {
	var index1, index2 int
	for {
		if index1 == len(nums1) {
			return nums2[index2+k-1]
		}
		if index2 == len(nums2) {
			return nums1[index1+k-1]
		}
		if k == 1 {
			// 使用index1, index2，而不是0
			return min(nums1[index1], nums2[index2])
		}
		half := k / 2
		newIndex1 := min(index1+half, len(nums1)) - 1
		newIndex2 := min(index2+half, len(nums2)) - 1
		if nums1[newIndex1] < nums2[newIndex2] {
			k = k - newIndex1 + index1 - 1
			index1 = newIndex1 + 1
		} else {
			k = k - newIndex2 + index2 - 1
			index2 = newIndex2 + 1
		}
	}
}
