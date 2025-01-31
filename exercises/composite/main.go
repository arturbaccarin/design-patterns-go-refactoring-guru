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

type Folder struct {
	children []Component
}

func (f *Folder) Display() {
	for _, child := range f.children {
		child.Display()
	}
}

func (f *Folder) Add(child Component) {
	f.children = append(f.children, child)
}

func main() {
	folder1 := &Folder{}

	file1 := &File{"file1"}
	file2 := &File{"file2"}
	file3 := &File{"file3"}
	folder1.Add(file1)
	folder1.Add(file2)
	folder1.Add(file3)

	folder2 := &Folder{}

	file4 := &File{"file4"}
	file5 := &File{"file5"}

	folder2.Add(file4)
	folder2.Add(file5)

	folder1.Add(folder2)

	folder1.Display()
}
