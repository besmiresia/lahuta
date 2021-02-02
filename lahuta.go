package main

import (
	"fmt"
	"github.com/GeertJohan/go.rice"
	"github.com/labstack/echo/v4"
	"html/template"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"sort"
	"strings"
)

type Book struct {
	Name string
	File string
}

var books = make(map[string]string)
var keys []string

func main() {
	getBooks()
	server()
}

func getBooks() {
	files, err := ioutil.ReadDir("./books")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	for _, f := range files {
		books[f.Name()] = strings.Title(
			strings.ReplaceAll(f.Name(), "_", " "))
	}
	fmt.Println(books)

	for k := range books {
		keys = append(keys, k)
	}
	sort.Strings(keys)

}

func server() {
	e := echo.New()
	e.Static("/assets", "public/assets")
	assetHandler := http.FileServer(
		rice.MustFindBox("books").HTTPBox(),
	)

	e.GET("/books/*",
		echo.WrapHandler(
			http.StripPrefix("/books/", assetHandler),
		),
	)
	e.Renderer = t
	e.GET("/books", BooksView)
	e.GET("/", BooksView)

	e.Logger.Fatal(e.Start(":3000"))
}

func BooksView(c echo.Context) error {
	var allBooks []Book
	for _, snake := range keys {
		allBooks = append(
			allBooks,
			Book{books[snake], "books/" + snake})
	}
	return c.Render(http.StatusOK, "books", allBooks)
}

var t = &Template{
	templates: template.Must(template.ParseGlob("public/views/*.html")),
}

type Template struct {
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}
