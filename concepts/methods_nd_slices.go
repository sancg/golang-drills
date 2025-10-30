package concepts

import (
	"book-tool-cli/internal/app"
	"fmt"
)

func RunMethodsAndSlices(b *app.Book) {
	fmt.Println("----------------------------")
	fmt.Println("Drill Exercise: Slices & Methods")
	fmt.Println("----------------------------")
	lib := app.Library{}
	lib.AddCollection("Fantasy")

	fmt.Printf("Collections: %v\n", lib)
	lib.AddCollection("Fantasy")
	fmt.Printf("Collections After duplicate?: %v\n", lib)

	lib.AddBook(*b, "")
	fmt.Printf("Collection Default: %v\n", lib)

	lib.AddBook(*b, "Fantasy")
	fmt.Printf("Collection Fantasy: %v\n", lib)

	lib.AddBook(*b, "")
	fmt.Printf("Collection Default Duplicate?: %v\n", lib)

	lib.AddBook(app.Book{Author: "Edgar Schein", Title: "Humble Inquiry: The Gentle Art of Asking Instead of Telling", Pages: 144}, "Daily")
	lib.ListBook()
}
