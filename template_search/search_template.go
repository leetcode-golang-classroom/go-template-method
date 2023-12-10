package template_search

import (
	"fmt"
	"io"
	"os"
)

type SearchTemplate[T int | string] interface {
	InitialSearchValue() T
	UpdateSearchValue(t T, ref string) T
	SearchEnd(arr []string, value T, idx int) bool
}

type TemplateSearchClass[T int | string] struct {
	searcher SearchTemplate[T]
	writer   io.Writer
}

func NewTemplateClass[T int | string](s SearchTemplate[T]) *TemplateSearchClass[T] {
	return &TemplateSearchClass[T]{
		searcher: s,
		writer:   os.Stdout,
	}
}

func (t *TemplateSearchClass[T]) Search(arr []string) T {
	searchValue := t.searcher.InitialSearchValue()
	for i := 0; i < len(arr); i++ {
		fmt.Fprintln(t.writer, arr[i])
		searchValue = t.searcher.UpdateSearchValue(searchValue, arr[i])
		if t.searcher.SearchEnd(arr, searchValue, i) {
			break
		}
	}
	return searchValue
}

func (t *TemplateSearchClass[T]) SetWriter(writer io.Writer) {
	t.writer = writer
}
