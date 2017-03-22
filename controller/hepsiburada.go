package controller

import (
	"log"

	"github.com/PuerkitoBio/goquery"
	"github.com/atakanozceviz/kitapbul/model"
	"github.com/headzoo/surf"
)

func Hepsiburada(books *model.Books, s string) {
	defer wg.Done()
	bow := surf.NewBrowser()
	err := bow.Open("http://www.hepsiburada.com/ara?q=" + s)
	if err != nil {
		log.Println(err)
	} else {
		bow.Find(".product").Each(func(index int, item *goquery.Selection) {
			title := item.Find(".product-title p").Text()
			author := item.Find(".who").First().Text()
			pub := ""
			img, _ := item.Find("img").Attr("src")
			price := item.Find(".product-price").Text()
			website, _ := item.Find("a").First().Attr("href")
			if title != "" && price != "" {
				p := model.Book{
					Title:     title,
					Author:    author,
					Publisher: pub,
					Img:       img,
					Price:     price,
					WebSite:   "http://www.hepsiburada.com" + website,
				}
				lock.Lock()
				model.Add(p, books)
				lock.Unlock()
			}

		})
	}
}
