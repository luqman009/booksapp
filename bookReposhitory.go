package main

var bookCollection = make([]Book, 0)

type BookReposhitory struct{}

func (br BookReposhitory) AddMewBook(book Book) {
	bookCollection = append(bookCollection, book)
}

func (br BookReposhitory) FindAllBook() []Book {
	return bookCollection
}
func (br BookReposhitory) FindBookField(fnFilter func(Book) bool) []Book {
	var searchResult = make([]Book, 0)
	for _, b := range bookCollection {
		if fnFilter(b) {
			searchResult = append(searchResult, b)
		}
	}

	return searchResult

}
