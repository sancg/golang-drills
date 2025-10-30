package main

import (
	"book-tool-cli/concepts"
	"book-tool-cli/internal/app"
)

func main() {
	book := app.Book{
		Title:  "The Hobbit",
		Author: "J.R.R. Tolkien",
		Pages:  310,
	}
	book.Summary()
	// concepts.RunBookPointers(&book)
	// concepts.RunOperatedWithPointers()
	concepts.RunMethodsAndSlices(&book)

}
