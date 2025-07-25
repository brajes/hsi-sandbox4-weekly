package main

import "fmt"

type Book struct {
	ID           int
	Title        string
	Author       string
	DynamicStock int
	Stock        int
}

type Library struct {
	books []Book
}

func CreateLibrary() *Library {
	return &Library{make([]Book, 0)}
}

func (l *Library) ListBooks() {
	fmt.Println("ID | Title | Author | Stock")
	for _, book := range l.books {
		fmt.Printf("%d | %s | %s | %d\n", book.ID, book.Title, book.Author, book.DynamicStock)
	}
}

func (l *Library) AddBook(id int, title string, author string, stock int) {
	existing_book := l.FindBook(id)

	if existing_book != nil {
		fmt.Println("Book with existing ID already exists")
		return
	}

	b := Book{id, title, author, stock, stock}
	l.books = append(l.books, b)
}

func (l *Library) FindBook(id int) *Book {
	for i, book := range l.books {
		if book.ID == id {
			return &l.books[i]
		}
	}
	return nil
}

func (l *Library) BorrowBook(id int, qty int) {
	book := l.FindBook(id)
	if book == nil {
		fmt.Println("Book not found.")
		return
	}
	if book.DynamicStock == 0 {
		fmt.Println("Book out off stocks.")
		return
	}

	if qty < 1 {
		fmt.Println("Quantity cannot be zero.")
		return
	}

	if book.DynamicStock < qty {
		fmt.Printf("Not enough stocks. Max %d book(s).\n", book.DynamicStock)
		return
	}

	book.DynamicStock = book.DynamicStock - 1
}

func (l *Library) ReturnBook(id int, qty int) {
	book := l.FindBook(id)

	if book == nil {
		fmt.Println("Book not found.")
		return
	}

	if qty < 1 {
		fmt.Println("Quantity cannot be zero.")
		return
	}

	if qty+book.DynamicStock > book.Stock {
		fmt.Println("Returned Quantity exceed original stocks.")
		return
	}

	book.DynamicStock = book.DynamicStock + qty
	fmt.Printf("%d books with title %s has been returned to library.\n", qty, book.Title)
}

func main() {
	library := CreateLibrary()
	library.AddBook(1, "Ushul Tsalatsah", "Syaikh Muhammad Abdul Wahab", 1)
	library.AddBook(2, "Fadhul Islam", "Syaikh Muhammad Abdul Wahab", 1)
	library.ListBooks()
	fmt.Println("")
	library.BorrowBook(1, 1)
	library.ListBooks()
	fmt.Println("")
	library.ReturnBook(1, 1)
	fmt.Println("")
	library.BorrowBook(2, 2)
	fmt.Println("")
	library.ReturnBook(1, 1)
}
