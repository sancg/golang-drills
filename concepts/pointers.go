package concepts

import (
	"book-tool-cli/internal/app"
	"fmt"
)

func RunBookPointers(b *app.Book) {
	// Reviewing pointers
	// When modified directly the struct changes as expected
	fmt.Println("----------------------------")
	fmt.Println("Mutating the Struct AUTHOR Directly")
	fmt.Println("----------------------------")
	b.Author = "Santiago Cano"
	b.Summary()

	// However, if a function needs to directly modify the object...
	fmt.Println("----------------------------")
	fmt.Println("Mutating the Struct PAGES from a Function Reference")
	fmt.Println("----------------------------")

	b.UpdatePages(10)
	b.Summary()

	fmt.Println("----------------------------")
	fmt.Println("Mutating the Struct PAGES from a Function *Pointer Reference")
	fmt.Println("----------------------------")
	b.UpdatePagesPointer(30)
	b.Summary()
}

func RunOperatedWithPointers() {
	fmt.Println("----------------------------")
	fmt.Println("Operating with Pointer Reference inside a Function")
	fmt.Println("----------------------------")

	age := 9
	var agePointer = &age
	fmt.Println("My address AGE is: ", agePointer)
	fmt.Println("My pointer AGE is: ", *agePointer)
	fmt.Println()

	fmt.Println("am I adult?: ", isAdult(agePointer))
}

func isAdult(a *int) bool {
	return *a-18 >= 0
}
