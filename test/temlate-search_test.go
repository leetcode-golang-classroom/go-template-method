package test

import (
	"bytes"
	"reflect"
	"testing"

	"github.com/leetcode-golang-classroom/go-template-method/template_search"
)

func WithStringArr(t testing.TB) []string {
	return []string{"Waterball", "test", "abc", "Who Is Your Daddy", "Jack", "", "gsonliang", "Waterball", "Waterball"}
}

func TestCountNumberOfWaterBalls(t *testing.T) {
	arr := WithStringArr(t)
	expected := 3
	template := template_search.NewTemplateClass[int](template_search.NewCountNumberWaterBalls())
	var out bytes.Buffer
	template.SetWriter(&out)
	got := template.Search(arr)
	if !reflect.DeepEqual(got, expected) {
		t.Errorf("expected: %v, got: %v", expected, got)
	}
}
func TestSearchEmptyMessageIndex(t *testing.T) {
	arr := WithStringArr(t)
	expected := 5
	template := template_search.NewTemplateClass[int](template_search.NewSearchEmptyMessageIndex())
	var out bytes.Buffer
	template.SetWriter(&out)
	got := template.Search(arr)
	if !reflect.DeepEqual(got, expected) {
		t.Errorf("expected: %v, got: %v", expected, got)
	}
}
func TestSearchLongestMessage(t *testing.T) {
	arr := WithStringArr(t)
	expected := "Who Is Your Daddy"
	template := template_search.NewTemplateClass[string](template_search.NewSearchLongestMessage())
	var out bytes.Buffer
	template.SetWriter(&out)
	got := template.Search(arr)
	if !reflect.DeepEqual(got, expected) {
		t.Errorf("expected: %v, got: %v", expected, got)
	}
}
