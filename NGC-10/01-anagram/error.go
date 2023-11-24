package main

import "fmt"

type ErrContract struct {
	Type    string
	Message string
}

func (e ErrContract) Error() string {
	return fmt.Sprintf("%v: %v\n", e.Type, e.Message)
}

func errValidation(msg string) *ErrContract {
	return &ErrContract{
		Type:    "invalid input",
		Message: msg,
	}
}
