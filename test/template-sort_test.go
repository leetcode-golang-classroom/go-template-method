package test

import (
	"reflect"
	"testing"

	"github.com/leetcode-golang-classroom/go-template-method/template_sort"
)

func WithArr(t testing.TB) []int {
	return []int{0, 3, 100, 2, 25, 200, -199, 7}
}

func TestClass1(t *testing.T) {
	originArr := WithArr(t)
	expected := []int{-199, 0, 2, 3, 7, 25, 100, 200}
	t1 := template_sort.NewTemplateClass(template_sort.NewClass1())
	t1.SortMethod(originArr)
	if !reflect.DeepEqual(originArr, expected) {
		t.Errorf("expected: %v, got: %v\n", expected, originArr)
	}
}

func TestClass2(t *testing.T) {
	originArr := WithArr(t)
	expected := []int{200, 100, 25, 7, 3, 2, 0, -199}
	t1 := template_sort.NewTemplateClass(template_sort.NewClass2())
	t1.SortMethod(originArr)
	if !reflect.DeepEqual(originArr, expected) {
		t.Errorf("expected: %v, got: %v\n", expected, originArr)
	}
}
