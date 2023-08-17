package utils

import (
	"image"
	"image/jpeg"
	"image/png"
	"log"
	"os"

	"github.com/jdeng/goheif"
	"golang.org/x/image/bmp"
)

const BMP = "bmp"
const HEIC = "HEIC"
const PNG = "png"

func DecodeImage(srcFile *os.File, srcFormat string) image.Image {
	switch srcFormat {
	case BMP:
		img, err := bmp.Decode(srcFile)
		HandleError(err)
		return img
	case HEIC:
		img, err := goheif.Decode(srcFile)
		HandleError(err)
		return img
	case PNG:
		img, err := png.Decode(srcFile)
		HandleError(err)
		return img
	default:
		log.Fatalln("Incorrect value received")
	}

	return nil
}

func EncodeImage(file *os.File, image image.Image, srcFormat string) {
	switch srcFormat {
	case BMP:
		err := png.Encode(file, image)
		HandleError(err)
	case HEIC, PNG:
		err := jpeg.Encode(file, image, nil)
		HandleError(err)
	default:
		log.Fatalln("Incorrect value received")
	}
}
