package utils

type Coordinate struct {
	X int
	Y int
}

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

func ComputeMax(values []int) int {
	max := values[0]
	for _, value := range values {
		if value > max {
			max = value
		}
	}
	return max
}

func ComputeMin(values []int) int {
	min := values[0]
	for _, value := range values {
		if value < min {
			min = value
		}
	}
	return min
}
