package main

import (
	"bufio"
	"fmt"
	"os"
)

type Buku struct {
	judul    string
	penulis  string
	penerbit string
}

func main() {
	var buku Buku
	var scanner = bufio.NewReader(os.Stdin)
	fmt.Println("Masukkan judul buku")
	buku.judul, _ = scanner.ReadString('\n')

	fmt.Println("Masukkan penulis buku")
	buku.penulis, _ = scanner.ReadString('\n')

	fmt.Println("Masukkan penerbit buku")
	buku.penerbit, _ = scanner.ReadString('\n')

	fmt.Println("Judul:", buku.judul, "\nPenulis:", buku.penulis, "\nPenerbit:", buku.penerbit)
}
