package main

import "fmt"

/*
https://refactoring.guru/design-patterns/iterator

Iterator is a behavioral design pattern that lets you traverse elements of
a collection without exposing its underlying representation (list, stack, tree, etc.).

Collections are one of the most used data types in programming. Nonetheless,
a collection is just a container for a group of objects.

But no matter how a collection is structured, it must provide some way of accessing its
elements so that other code can use these elements. There should be a way to go through
each element of the collection without accessing the same elements over and over.

Adding more and more traversal algorithms to the collection gradually blurs its primary responsibility,
which is efficient data storage.

On the other hand, the client code that’s supposed to work with various collections may not even care
how they store their elements. However, since collections all provide different ways of accessing
their elements, you have no option other than to couple your code to the specific collection classes.

The main idea of the Iterator pattern is to extract the traversal behavior of a collection
into a separate object called an iterator.

In addition to implementing the algorithm itself, an iterator object encapsulates all of the traversal details,
such as the current position and how many elements are left till the end.
Because of this, several iterators can go through the same collection at the same time, independently of each other.

All iterators must implement the same interface.

If you need a special way to traverse a collection, you just create a new iterator class,
without having to change the collection or the client.

Iterator is a behavioral design pattern that allows sequential traversal through a
complex data structure without exposing its internal details.
*/

// Collection
type Collection interface {
	createIterator() Iterator
}

// Concrete Collection
type UserCollection struct {
	users []*User
}

func (u *UserCollection) createIterator() Iterator {
	return &UserIterator{
		users: u.users,
	}
}

// Iterator
type Iterator interface {
	hasNext() bool
	getNext() *User
}

// Concrete Iterator
type UserIterator struct {
	index int
	users []*User
}

func (u *UserIterator) hasNext() bool {
	if u.index < len(u.users) {
		return true
	}
	return false

}
func (u *UserIterator) getNext() *User {
	if u.hasNext() {
		user := u.users[u.index]
		u.index++
		return user
	}
	return nil
}

type User struct {
	name string
	age  int
}

func main() {

	user1 := &User{
		name: "a",
		age:  30,
	}
	user2 := &User{
		name: "b",
		age:  20,
	}

	userCollection := &UserCollection{
		users: []*User{user1, user2},
	}

	iterator := userCollection.createIterator()

	for iterator.hasNext() {
		user := iterator.getNext()
		fmt.Printf("User is %+v\n", user)
	}
}
