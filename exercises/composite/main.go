package main

/*
Let’s imagine a simple file system where we have File and Directory as
components of the system. A Directory can contain both File objects and
other Directory objects.
*/

type Component interface {
	Display()
}

type File struct {
	name string
}

func (f *File) Display() {
	println(f.name)
}

type Composite struct {
	children []Component
}

func (c *Composite) Display() {
	for _, child := range c.children {
		child.Display()
	}
}

func (c *Composite) Add(child Component) {
	c.children = append(c.children, child)
}

func main() {

}
