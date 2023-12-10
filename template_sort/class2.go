package template_sort

type Class2 struct{}

func NewClass2() *Class2 {
	return &Class2{}
}
func (c2 *Class2) Compare(i, j int, arr []int) bool {
	return arr[i] < arr[j]
}
