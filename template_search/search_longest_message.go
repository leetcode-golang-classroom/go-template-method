package template_search

// input: n 個 string
// initial value
// for i= 0 .. n-1
// update laragest string
// print ith string
// return largest string
type SearchLongestMessage struct{}

func (s *SearchLongestMessage) InitialSearchValue() string {
	return ""
}
func (s *SearchLongestMessage) UpdateSearchValue(value string, ref string) string {
	if len(value) < len(ref) {
		return ref
	}
	return value
}

func (s *SearchLongestMessage) SearchEnd(arr []string, value string, idx int) bool {
	return idx == len(arr)-1
}
func NewSearchLongestMessage() *SearchLongestMessage {
	return &SearchLongestMessage{}
}
