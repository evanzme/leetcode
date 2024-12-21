package top100

import "sort"

/*
给你一个整数数组 nums ，判断是否存在三元组 [nums[i], nums[j], nums[k]] 满足 i != j、i != k 且 j != k ，同时还满足 nums[i] + nums[j] + nums[k] == 0 。请你返回所有和为 0 且不重复的三元组。

注意：答案中不可以包含重复的三元组。

[-1,0,1,2,-1,-4]
*/
func threeSum(nums []int) [][]int {
	ans := make([][]int, 0)
	numFlag := make(map[int]map[int]bool)
	ansExist := make(map[int]map[int]int)
	for i := 0; i < len(nums); i++ {
		if _, ok := numFlag[nums[i]]; !ok {
			numFlag[nums[i]] = make(map[int]bool)
		}
		numFlag[nums[i]][i] = true
	}
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			isOk := false
			diff := 0 - nums[i] - nums[j]
			if posMap, ok := numFlag[diff]; ok {
				for v := range posMap {
					if v != i && v != j {
						isOk = true
						break
					}
				}
			}
			if isOk {
				var ansArr = []int{diff, nums[i], nums[j]}
				sort.Ints(ansArr)
				_, ok := ansExist[ansArr[0]]
				if !ok {
					ans = append(ans, []int{ansArr[0], ansArr[1], ansArr[2]})
				} else if _, ok := ansExist[ansArr[0]][ansArr[1]]; !ok {
					ans = append(ans, []int{ansArr[0], ansArr[1], ansArr[2]})
				} else if ansExist[ansArr[0]][ansArr[1]] != ansArr[2] {
					ans = append(ans, []int{ansArr[0], ansArr[1], ansArr[2]})
				}
				// 标记
				if _, ok := ansExist[ansArr[0]]; !ok {
					ansExist[ansArr[0]] = make(map[int]int)
				}
				ansExist[ansArr[0]][ansArr[1]] = ansArr[2]
			}
		}
	}
	return ans
}
