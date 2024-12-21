package top100

func maxArea(height []int) int {
	n := len(height)
	ans := min(height[0], height[n-1]) * (n - 1)
	i := 0
	j := n - 1
	for i < j {
		// 取左边柱子场景
		left := i + 1
		right := j
		res1 := min(height[left], height[right]) * (right - left - 1)
		// 取右边柱子场景
		left = i
		right = j - 1
		res2 := min(height[left], height[right]) * (right - left - 1)
		if res1 > res2 {
			i++
			ans = max(ans, res1)
		} else {
			j--
			ans = max(ans, res2)
		}
	}
	return ans
}
