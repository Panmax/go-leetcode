package main

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	length := len(nums1) + len(nums2)
	if length%2 == 1 {
		return float64(getKthElement(nums1, nums2, length/2+1))
	} else {
		return float64(getKthElement(nums1, nums2, length/2)+getKthElement(nums1, nums2, length/2+1)) / 2.0
	}
}

func getKthElement(nums1 []int, nums2 []int, k int) int {
	index1, index2 := 0, 0
	for {
		if len(nums1) == index1 {
			return nums2[index2+k-1]
		}
		if len(nums2) == index2 {
			return nums1[index1+k-1]
		}
		if k == 1 {
			return min(nums1[index1], nums2[index2])
		}
		half := k / 2
		newIndex1 := min(index1+half, len(nums1)) - 1
		newIndex2 := min(index2+half, len(nums2)) - 1
		pivot1 := nums1[newIndex1]
		pivot2 := nums2[newIndex2]
		if pivot1 <= pivot2 {
			k -= newIndex1 - index1 + 1
			index1 = newIndex1 + 1
		} else {
			k -= newIndex2 - index2 + 1
			index2 = newIndex2 + 1
		}
	}
	return 0
}

func min(i, j int) int {
	if i < j {
		return i
	}
	return j
}
