package main

import (
	"fmt"

	"golang.org/x/example/hello/reverse"
)

func main() {
	albums, err := AlbumByArtist("bsb")
	if err != nil {
		fmt.Println(albums)
	}
	fmt.Println(albums)
	fmt.Println(reverse.String("Hello"), reverse.Int(24601))
}
