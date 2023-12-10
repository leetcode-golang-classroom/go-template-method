package template_sort

type Class1 struct{}

func NewClass1() *Class1 {
	return &Class1{}
}
func (c1 *Class1) Compare(i, j int, arr []int) bool {
	return arr[i] > arr[j]
}
