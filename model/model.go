package model

import (
	"encoding/json"
	"log"
)

type Book struct {
	Title     string `json:"title"`
	Author    string `json:"author"`
	Publisher string `json:"publisher"`
	Img       string `json:"img"`
	Price     string `json:"price"`
	WebSite   string `json:"website"`
}

type Books []Book

func Add(b Book, bs *Books) {
	*bs = append(*bs, Book{
		Title:     b.Title,
		Author:    b.Author,
		Publisher: b.Publisher,
		Img:       b.Img,
		Price:     b.Price,
		WebSite:   b.WebSite,
	})
}

func (bs *Books) ToJson() []byte {
	j, err := json.Marshal(bs)
	if err != nil {
		log.Println(err)
	}
	return j
}
