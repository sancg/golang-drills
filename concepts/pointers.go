package concepts

import (
	"book-tool-cli/internal/app"
	"fmt"
)

func RunPointers(b *app.Book) {
	// Reviewing pointers
	// When modified directly the struct changes as expected
	fmt.Println("----------------------------")
	fmt.Println("Mutating the Struct Directly")
	fmt.Println("----------------------------")
	b.Author = "Santiago Cano"
	b.Summary()

	// However, if a function needs to directly modify the object...
	fmt.Println("----------------------------")
	fmt.Println("Mutating the Struct from a Function Reference")
	fmt.Println("----------------------------")

	b.UpdatePages(10)
	b.Summary()

	fmt.Println("----------------------------")
	fmt.Println("Mutating the Struct from a Function *Pointer Reference")
	fmt.Println("----------------------------")
	b.UpdatePagesPointer(30)
	b.Summary()
}
