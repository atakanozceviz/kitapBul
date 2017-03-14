package controller

import (
	"log"

	"github.com/PuerkitoBio/goquery"
	"github.com/atakanozceviz/kitapBul/model"
	"github.com/headzoo/surf"
)

func Idefix(books *model.Books, s string) {
	defer wg.Done()
	bow := surf.NewBrowser()
	err := bow.Open("http://www.idefix.com/search?q=" + s)
	if err != nil {
		log.Fatal(err)
	} else {
		bow.Find(".list-cell").Each(func(index int, item *goquery.Selection) {
			title := item.Find(".item-name h3").Text()
			author := item.Find(".who").Text()
			pub := item.Find(".mb10").Text()
			img, _ := item.Find("figure img").Attr("src")
			price := item.Find(".price").Text()

			if title != "" && price != "" {
				p := model.Book{
					Title:     title,
					Author:    author,
					Publisher: pub,
					Img:       img,
					Price:     price,
					WebSite:   "www.idefix.com",
				}
				model.Add(p, books)
			}

		})
	}
}
