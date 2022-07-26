package main

import (
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

	srcDir = utils.GetFolderName("Enter source directory")
	convDir = utils.GetFolderName("Enter convert destination directory")

	srcFormat = utils.GetSourceFormat()

	files, err := ioutil.ReadDir(srcDir)
	utils.HandleError(err)

	fmt.Printf("Found %d total files\n", len(files))

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
