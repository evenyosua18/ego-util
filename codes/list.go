package codes

import (
	"gopkg.in/yaml.v3"
	"log"
	"os"
)

var (
	customCodes        map[int]Code
	defaultUnknownCode = 999
)

func init() {
	customCodes = map[int]Code{
		defaultUnknownCode: {
			CustomCode:      defaultUnknownCode,
			ResponseMessage: "need to register your custom code",
			ErrorMessage:    "unknown codes code",
			ResponseCode:    500,
		},
	}
}

func RegisterCode(path string) {
	//read file
	f, err := os.ReadFile(path)

	if err != nil {
		panic(err)
	}

	e := struct {
		Codes []Code `yaml:"codes"`
	}{}

	//unmarshal yaml file
	if err = yaml.Unmarshal(f, &e); err != nil {
		panic(err)
	}

	//save to map
	for _, code := range e.Codes {
		Add(code)
	}

	log.Printf("success register %d codes", len(e.Codes))
}

func Add(e Code) {
	if e.CustomCode != defaultUnknownCode {
		customCodes[e.CustomCode] = e
	}
}

func Remove(code int) {
	if code != defaultUnknownCode {
		delete(customCodes, code)
	}
}

func SetUnknownCode(e Code) {
	if e.CustomCode == defaultUnknownCode {
		customCodes[e.CustomCode] = e
	}
}

func List() map[int]Code {
	return customCodes
}
