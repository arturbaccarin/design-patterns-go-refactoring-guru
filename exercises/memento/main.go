package main

import "fmt"

/*
Imagine you’re developing a text editor that needs
an undo feature to revert changes. To achieve this,
implement the Memento design pattern, which captures
and restores an object’s state without exposing its
internal structure. Create a TextEditor class to manage
text input, a Memento to store text states, and a
Caretaker to handle the undo functionality by saving
and restoring text states. Ensure the editor can save
its state, modify the text, and undo changes, printing
the text after each operation.
*/

type Memento struct {
	text string
}

type TextEditor struct {
	text string
}

func (e *TextEditor) SetText(text string) {
	e.text = text
}

func (e *TextEditor) GetText() string {
	return e.text
}

func (e *TextEditor) CreateMemento() *Memento {
	return &Memento{text: e.text}
}

func (e *TextEditor) Restore(m Memento) {
	e.text = m.text
}

type Caretaker struct {
	undoStack [](*Memento)
	editor    *TextEditor
}

func (c *Caretaker) Save() {
	c.undoStack = append(c.undoStack, c.editor.CreateMemento())
}

func (c *Caretaker) Undo() {
	if len(c.undoStack) == 0 {
		fmt.Println("Nothing to undo")
		return
	}

	memento := c.undoStack[len(c.undoStack)-1]
	c.editor.Restore(*memento)
	c.undoStack = c.undoStack[:len(c.undoStack)-1]
}

func main() {
	editor := &TextEditor{}
	caretaker := &Caretaker{editor: editor}

	// Initial text
	editor.SetText("Hello")
	caretaker.Save() // Save state after text entry.
	fmt.Println("Initial Text:", editor.GetText())

	// Modify text
	editor.SetText("Hello, World!")
	fmt.Println("Text after edit:", editor.GetText())

	caretaker.Undo()
	fmt.Println("Text after undo:", editor.GetText())
}
