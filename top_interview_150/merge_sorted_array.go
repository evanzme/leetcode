package top_interview_150

import "sort"

/*
 [1, 4, 6, 0, 0, 0, 0] 3

 [2, 5, 6, 7] 4
*/

func merge(nums1 []int, m int, nums2 []int, n int) {
	copy(nums1[m:], nums2)
	sort.Ints(nums1)
}
