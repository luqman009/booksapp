package main

type Book struct {
	Id     string
	Title  string
	Author string
	Genre  string
	Price  float64
}

func NewBook(title string, author string, genre string, price float64) Book {
	return Book{
		Title:  title,
		Author: author,
		Genre:  genre,
		Price:  price,
	}
}
