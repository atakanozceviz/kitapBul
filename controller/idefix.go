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
			a := item.Find(".item-name")
			title := a.Find("h3").Text()
			author := item.Find(".who").First().Text()
			pub := item.Find(".mb10").Text()
			img, _ := item.Find("figure img").Attr("src")
			price := item.Find(".price").Text()
			website, _ := a.Attr("href")
			if title != "" && price != "" {
				p := model.Book{
					Title:     title,
					Author:    author,
					Publisher: pub,
					Img:       img,
					Price:     price,
					WebSite:   "http://www.idefix.com" + website,
				}
				lock.Lock()
				model.Add(p, books)
				lock.Unlock()
			}

		})
	}
}
