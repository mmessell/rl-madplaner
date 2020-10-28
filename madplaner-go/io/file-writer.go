package io

import (
	"fmt"
	"os"

	"../validator"
)

func CreateFile(dir string, fn string, text string) {
	fmt.Println("--- file-writer.go", "start CreateFile ---")

	ensureDirectory(dir)
	file := createFile(dir, fn)
	file.WriteString(text)
	file.Close()

	fmt.Println("--- file-writer.go", "end CreateFile ---")
}

func ensureDirectory(dir string) {
	err := os.MkdirAll(dir, os.ModePerm)
	validator.ValidateError(err, "Error while ensuring directory: "+dir)
}

func createFile(dir string, fn string) *os.File {
	filepath := dir + "/" + fn
	file, err := os.Create(filepath)
	validator.ValidateError(err, "Error while creating file: "+filepath)
	return file
}
