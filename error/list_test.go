package error

import "testing"

func TestManageListError_Add(t *testing.T) {
	if len(CustomError) != 1 {
		t.Errorf("default list custom error should be one")
	}

	//add first error
	Add(Error{
		CustomCode:      500,
		ResponseMessage: "TEST",
		ErrorMessage:    "TEST",
		ResponseCode:    500,
	})

	//check
	if len(CustomError) != 2 {
		t.Errorf("list custom error should be two")
	}

	//add second error
	Add(Error{
		CustomCode:      400,
		ResponseMessage: "TEST",
		ErrorMessage:    "TEST",
		ResponseCode:    400,
	})

	//check
	if len(CustomError) != 3 {
		t.Errorf("list custom error should be three, total current error: %d", len(CustomError))
	}
}

func TestManageListError_SetUnknownError(t *testing.T) {
	//call function SetUnknownError, custom code not 999
	SetUnknownError(Error{
		CustomCode:      999,
		ResponseMessage: "TEST",
		ErrorMessage:    "TEST",
		ResponseCode:    100,
	})

	if CustomError[999].ErrorMessage != "TEST" {
		t.Errorf("unknown error message should be 'TEST', current message: %s", CustomError[999].ErrorMessage)
	}

	if CustomError[999].ResponseMessage != "TEST" {
		t.Errorf("unknown response message should be 'TEST', current message: %s", CustomError[999].ResponseMessage)
	}

	//call function SetUnknownError, but custom code not 999
	SetUnknownError(Error{
		CustomCode:      500,
		ResponseMessage: "TEST INVALID",
		ErrorMessage:    "TEST INVALID",
		ResponseCode:    100,
	})

	if CustomError[999].ErrorMessage != "TEST" {
		t.Errorf("unknown error message should be 'TEST', current message: %s", CustomError[999].ErrorMessage)
	}

	if CustomError[999].ResponseMessage != "TEST" {
		t.Errorf("unknown response message should be 'TEST', current message: %s", CustomError[999].ResponseMessage)
	}
}

func TestManageListError_RegisterError(t *testing.T) {
	//register error from yaml file, add one error
	RegisterError("./test.yaml")

	if len(CustomError) != 4 {
		t.Errorf("list custom error should be four, total current error: %d", len(CustomError))
	}
}

func TestManageListError_Remove(t *testing.T) {
	//remove one code from list
	Remove(300)

	if len(CustomError) != 3 {
		t.Errorf("list custom error should be three, total current error: %d", len(CustomError))
	}
}
