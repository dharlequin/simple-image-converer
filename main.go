package main

import (
	"bufio"
	"fmt"
	"image/png"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"golang.org/x/image/bmp"
)

func main() {
	var srcDir string
	var convDir string

	fmt.Println("Enter source directory")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	srcDir = scanner.Text()

	fmt.Println("Enter convert destination directory")
	scanner.Scan()
	convDir = scanner.Text()

	if !strings.HasSuffix(srcDir, "/") {
		srcDir += "/"
	}

	if !strings.HasSuffix(convDir, "/") {
		convDir += "/"
	}

	files, err := ioutil.ReadDir(srcDir)
	handleError(err)

	fmt.Printf("Found %d files\n", len(files))

	var counter = 0

	for _, f := range files {
		if !f.IsDir() {
			counter++

			fmt.Printf("File #%d\n", counter)
			fmt.Printf("\tOriginal file name: %v\n", f.Name())
			fmt.Printf("\tWill be using this original date of modification: %v\n", f.ModTime())

			oldFile, err := os.Open(srcDir + f.Name())
			handleError(err)

			img, err := bmp.Decode(oldFile)
			handleError(err)

			newName := strings.Replace(f.Name(), "bmp", "png", -1)
			// fmt.Println(newName)

			newFile, err := os.Create(convDir + newName)
			handleError(err)

			err = png.Encode(newFile, img)
			handleError(err)

			os.Chtimes(convDir+newName, f.ModTime(), f.ModTime())
		}
	}

	fmt.Printf("Finished converting %d files\n\n", counter)
}

func handleError(err error) {
	if err != nil {
		log.Fatal(err)
		os.Exit(100)
	}
}
