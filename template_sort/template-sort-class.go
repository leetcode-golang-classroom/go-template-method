package template_sort

type TemplateClass struct {
	comparator CompareHandle
}
type CompareHandle interface {
	Compare(i, j int, arr []int) bool
}

func NewTemplateClass(c CompareHandle) *TemplateClass {
	return &TemplateClass{
		comparator: c,
	}
}

func (t *TemplateClass) SortMethod(arr []int) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if t.comparator.Compare(j, j+1, arr) {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
}
