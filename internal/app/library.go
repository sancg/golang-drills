package app

import "fmt"

type Library struct {
	Shelf []Collection
}

type Collection struct {
	Name      string
	BookShelf []Book
}

func (l *Library) AddCollection(n string) {
	cl := Collection{Name: n}
	cl.NewCollection(n)

	if cl.isCollection(l) {
		fmt.Println("-- WARN: Collection already exists --")
		return
	}

	l.Shelf = append(l.Shelf, cl)
}

func (l *Library) AddBook(b Book, n string) {
	var cl Collection

	if n != "" {
		cl.Name = n

		fmt.Printf("\nAdding book to %s Shelf\n", n)
		if cl.isCollection(l) {
			l.updateShelf(&cl, b)
			fmt.Printf("-- Library: %v\n", l)
			return
		} else {
			fmt.Println("-- WARN: Fallback for adding book to Default Shelf")
		}
	}

	fmt.Println("\nAppending book to default Shelf")
	cl.NewCollection("default")
	cl.BookShelf = append(cl.BookShelf, b)

	// Check if the collection was already added
	if !cl.isCollection(l) {
		l.Shelf = append(l.Shelf, cl)
	} else {
		fmt.Println("Default Shelf already in the Library")
		l.updateShelf(&cl, b)
	}
}

func (l *Library) updateShelf(c *Collection, b Book) {
	for i, v := range l.Shelf {
		if v.Name == c.Name {
			// TODO: Avoid duplicate entries
			v.BookShelf = append(v.BookShelf, b)
			l.Shelf[i] = v
			fmt.Printf("Modified Bookshelf %v\n", v)
			break
		}
	}
}

func (l *Library) ListBook() {
	for _, s := range l.Shelf {
		fmt.Println(s.Name)
		for _, b := range s.BookShelf {
			fmt.Println(b)
		}
		fmt.Println("--------------------")
	}
}

func (c *Collection) NewCollection(n string) {
	c.Name = n
}

func (c *Collection) isCollection(l *Library) bool {
	for _, v := range l.Shelf {
		if v.Name == c.Name {
			return true
		}
	}
	return false
}
