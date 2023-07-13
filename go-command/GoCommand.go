package main

import (
	"errors"
	"github.com/aws/smithy-go"
	"log"
)

func main() {
	var ae smithy.APIError
	if errors.As(nil, &ae) {
		log.Printf("code: %s, message: %s, fault: %s", ae.ErrorCode(), ae.ErrorMessage(), ae.ErrorFault().String())
	}
}
