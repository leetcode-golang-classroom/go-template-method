package main

import (
	"fmt"

	"github.com/leetcode-golang-classroom/go-template-method/template_search"
)

func main() {
	arr := []string{"Waterball", "test", "abc", "Who Is Your Daddy", "Jack", "", "gsonliang", "Waterball", "Waterball"}
	t1 := template_search.NewTemplateClass[int](template_search.NewCountNumberWaterBalls())
	result := t1.Search(arr)
	fmt.Printf("[count number waterball]: %v\n", result)
	t2 := template_search.NewTemplateClass[int](template_search.NewSearchEmptyMessageIndex())
	result = t2.Search(arr)
	fmt.Printf("[find empty index]: %v\n", result)
	t3 := template_search.NewTemplateClass[string](template_search.NewSearchLongestMessage())
	result3 := t3.Search(arr)
	fmt.Printf("[find longest string]: %v\n", result3)
}
