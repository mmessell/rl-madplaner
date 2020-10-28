package main

import (
	"runtime"

	"path/filepath"
	"./log"
)

type configuration struct {
	LatexDir           string
	MadplanDir         string
	MadplanImageDir    string
	OpskrifterDir      string
	OpskrifterImageDir string
}

var Config configuration

func init() {
	log.Start("config", "init")

	_, currentFile, _, _ := runtime.Caller(0)
	relPath := filepath.Dir(currentFile)

	outputDIR := "/../output"
	latexDir := relPath + outputDIR + "/latexfiles"

	configDif := relPath + "/../configuration"
	madplanDir := configDif + "/madplaner"
	madplanImageDir := madplanDir + "/billeder"
	opskrifterDir := configDif + "/opskrifter"
	opskrifterImageDir := opskrifterDir + "/billeder"

	Config = configuration{
		LatexDir:           latexDir,
		MadplanDir:         madplanDir,
		MadplanImageDir:    madplanImageDir,
		OpskrifterDir:      opskrifterDir,
		OpskrifterImageDir: opskrifterImageDir,
	}

	log.End("config", "init")
}
