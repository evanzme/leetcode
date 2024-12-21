package top100

/*
输入：m = 3, n = 7
输出：28
0, 1, 1, 1,  1,  1,  1
1, 2, 3, 4,  5,  6,  7
1, 3, 6, 10, 15, 21, 28

输入：m = 3, n = 2
输出：3

0 1 1
1 2 3
1 3 6
*/
func uniquePaths(m int, n int) int {
	dp := make(map[int]map[int]int)
	dp[0] = map[int]int{0: 0}
	for i := 1; i <= n; i++ {
		if _, ok := dp[i]; !ok {
			dp[i] = make(map[int]int)
		}
		for j := 1; j <= m; j++ {
			if i == 1 || j == 1 {
				dp[i][j] = 1
			} else {
				dp[i][j] = dp[i-1][j] + dp[i][j-1]
			}
		}
	}
	return dp[n][m]
}
