package app

import "fmt"

type Book struct {
	Title  string
	Author string
	Pages  int
}

func (b Book) Summary() {
	fmt.Printf("%s by %s has %d pages\n", b.Title, b.Author, b.Pages)
}

func (b Book) UpdatePages(p int) {
	b.Pages = p
	fmt.Printf("Updated book: %v\n", b)
}

func (b *Book) UpdatePagesPointer(p int) {
	b.Pages = p
	fmt.Printf("Updated with Pointer Book: %v\n", b)
}
