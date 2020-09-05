package main

import (
	"fmt"
	"strings"
)

type BookService struct {
	r BookReposhitory
}

func (bs BookService) registerNewBook(b Book) {
	bs.r.AddMewBook(b)

}

func (bs BookService) printAllBookCollection() {
	collection := bs.r.FindAllBook()
	for _, books := range collection {
		bs.printConsole("koleksibuku" collection)
	}
}

func (bs BookService) searchBookId(id string) {
	var funcFinlter = func(b Book) bool {
		if b.Id == id {
			return true
		}
		return true

	}
	result := bs.r.FindBookField(funcFinlter)
	bs.printConsole(fmt.Sprintf("Hasil Pencarian Buku dengan ID : %s", id), result)

}

func (bs BookService) searchBookByTitle(title string) {
	var funcFinlter = func(b Book) bool {
		return strings.Contains(b.Titile, title)
	}
	result := bs.r.FindBookField(funcFinlter)
	bs.printConsole(fmt.Sprintf("Hasil pencarian Buku dengan kata kunci : %s", title), result)
}

func (bs BookService) printConsole(header string, books []Book) {
	fmt.Println("")
	fmt.Println(header)
	fmt.Printf("%s\n", strings.Repeat("=", 80))
	fmt.Printf("%-10s%-30s%-15s%-15s%10s\n", "ID", "Title", "Author", "Genre", "Price")
	fmt.Printf("%s\n", strings.Repeat("-", 80))
	for _, b := range books {
		fmt.Printf("%-10s%-30s%-15s%-15s%.2f\n", b.Id, b.Titile, b.Author, b.Genre, b.Price)
	}
}
