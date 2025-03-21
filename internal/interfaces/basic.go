package interfaces

import "fmt"

type Printer interface {
	PrintInfo() string
}

type Book struct {
	Title  string
	Author string
}

func (b Book) PrintInfo() string {
	return fmt.Sprintf("Book: %s by %s", b.Title, b.Author)
}

type Device struct {
	Model string
}

func (d Device) PrintInfo() string {
	return fmt.Sprintf("Device model: %s", d.Model)
}
