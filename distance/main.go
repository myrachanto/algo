package main

import (
	"fmt"
	"math"
	"sort"
	"time"
)

type (
	Point struct{ x, y float64 }
	line  struct{ begin, end Point }
)

func (l line) Distance() float64 {
	return math.Hypot(l.end.x-l.begin.x, l.end.y-l.begin.y)
}
func (l line) String() string {
	return fmt.Sprintf("%.2f is the distance better yet the hypotenues", l.Distance())
}
func main() {
	side := line{Point{1, 2}, Point{4, 6}}
	fmt.Println(side)
	ls := []string{"mary", "anthony", "jane", "bob"}
	// sort.Sort(sort.StringSlice(ls))
	sort.Strings(ls)
	fmt.Println(ls)
	timer := time.Now()
	hour := timer.Hour()
	fmt.Println("=========", hour)
	tim := timer.Add(-time.Hour * time.Duration(hour))
	fmt.Println("=========", tim)

	times := timer.String()
	fmt.Println(times[:10])

}
