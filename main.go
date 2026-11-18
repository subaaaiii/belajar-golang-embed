package main

import (
	"embed"
	_ "embed"
	"fmt"
	"os"
)

//go:embed version.txt
var version string

//go:embed version.txt
var version2 string

//go:embed foto.jpeg
var foto []byte

//go:embed files/*.txt
var path embed.FS

/*
saat di build semua informasi akan disimpan dalam binary go, bukan dibaca secara real time
jadi walaupun file hilang atau dihapus tetap tersimpan dalam variable yang telah ditentukan
*/
func main() {
	fmt.Println(version)
	fmt.Println(version2)
	err := os.WriteFile("foto_new.png", foto, 0644)
	if err != nil {
		panic(err)
	}
	dirEntries, _ := path.ReadDir("files")
	for _, entry := range dirEntries {
		if !entry.IsDir() {
			fmt.Println(entry.Name())
			result, _ := path.ReadFile("files/" + entry.Name())
			fmt.Println(string(result))
		}
	}
}
