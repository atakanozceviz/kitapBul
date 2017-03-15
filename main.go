package main

import (
	"os"
	"runtime"

	"strings"

	"log"

	"github.com/atakanozceviz/kitapBul/controller"
	"github.com/atakanozceviz/kitapBul/model"
	"gopkg.in/kataras/iris.v6"
	"gopkg.in/kataras/iris.v6/adaptors/httprouter"
)

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func main() {
	port := os.Getenv("PORT")
	app := iris.New()
	app.Adapt(
		iris.DevLogger(),
		httprouter.New(),
	)
	app.Get("/", searchKeyword)
	app.Get("/:params", searchUrl)
	app.Get("/jsonp/:params", searchJsonp)
	app.Listen(":" + port)
}

func searchKeyword(ctx *iris.Context) {
	k := ctx.URLParam("keyword")
	if k != "" {
		var books model.Books
		ctx.JSON(iris.StatusOK, controller.Search(&books, k))
	}
}

func searchUrl(ctx *iris.Context) {
	rep := strings.NewReplacer("/", "")
	k := rep.Replace(ctx.Request.URL.Path)
	if k != "" {
		var books model.Books
		ctx.JSON(iris.StatusOK, controller.Search(&books, k))
	}
}

func searchJsonp(ctx *iris.Context) {
	k := ctx.URLParam("keyword")
	if k != "" {
		var books model.Books
		err := ctx.JSONP(iris.StatusOK, ctx.URLParam("callback"), controller.Search(&books, k))
		if err != nil {
			log.Println(err)
		}
	}
}
