package main

import (
	"fmt"

	"./io"
	"./latex"
	"./log"
	"./models"
)

func main() {
	log.Start("program", "main")

	for _, filename := range io.GetFilenamesInDirectoryWithSuffix(Config.MadplanDir, ".json") {
		fmt.Println("\nProcessing: " + filename)
		mp := models.GetMadplan(filename, Config.MadplanDir, Config.OpskrifterDir)
		latex.CreateLatexFile(mp, Config.LatexDir, Config.OpskrifterImageDir)
	}

	log.End("program", "main")
}
