package top_interview_150

/*
输入：nums = [1,1,1,2,2,3]
输出：5, nums = [1,1,2,2,3]
*/
func removeDuplicatesII(nums []int) int {
	var res, sameLen = 1, 0
	for i := 1; i < len(nums); i++ {
		if nums[i] == nums[i-1] {
			sameLen++
		} else {
			sameLen = 0
		}
		if sameLen == 2 {

			continue
		}
		nums[res] = nums[i]
		res++
	}
	return res
}
