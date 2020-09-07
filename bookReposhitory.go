package main

import (
	"crypto/md5"
	"fmt"
)

// var bookCollection = make([]Book, 0)

type BookReposistory struct {
	bookCollection      *[]*Book
	bookCollectionInfra *BookRepositoryInfrastructure
}

func NewBookRepository(dataPath string, bookCollection []*Book) *BookReposistory {
	bookRepoInfrac := NewBookRepoInfra(dataPath)
	return &BookReposistory{bookCollection: &bookCollection, bookCollectionInfra: bookRepoInfrac}
}

func (br *BookReposistory) AddMewBook(book *Book) {
	data := []byte(book.Title)
	book.Id = fmt.Sprintf("%x", md5.Sum(data))
	*br.bookCollection = append(*br.bookCollection, book)
	br.bookCollectionInfra.saveToFile(br.bookCollection)
}

func (br *BookReposistory) FindAllBook() []*Book {
	br.bookCollectionInfra.readFile(br.bookCollection)
	return *br.bookCollection
}
func (br BookReposistory) FindBookField(fnFilter func(Book) bool) []*Book {
	br.bookCollectionInfra.readFile(br.bookCollection)
	var searchResult = make([]*Book, 0)
	for _, b := range *br.bookCollection {
		if fnFilter(*b) {
			searchResult = append(searchResult, b)
		}
	}
	return searchResult
}
