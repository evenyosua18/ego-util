package errors

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type Error struct {
	CustomCode      int    `yaml:"code"`
	ResponseMessage string `yaml:"message"`
	ErrorMessage    string `yaml:"errors"`
	ResponseCode    int    `yaml:"response_code"`
}

func (e Error) Error() error {
	return errors.New(e.ErrorMessage)
}

func (e Error) Message() string {
	return e.ResponseMessage
}

func (e Error) Code() int {
	return e.CustomCode
}

func (e Error) CodeResponse() int {
	return e.ResponseCode
}

func Extract(e error) (err Error) {
	//split errors message
	if e != nil && e.Error() != "" {
		msg := e.Error()

		if !(strings.Index(msg, ":") == -1 || len(msg[:strings.Index(msg, ":")]) != 3) {
			//if the 3 letter before ":" is not digit
			if code, errConversion := strconv.Atoi(msg[:strings.Index(msg, ":")]); errConversion == nil {
				err.CustomCode = code
			}

			//set errors message
			err.ErrorMessage = msg[strings.Index(msg, ":")+2:]
		}
	}

	//check from map
	if CustomError[err.CustomCode].CustomCode == 0 {
		err.CustomCode = 999
	}

	//if errors message is empty, use from custom code
	if err.ErrorMessage == "" {
		err.ErrorMessage = CustomError[err.CustomCode].ErrorMessage
	}

	//set response message
	err.ResponseMessage = CustomError[err.CustomCode].ResponseMessage

	//set response code
	err.ResponseCode = CustomError[err.CustomCode].ResponseCode

	return
}

func Wrap(err error, code ...int) error {
	if len(code) == 1 {
		return fmt.Errorf("%d: %w", code[0], err)
	} else {
		return err
	}
}

func Get(code int) (err Error) {
	err = CustomError[code]

	if err.CustomCode == 0 {
		err = CustomError[999]
	}

	return
}
