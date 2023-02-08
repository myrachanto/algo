package main

import "log"

func init() {
	log.SetPrefix("Greedy Method --")
}
func main() {
	// sort the intervals by left and order by right, e.g. goal =>[[1,2], [1,3], [1,4],[2,3],[3,4]]
	// store the target range (left, right) for an arrow
	// if current ballon is within the range of (left, right), update the range
	// if the current ballon is not in the range of (left, right), fire an arrow, update the range with the current ballon coordinates
	// count the number of arrows we have fired
	select {}

}

func findMinArrowShots(points [][]int) int {
	if len(points) < 2 {
		return len(points)
	}
	// sort the points by left , order by right
	return 0
}
