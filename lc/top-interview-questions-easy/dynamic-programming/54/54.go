package _54

func climbStairs(n int) int {
	switch n {
	case 0, 1:
		return 1
	case 2:
		return 2
	}
	result := make([]int, n+1)
	result[1] = 1
	result[2] = 2
	for i := 3; i <= n; i++ {
		result[i] = result[i-1] + result[i-2]
	}
	return result[n]
}

func climbStairsBetter(n int) int {
	var dp [3]int
	dp[0] = 1
	dp[1] = 1
	for i := 2; i <= n; i++ {
		dp[i%3] = dp[(i-1)%3] + dp[(i-2)%3]
	}
	return dp[n%3]
}

func climbStairsBetter2(n int) int {
	var p, q, r = 0, 0, 1
	for i := 1; i <= n; i++ {
		p = q
		q = r
		r = p + q
	}
	return r
}
