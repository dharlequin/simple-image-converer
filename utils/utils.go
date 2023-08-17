package utils

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func GetFolderName(instruction string) string {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println(instruction)
	scanner.Scan()
	directory := scanner.Text()

	validateFolderName(directory)
	return normalizeFolderName(directory)
}

func normalizeFolderName(name string) string {
	if !strings.HasSuffix(name, "/") {
		name += "/"
	}

	return name
}

func validateFolderName(name string) {
	if name == "" {
		log.Fatalln("Value cannot be empty")
	}
}

func GetSourceFormat() string {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Choose source format by entering corresponding number:")
	fmt.Println("1 - BMP to PNG")
	fmt.Println("2 - HEIC to JPEG")
	fmt.Println("3 - PNG to JPG")
	scanner.Scan()
	srcFormat := scanner.Text()

	return assignFormat(srcFormat)
}

func assignFormat(format string) string {
	switch format {
	case "1":
		return BMP
	case "2":
		return HEIC
	case "3":
		return PNG
	default:
		log.Fatalln("Incorrect value received")
	}

	return ""
}

func SetNewFileName(name string, format string) string {
	switch format {
	case BMP:
		return strings.Replace(name, format, "png", -1)
	case HEIC, PNG:
		return strings.Replace(name, format, "jpg", -1)
	default:
		log.Fatalln("Incorrect value received")
	}

	return ""
}
