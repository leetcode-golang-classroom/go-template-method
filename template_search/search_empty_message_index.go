package template_search

// input: N 個 string
// init value
// for i:= 0 .. len(s)== 0
// update index
// print string
// return idx
type SearchEmptyMessageIndex struct{}

func (s *SearchEmptyMessageIndex) InitialSearchValue() int {
	return 0
}
func (s *SearchEmptyMessageIndex) UpdateSearchValue(value int, ref string) int {
	if len(ref) != 0 {
		return value + 1
	}
	return value
}
func (s *SearchEmptyMessageIndex) SearchEnd(arr []string, value int, idx int) bool {
	return len(arr[idx]) == 0
}
func NewSearchEmptyMessageIndex() *SearchEmptyMessageIndex {
	return &SearchEmptyMessageIndex{}
}
