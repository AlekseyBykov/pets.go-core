package interfaces

import "testing"

func TestPrinterInterface(t *testing.T) {
	b := Book{Title: "Go in Action", Author: "William Kennedy"}
	d := Device{Model: "MacBook Pro"}

	var p Printer

	p = b
	if p.PrintInfo() != "Book: Go in Action by William Kennedy" {
		t.Errorf("unexpected print info for Book")
	}

	p = d
	if p.PrintInfo() != "Device model: MacBook Pro" {
		t.Errorf("unexpected print info for Device")
	}
}

func TestPrintableComposition(t *testing.T) {
	b := Book{Title: "Go in Action", Author: "William Kennedy"}
	d := Device{Model: ""}

	var pb Printable

	pb = b
	if !pb.IsValid() {
		t.Errorf("expected valid printable Book")
	}

	pb = d
	if pb.IsValid() {
		t.Errorf("expected invalid printable Device")
	}
}
