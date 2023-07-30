package codes

import (
	"gopkg.in/yaml.v3"
	"log"
	"os"
)

var (
	CustomError map[int]Code
)

func init() {
	CustomError = map[int]Code{
		999: {
			CustomCode:      999,
			ResponseMessage: "need to register your custom code",
			ErrorMessage:    "unknown codes code",
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
		Errors []Code `yaml:"codes"`
	}{}

	//unmarshal yaml file
	if err = yaml.Unmarshal(f, &e); err != nil {
		panic(err)
	}

	//save to map
	for _, errorCode := range e.Errors {
		Add(errorCode)
	}

	log.Printf("success register %d codes", len(e.Errors))
}

func Add(e Code) {
	if e.CustomCode != 999 {
		CustomError[e.CustomCode] = e
	}
}

func Remove(code int) {
	if code != 999 {
		delete(CustomError, code)
	}
}

func SetUnknownError(e Code) {
	if e.CustomCode == 999 {
		CustomError[e.CustomCode] = e
	}
}
