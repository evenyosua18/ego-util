package error

import (
	"errors"
	"testing"
)

func TestError_Wrap(t *testing.T) {
	err := Wrap(errors.New("TEST"), 100)

	if err.Error() != "100: TEST" {
		t.Errorf("error should be '100: TEST', current error is %s", err.Error())
	}
}

func TestError_Extract(t *testing.T) {
	//add sample error
	Add(Error{
		CustomCode:      105,
		ResponseMessage: "TEST",
		ErrorMessage:    "TEST",
		ResponseCode:    105,
	})

	err := Wrap(errors.New("TEST"), 105)

	newErr := Extract(err)

	if newErr.ErrorMessage != "TEST" {
		t.Errorf("error message should be 'TEST', current error message is %s", newErr.ErrorMessage)
	}

	if newErr.CustomCode != 105 {
		t.Errorf("error code should be 105, current error code is %d", newErr.CustomCode)
	}

	err = Wrap(errors.New("TEST"), 333)

	newErr = Extract(err)

	if newErr.CustomCode != 999 {
		t.Errorf("error code should be 999, current error code is %d", newErr.CustomCode)
	}

	//remove sample error (buat 1 lgi fungsi buat remove error / reset)
	Remove(105)
}
