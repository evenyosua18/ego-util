package errors

import (
	"errors"
	"testing"
)

func TestError_Wrap(t *testing.T) {
	err := Wrap(errors.New("TEST"), 100)

	if err.Error() != "100: TEST" {
		t.Errorf("errors should be '100: TEST', current errors is %s", err.Error())
	}
}

func TestError_Extract(t *testing.T) {
	//add sample errors
	Add(Code{
		CustomCode:      105,
		ResponseMessage: "TEST",
		ErrorMessage:    "TEST",
		ResponseCode:    105,
	})

	err := Wrap(errors.New("TEST"), 105)

	newErr := Extract(err)

	if newErr.ErrorMessage != "TEST" {
		t.Errorf("errors message should be 'TEST', current errors message is %s", newErr.ErrorMessage)
	}

	if newErr.CustomCode != 105 {
		t.Errorf("errors code should be 105, current errors code is %d", newErr.CustomCode)
	}

	err = Wrap(errors.New("TEST"), 333)

	newErr = Extract(err)

	if newErr.CustomCode != 999 {
		t.Errorf("errors code should be 999, current errors code is %d", newErr.CustomCode)
	}

	//remove sample errors (buat 1 lgi fungsi buat remove errors / reset)
	Remove(105)
}
