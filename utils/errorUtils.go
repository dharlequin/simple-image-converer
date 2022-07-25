package utils

import (
	"log"
	"os"
)

func HandleError(err error) {
	if err != nil {
		ThrowFatal()
	}
}

func ThrowFatal() {
	log.Fatal("Incorrect value received")
	os.Exit(100)
}
