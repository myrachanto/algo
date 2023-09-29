package main

func getSum(d []int) int {
	if len(d) == 0 {
		return 0
	}
	return d[0] + getSum(d[1:])
}
func getSum2(d []int) (sum int) {
	for _, k := range d {
		sum += k
	}
	return sum
}
func removeItem(item int, d []int) []int {
	res := make([]int, 0, len(d))
	for _, g := range d {
		if g != item {
			res = append(res, g)
		}
	}
	return res
}
func removeItem2(item int, d []int) []int {
	return nil
}
