package utils

import (
	"image"
	"image/jpeg"
	"image/png"
	"os"

	"github.com/jdeng/goheif"
	"golang.org/x/image/bmp"
)

const BMP = "bmp"
const HEIC = "HEIC"

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
	default:
		ThrowFatal()
	}

	return nil
}

func EncodeImage(file *os.File, image image.Image, srcFormat string) {
	switch srcFormat {
	case BMP:
		err := png.Encode(file, image)
		HandleError(err)
	case HEIC:
		err := jpeg.Encode(file, image, nil)
		HandleError(err)
	default:
		ThrowFatal()
	}
}
