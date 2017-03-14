package main

import (
	"os"
	"runtime"

	"strings"

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
	app.Get("/", search)
	app.Get("/:params", search2)
	app.Listen(":" + port)
}

func search(ctx *iris.Context) {
	k := ctx.URLParam("keyword")
	if k != "" {
		var books model.Books
		ctx.JSON(iris.StatusOK, controller.Search(&books, k))
	}
}
func search2(ctx *iris.Context) {
	rep := strings.NewReplacer("/", "")
	k := rep.Replace(ctx.Request.URL.Path)
	if k != "" {
		var books model.Books
		ctx.JSON(iris.StatusOK, controller.Search(&books, k))
	}
}
