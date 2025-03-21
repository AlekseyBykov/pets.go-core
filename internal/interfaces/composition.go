package interfaces

// Printable = Printer + Validator
type Printable interface {
	Printer
	Validator
}

type Validator interface {
	IsValid() bool
}

func (b Book) IsValid() bool {
	return b.Title != "" && b.Author != ""
}

func (d Device) IsValid() bool {
	return d.Model != ""
}
