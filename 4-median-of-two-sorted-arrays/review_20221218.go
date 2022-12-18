package main

func findMedianSortedArrays_20221218(nums1 []int, nums2 []int) float64 {
	length := len(nums1) + len(nums2)
	if length%2 == 1 {
		return float64(getKthElement_20221218(nums1, nums2, length/2+1))
	} else {
		return float64(getKthElement_20221218(nums1, nums2, length/2)+getKthElement_20221218(nums1, nums2, length/2+1)) / 2
	}
}

func getKthElement_20221218(nums1, nums2 []int, k int) int {
	var index1, index2 int
	for {
		// index1 == len(nums1)，不用+k，不是>
		if index1 == len(nums1) {
			return nums2[index2+k-1]
		}
		if index2 == len(nums2) {
			return nums1[index1+k-1]
		}
		if k == 1 {
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
