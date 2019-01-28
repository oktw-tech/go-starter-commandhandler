package crosscutting

import (
	"fmt"
	"log"
)

//RaiseError : Raise the error with panic
func RaiseError(message string, err error) {
	if err != nil {
		log.Fatalf("%s: %s", message, err)
		panic(fmt.Sprintf("%s: %s", message, err))
	}
}
