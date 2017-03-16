package controller

import (
	"log"

	"github.com/atakanozceviz/kitapBul/model"
	"github.com/headzoo/surf"
)

func Pandora(books *model.Books, s string) {
	defer wg.Done()
	bow := surf.NewBrowser()
	err := bow.Open("http://www.pandora.com.tr/Arama/?type=9&kitapadi=&yazaradi=&yayinevi=&isbn=" + s + "&resimli=1&dil=0&sirala=0")
	if err != nil {
		log.Fatal(err)
	}
	item := bow.Find(".urunorta")

	title := item.Find(".kt").Text()
	author := item.Find(".yz").Text()
	pub := item.Find(".yy").Text()
	img, _ := item.Find(".imgcont img").Attr("src")
	price := item.Find(".fyt strong").Text()
	website, _ := item.Find(".imgcont a").Attr("href")
	if title != "" && price != "" {
		p := model.Book{
			Title:     title,
			Author:    author,
			Publisher: pub,
			Img:       "http://www.pandora.com.tr" + img,
			Price:     price,
			WebSite:   "http://www.pandora.com.tr" + website,
		}
		lock.Lock()
		model.Add(p, books)
		lock.Unlock()
	}

}
