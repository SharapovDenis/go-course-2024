package main

import (
	"fmt"
	"homework/validator"
)

type MyStruct struct {
	InA     string `validate:"in:ab,cd"`
	InB     string `validate:"in:aa,bb,cd,ee"`
	InC     int    `validate:"in:-1,-3,5,7"`
	InD     int    `validate:"in:5-"`
	InEmpty string `validate:"in:"`
}

func main() {

	ms := MyStruct{
		InA:     "ef",
		InB:     "ab",
		InC:     2,
		InD:     12,
		InEmpty: "",
	}
	err := validator.Validate(ms)
	if err != nil {
		fmt.Print(err)
	}

}
