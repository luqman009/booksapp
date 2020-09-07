package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type BookDelivery struct {
	bookService *BookService
}

func NewBookDelivery(c *AppConfig) *BookDelivery {
	repo := NewBookRepository(c.getDbPath(), c.getDb())
	bookService := NewBookService(repo)
	return &BookDelivery{bookService}
}

func (bd *BookDelivery) create() {
	var isExit = false
	var userChoice string

	bd.mainMenuForm()
	for isExit == false {
		fmt.Printf("\n%s", "Your choice : ")
		fmt.Scanln(&userChoice)
		switch userChoice {
		case "01":
			bd.registerationBookForm()
		case "02":
			bd.viewBookCollectionForm()
		case "q":
			isExit = true
		default:
			fmt.Println("Unknown menu code")
		}
	}
}

func (bd *BookDelivery) menuChoiceOrdered(m map[string]string) []string {
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	return keys
}

func (bd *BookDelivery) mainMenuForm() {
	var appMenu = map[string]string{
		"01": "Add book to collection",
		"02": "View book to collection",
		"q":  " Quit aplication",
	}
	fmt.Printf("%s\n", strings.Repeat("*", 30))
	fmt.Printf("%26s\n", "My 1st Book Conllection")
	fmt.Printf("%s\n", strings.Repeat("*", 30))
	for _, menuCode := range bd.menuChoiceOrdered(appMenu) {
		fmt.Printf("%s %s\n", menuCode, appMenu[menuCode])
	}
}

func (bd *BookDelivery) registerationBookForm() {
	consoleClear()
	var title string
	var author string
	var genre string
	var price float64
	var confirmation string
	fmt.Println()
	fmt.Printf("%s\n", "Book Registration Form")
	fmt.Printf("%s\n", strings.Repeat("-", 30))
	scanner := bufio.NewReader(os.Stdin)
	fmt.Print("Title : ")
	sTitle, _ := scanner.ReadString('\n')
	title = strings.TrimSpace(sTitle)
	fmt.Print("Author : ")
	sAuthhor, _ := scanner.ReadString('\n')
	author = strings.TrimSpace(sAuthhor)
	fmt.Print("Genre : ")
	sGenre, _ := scanner.ReadString('\n')
	genre = strings.TrimSpace(sGenre)
	fmt.Print("Price : ")
	sPrice, _ := scanner.ReadString('\n')
	price, _ = strconv.ParseFloat(strings.TrimSpace(sPrice), 64)

	fmt.Println("Save to collection ? : (Y/N)")
	fmt.Scanln(&confirmation)

	if confirmation == "y" {
		newBook := NewBook(title, author, genre, price)
		bd.bookService.registerNewBook(&newBook)
	}
	consoleClear()
	bd.mainMenuForm()
}

func (bd *BookDelivery) viewBookCollectionForm() {
	consoleClear()
	books := bd.bookService.getAllBookNewBook()
	fmt.Println("")
	fmt.Printf("My Book Collection")
	fmt.Printf("%s\n", strings.Repeat("=", 115))
	fmt.Printf("%-40s%-30s%-15s%-15s%15s\n", "ID", "Title", "Author", "Genre", "Price (Rp.)")
	fmt.Printf("%s\n", strings.Repeat("=", 115))
	for _, b := range books {
		fmt.Printf("%-40s%-30s%-15s%-15s%15.2f\n", b.Id, b.Title, b.Author, b.Genre, b.Price)
	}
	var kembali string
	fmt.Print("Back To menu ? : (Y)")
	fmt.Scanln(&kembali)
	if kembali == "y" {
		consoleClear()
		bd.mainMenuForm()
	}
	// for {
	// 	var kembali string
	// 	fmt.Print("Back To menu ? : (Y)")
	// 	fmt.Scanln(&kembali)
	// 	if kembali == "y" {
	// 		consoleClear()
	// 		bd.mainMenuForm()
	// 	} else {
	// 		fmt.Println("inputan salah")
	// 	}

	// }
}
