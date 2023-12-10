package template_search

import "strings"

// input： N 個 string
// inital count
// for i= 0 ... N-1
// update count
// print i message
// return count
type CountNumberWaterBalls struct{}

func NewCountNumberWaterBalls() *CountNumberWaterBalls {
	return &CountNumberWaterBalls{}
}
func (c *CountNumberWaterBalls) InitialSearchValue() int {
	return 0
}
func (c *CountNumberWaterBalls) UpdateSearchValue(value int, ref string) int {
	if strings.Compare("Waterball", ref) == 0 {
		return value + 1
	}
	return value
}
func (s *CountNumberWaterBalls) SearchEnd(arr []string, value int, idx int) bool {
	return idx == len(arr)-1
}
