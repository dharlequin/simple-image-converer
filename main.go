package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"dharlequin/go-image-converter/utils"
)

func main() {
	var srcDir string
	var convDir string

	var srcFormat string

	fmt.Println("Enter source directory")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	srcDir = scanner.Text()

	fmt.Println("Enter convert destination directory")
	scanner.Scan()
	convDir = scanner.Text()

	srcDir = utils.NormalizeFolderName(srcDir)
	convDir = utils.NormalizeFolderName(convDir)

	fmt.Println("Choose source format:")
	fmt.Println("1 - BMP")
	fmt.Println("2 - HEIC")
	scanner.Scan()
	srcFormat = scanner.Text()

	srcFormat = utils.AssignFormat(srcFormat)

	files, err := ioutil.ReadDir(srcDir)
	utils.HandleError(err)

	fmt.Printf("Found %d files\n", len(files))

	var counter = 0

	for _, f := range files {
		if !f.IsDir() && strings.HasSuffix(f.Name(), srcFormat) {
			counter++

			fmt.Printf("File #%d\n", counter)
			fmt.Printf("\tOriginal file name: %v\n", f.Name())
			fmt.Printf("\tWill be using this original date of modification: %v\n", f.ModTime())

			oldFile, err := os.Open(srcDir + f.Name())
			utils.HandleError(err)

			img := utils.DecodeImage(oldFile, srcFormat)

			newName := utils.SetNewFileName(f.Name(), srcFormat)

			newFile, err := os.Create(convDir + newName)
			utils.HandleError(err)

			utils.EncodeImage(newFile, img, srcFormat)

			os.Chtimes(convDir+newName, f.ModTime(), f.ModTime())
		}
	}

	fmt.Printf("Finished converting %d files\n\n", counter)
}
