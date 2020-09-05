package main

func main() {

	book01 := Book{"01", "Dragon Ball 01", "Akira Toriyama", "adventure", 80000}
	book02 := Book{"02", "Dragon Ball 02", "Akira Toriyama", "adventure", 80000}
	book03 := Book{"03", "Dragon Ball 03", "Akira Toriyama", "adventure", 80000}

	repo := BookReposhitory{}
	bookService := BookService{r: repo}

	bookService.registerNewBook(book01)
	bookService.registerNewBook(book02)
	bookService.registerNewBook(book03)
	bookService.printAllBookCollection()

	bookService.searchBookByTitle("dragon")

	bookService.searchBookId("01")

}
