package main

import (
	"bufio"
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/jdeng/goheif"
	"golang.org/x/image/bmp"
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

	srcDir = normalizeFolderName(srcDir)
	convDir = normalizeFolderName(convDir)

	fmt.Println("Choose source format:")
	fmt.Println("1 - BMP")
	fmt.Println("2 - HEIC")
	scanner.Scan()
	srcFormat = scanner.Text()

	srcFormat = assignFormat(srcFormat)

	files, err := ioutil.ReadDir(srcDir)
	handleError(err)

	fmt.Printf("Found %d files\n", len(files))

	var counter = 0

	for _, f := range files {
		if !f.IsDir() && strings.HasSuffix(f.Name(), srcFormat) {
			counter++

			fmt.Printf("File #%d\n", counter)
			fmt.Printf("\tOriginal file name: %v\n", f.Name())
			fmt.Printf("\tWill be using this original date of modification: %v\n", f.ModTime())

			oldFile, err := os.Open(srcDir + f.Name())
			handleError(err)

			img := decodeImage(oldFile, srcFormat)

			newName := setNewFileName(f.Name(), srcFormat)

			newFile, err := os.Create(convDir + newName)
			handleError(err)

			encodeImage(newFile, img, srcFormat)

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

func normalizeFolderName(name string) string {
	if !strings.HasSuffix(name, "/") {
		name += "/"
	}

	return name
}

func assignFormat(format string) string {
	switch format {
	case "1":
		return "bmp"
	case "2":
		return "HEIC"
	default:
		log.Fatal("Incorrect value received")
		os.Exit(100)
	}

	return ""
}

func decodeImage(srcFile *os.File, srcFormat string) image.Image {
	switch srcFormat {
	case "bmp":
		img, err := bmp.Decode(srcFile)
		handleError(err)

		return img
	case "HEIC":
		img, err := goheif.Decode(srcFile)
		handleError(err)
		return img
	default:
		log.Fatal("Incorrect value received")
		os.Exit(100)
	}

	return nil
}

func setNewFileName(name string, format string) string {
	switch format {
	case "bmp":
		return strings.Replace(name, format, "png", -1)
	case "HEIC":
		return strings.Replace(name, format, "jpeg", -1)
	default:
		log.Fatal("Incorrect value received")
		os.Exit(100)
	}

	return ""
}

func encodeImage(file *os.File, image image.Image, srcFormat string) {
	switch srcFormat {
	case "bmp":
		err := png.Encode(file, image)
		handleError(err)
	case "HEIC":
		err := jpeg.Encode(file, image, nil)
		handleError(err)
	default:
		log.Fatal("Incorrect value received")
		os.Exit(100)
	}
}
