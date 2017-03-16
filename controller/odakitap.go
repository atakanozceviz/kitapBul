package controller

import (
	"log"

	"github.com/atakanozceviz/kitapBul/model"
	"github.com/headzoo/surf"
)

func Odakitap(books *model.Books, s string) {
	defer wg.Done()
	bow := surf.NewBrowser()
	err := bow.Open("https://www.odakitap.com/index.php?p=Products&q=" + s)
	if err != nil {
		log.Fatal(err)
	}
	item := bow.Find(".main-content")

	title := item.Find(".pd-name").Text()
	author := item.Find(".pd-owner a").Text()
	pub := item.Find(".pd-publisher a span").Text()
	img, _ := item.Find("#main_img").Attr("src")
	price := item.Find("#prd_final_price_display").Text()
	website := bow.Url().String()

	if title != "" && price != "" {
		p := model.Book{
			Title:     title,
			Author:    author,
			Publisher: pub,
			Img:       "https://www.odakitap.com" + img,
			Price:     price,
			WebSite:   website,
		}
		lock.Lock()
		model.Add(p, books)
		lock.Unlock()
	}

}
