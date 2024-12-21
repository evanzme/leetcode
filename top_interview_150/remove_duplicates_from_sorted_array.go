package top_interview_150

/*
输入：nums = [0,0,1,1,1,2,2,3,3,4]
输出：5, nums = [0,1,2,3,4]
*/
func removeDuplicates(nums []int) int {
	numMap := make(map[int]bool)
	var num = 0
	for i := 0; i < len(nums); i++ {
		if !numMap[nums[i]] {
			nums[num] = nums[i]
			num++
		}
		numMap[nums[i]] = true
	}
	return num
}
