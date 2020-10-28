package models

import (
	"encoding/json"
	"fmt"

	"../log"
)

type myError struct {
	err     error
	message string
	exit    bool
}

func (myErr myError) ToString() string {

	out, err := json.Marshal(myErr)

	if err != nil {
		log.Fatal(err, "A myError model couldn't be parse into bytes")
	}

	fmt.Println(string(out))

	return string(out)
}
