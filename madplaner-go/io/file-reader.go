package io

import (
	"io/ioutil"
	"os"
	"strings"

	"../log"
	"../validator"
)

func GetFileByteArray(dir string, fn string) []byte {
	filepath := dir + "/" + fn
	file, err := os.Open(filepath)
	validator.ValidateError(err, "Cannot open: "+filepath)

	s, err := ioutil.ReadAll(file)
	validator.ValidateError(err, "Cannot readall: "+filepath)

	return s
}

func GetFilenamesInDirectoryWithSuffix(dir string, suffix string) []string {
	log.Start("file-reader", "GetFilenamesInDirectoryWithSuffix")

	fileInfos, err := ioutil.ReadDir(dir)
	validator.ValidateError(err, "Cannot read directory: "+dir)

	mpFilenames := getFilenamesWithSuffix(fileInfos, suffix)

	log.End("file-reader", "GetFilenamesInDirectoryWithSuffix")
	return mpFilenames
}

func getFilenamesWithSuffix(fileinfos []os.FileInfo, suffix string) []string {
	fns := []string{}

	for _, file := range fileinfos {
		if !file.IsDir() && strings.HasSuffix(file.Name(), suffix) {
			fns = append(fns, file.Name())
		}
	}

	return fns
}
