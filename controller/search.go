package controller

import (
	"sync"

	"github.com/atakanozceviz/kitapbul/model"
)

var wg sync.WaitGroup
var lock sync.Mutex

func Search(books *model.Books, s string) *model.Books {
	wg.Add(4)
	go Idefix(books, s)
	go Odakitap(books, s)
	go Pandora(books, s)
	go Hepsiburada(books, s)
	wg.Wait()
	return books
}
