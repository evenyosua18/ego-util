package fiber_helper

type HttpResponse struct {
	// Code will store custom code that define in errors.yml
	Code int `json:"code"`
	// Message will store message from custom code that define in errors.yml
	Message string `json:"message" example:"message from custom code"`
	// ErrorMessage will store error message from system
	ErrorMessage string `json:"error_message" example:"error message from system"`
	// Data will store response for API
	Data interface{} `json:"data" swaggertype:"string" example:"object result data | will be nil if error"`
}

type Response struct {
	CustomCode      int    `yaml:"code"`
	ResponseMessage string `yaml:"message"`
	ErrorMessage    string `yaml:"errors"`
	ResponseCode    int    `yaml:"response_code"`
}
