package controller

import (
	"sync"

	"github.com/atakanozceviz/kitapBul/model"
)

var wg sync.WaitGroup

func Search(books *model.Books, s string) *model.Books {
	wg.Add(3)
	go Idefix(books, s)
	go Odakitap(books, s)
	go Pandora(books, s)
	wg.Wait()
	return books
}
