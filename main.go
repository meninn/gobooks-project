package main

import (
	"fmt"
	"gobooks/internal/service"
)

func main() {
	book := service.Book{
		ID:     1,
		Title:  "Harry Potter and the Philosopher's Stone",
		Author: "J.K. Rowling",
		Genre:  "Fantasy",
	}

	fmt.Println(book.GetFullBook())
}
