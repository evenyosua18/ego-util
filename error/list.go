package error

import (
	"gopkg.in/yaml.v3"
	"log"
	"os"
)

var (
	CustomError map[int]Error
)

func init() {
	CustomError = map[int]Error{
		999: {
			CustomCode:      999,
			ResponseMessage: "need to register your custom code",
			ErrorMessage:    "unknown error code",
			ResponseCode:    500,
		},
	}
}

func RegisterError(path string) {
	//read file
	f, err := os.ReadFile(path)

	if err != nil {
		panic(err)
	}

	e := struct {
		Errors []Error `yaml:"errors"`
	}{}

	//unmarshal yaml file
	if err = yaml.Unmarshal(f, &e); err != nil {
		panic(err)
	}

	//save to map
	for _, errorCode := range e.Errors {
		Add(errorCode)
	}

	log.Printf("success register %d errors", len(e.Errors))
}

func Add(e Error) {
	if e.CustomCode != 999 {
		CustomError[e.CustomCode] = e
	}
}

func Remove(code int) {
	if code != 999 {
		delete(CustomError, code)
	}
}

func SetUnknownError(e Error) {
	if e.CustomCode == 999 {
		CustomError[e.CustomCode] = e
	}
}
