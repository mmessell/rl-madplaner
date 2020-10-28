package main

import (
	"os"

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

	usrDir, _ := os.UserHomeDir()
	baseDir := usrDir + "/Dropbox/Ren lykke/Madplaner"
	configDif := baseDir + "/configuration"

	latexDir := configDif + "/latex/texfiles"
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
