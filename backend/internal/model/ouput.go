package model

import "errors"

// Output collects error and/or response, and print them by specified mode.
type Output struct {
	Error   error  `json:"-"`
	Message string `json:"message"`
	Status  int    `json:"-"`
}

//ErrorNil creates empty Output
func ErrorNil() Output {
	return Output{}
}

//Create creates Output object
func ErrorString(msg string, status int) Output {
	return Output{
		Error:   errors.New(msg),
		Message: msg,
		Status:  status,
	}
}

func (out Output) IsError() bool {
	return out.Error != nil
}
