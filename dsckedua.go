package main

import "fmt"

type Printable interface {
	printPrimary()
	printSecondary(message string)
}

type Buku struct {
	judul    string
	penulis  string
	penerbit string
}

func (b Buku) printPrimary() {
	fmt.Println(b.judul)
}

func (b Buku) printSecondary(message string) {
	fmt.Println(message)
	fmt.Println(b.judul)
}

func main() {
	var printable Printable
	printable = Buku{"Perahu Kertas", "Dewi Lestari", "Bentang"}
	printable.printPrimary()
	printable.printSecondary("Ini adalah judul")
}
