package models

import (
	"encoding/json"

	"../io"
	"../log"
	. "../types"
)

type MadplanFile struct {
	Name       string
	Persons    Double
	Weeks      []Double
	Dagsplaner []DagsplanFile
}

func getMadplanFile(dir string, fn string) MadplanFile {
	log.Start("madplan-file", "GetMadplanFile")

	b := io.GetFileByteArray(dir, fn)
	mpFile := bytesToMadplanFile(b)

	log.End("madplan-file", "GetMadplanFile")
	return mpFile
}

func bytesToMadplanFile(b []byte) MadplanFile {
	var mp = MadplanFile{}
	json.Unmarshal(b, &mp)
	return mp
}
