package validator

import (
	"fmt"
	"os"
)

func ValidateError(err error, msg string) {
	if err != nil {
		fmt.Println("Exits program due to error", msg, err)
		os.Exit(1)
	}
}
