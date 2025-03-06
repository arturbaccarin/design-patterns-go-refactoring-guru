package main

import "fmt"

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
