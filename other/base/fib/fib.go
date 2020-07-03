package fib

func fib(n int64) int64 {
	var i, s1, s2 int64
	if n == 1 || n == 2 {
		return 1
	}
	s1 = 1
	s2 = 1
	for i = 3; i <= n; i++ {
		s2 = s1 + s2
		s1 = s2 - s1
	}
	return s2
}
