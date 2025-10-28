package main

import (
	"book-tool-cli/internal/app"
	"fmt"
)

func main() {
	book := app.Book{
		Title:  "The Hobbit",
		Author: "J.R.R. Tolkien",
		Pages:  310,
	}
	book.Summary()

	// Reviewing pointers
	// When modified directly the struct changes as expected
	fmt.Println("----------------------------")
	fmt.Println("Mutating the Struct Directly")
	fmt.Println("----------------------------")
	book.Author = "Santiago Cano"
	book.Summary()

	// However, if a function needs to directly modify the object...
	fmt.Println("----------------------------")
	fmt.Println("Mutating the Struct from a Function Reference")
	fmt.Println("----------------------------")

	book.UpdatePages(10)
	book.Summary()

	fmt.Println("----------------------------")
	fmt.Println("Mutating the Struct from a Function *Pointer Reference")
	fmt.Println("----------------------------")
	book.UpdatePagesPointer(30)
	book.Summary()
}
