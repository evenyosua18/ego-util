package codes

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type Code struct {
	CustomCode      int    `yaml:"code"`
	ResponseMessage string `yaml:"message"`
	ErrorMessage    string `yaml:"codes"`
	ResponseCode    int    `yaml:"response_code"`
}

func (e Code) Error() error {
	return errors.New(e.ErrorMessage)
}

func (e Code) Message() string {
	return e.ResponseMessage
}

func (e Code) Code() int {
	return e.CustomCode
}

func (e Code) CodeResponse() int {
	return e.ResponseCode
}

func Extract(e error) (code Code) {
	//split codes message
	if e != nil && e.Error() != "" {
		msg := e.Error()

		if !(strings.Index(msg, ":") == -1 || len(msg[:strings.Index(msg, ":")]) != 3) {
			//if the 3 letter before ":" is not digit
			if c, errConversion := strconv.Atoi(msg[:strings.Index(msg, ":")]); errConversion == nil {
				code.CustomCode = c
			}

			//set codes message
			code.ErrorMessage = msg[strings.Index(msg, ":")+2:]
		}
	}

	//check from map
	if customCodes[code.CustomCode].CustomCode == 0 {
		code.CustomCode = defaultUnknownCode
	}

	//if codes message is empty, use from custom code
	if code.ErrorMessage == "" {
		code.ErrorMessage = customCodes[code.CustomCode].ErrorMessage
	}

	//set response message
	code.ResponseMessage = customCodes[code.CustomCode].ResponseMessage

	//set response code
	code.ResponseCode = customCodes[code.CustomCode].ResponseCode

	return
}

func Wrap(err error, code ...int) error {
	if len(code) == 1 {
		return fmt.Errorf("%d: %w", code[0], err)
	} else {
		return err
	}
}

func Create(code int) (err Code) {
	//check from map
	if customCodes[code].CustomCode == 0 {
		return customCodes[code]
	} else {
		return customCodes[defaultUnknownCode]
	}
}

func FromError(err error, codes ...int) (code Code) {
	//check err should be not empty
	if err == nil {
		return customCodes[defaultUnknownCode]
	}

	//set up error message
	code.ErrorMessage = err.Error()

	if len(codes) == 1 && customCodes[codes[0]].CustomCode != 0 {
		code.ResponseCode = customCodes[codes[0]].ResponseCode
		code.ResponseMessage = customCodes[codes[0]].ResponseMessage
		code.CustomCode = codes[0]
	}

	return
}

func Get(code int) (err Code) {
	err = customCodes[code]

	if err.CustomCode == 0 {
		err = customCodes[defaultUnknownCode]
	}

	return
}
