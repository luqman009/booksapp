package main

import (
	"strings"
)

type BookService struct {
	r *BookReposistory
}

func NewBookService(repo *BookReposistory) *BookService {
	return &BookService{r: repo}
}

func (bs *BookService) registerNewBook(b *Book) {
	bs.r.AddMewBook(b)
}

func (bs *BookService) getAllBookNewBook() []*Book {
	return bs.r.FindAllBook()
}

func (bs *BookService) searchBookId(id string) []*Book {
	var funcFinlter = func(b Book) bool {
		return b.Id == id
	}
	return bs.r.FindBookField(funcFinlter)

}

func (bs *BookService) searchBookByTitle(title string) []*Book {
	var funcFinlter = func(b Book) bool {
		return strings.Contains(b.Title, title)
	}
	return bs.r.FindBookField(funcFinlter)

}
