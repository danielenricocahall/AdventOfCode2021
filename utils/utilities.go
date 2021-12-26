package utils

func Map(vs []string, f func(string) (int, error)) []int {
	vsm := make([]int, len(vs))
	for i, v := range vs {
		vsm[i], _ = f(v)
	}
	return vsm
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func Min(a, b int) int {
	return -Max(-a, -b)
}
