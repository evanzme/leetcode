package top100

/*
[4,2,0,3,2,5]

	+

+    +
+  + +
++ +++
++ +++
-------

[0,1,0,2,1,0,1,3,2,1,2,1]
*/
func trap(height []int) int {
	if len(height) == 0 {
		return 0
	}
	maxHeight := height[0]
	for i := 1; i < len(height); i++ {
		maxHeight = max(maxHeight, height[i])
		if height[i] >= height[i-1] {
			continue
		}

	}
	return maxHeight
}
