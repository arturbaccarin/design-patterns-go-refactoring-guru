package main

import "fmt"

/*
In this exercise, we'll implement the Iterator pattern for a collection of Books.
We’ll create a BookCollection that holds a list of books and an iterator
to traverse over them.
*/

type Iterator interface {
	hasNext() bool
	getNext() *Book
}

type Book struct {
	title  string
	author string
}

type BookCollection struct {
	books []*Book
	index int
}

func (b *BookCollection) hasNext() bool {
	if b.index < len(b.books) {
		return true
	}
	return false
}

func (b *BookCollection) getNext() *Book {
	if b.hasNext() {
		book := b.books[b.index]
		b.index++
		return book
	}
	return nil
}

func main() {
	books := []*Book{
		{title: "The Great Gatsby", author: "F. Scott Fitzgerald"},
		{title: "To Kill a Mockingbird", author: "Harper Lee"},
		{title: "Pride and Prejudice", author: "Jane Austen"},
	}

	collection := &BookCollection{books: books}

	var iterator Iterator
	iterator = collection

	for iterator.hasNext() {
		book := iterator.getNext()
		fmt.Printf("Title: %s, Author: %s\n", book.title, book.author)
	}
}
