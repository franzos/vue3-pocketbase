package main

import (
	"github.com/asaskevich/govalidator"
    "github.com/microcosm-cc/bluemonday"
)

func sanitizeInput(input string) string {
    p := bluemonday.UGCPolicy()
    return p.Sanitize(input)
}

func validateInput(input string) bool {
    return govalidator.IsPrintableASCII(input) // You can add more validation rules here
}