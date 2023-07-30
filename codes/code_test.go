package codes

import (
	"errors"
	"testing"
)

func TestError_Wrap(t *testing.T) {
	err := Wrap(errors.New("TEST"), 100)

	if err.Error() != "100: TEST" {
		t.Errorf("codes should be '100: TEST', current codes is %s", err.Error())
	}
}

func TestError_Extract(t *testing.T) {
	//add sample codes
	Add(Code{
		CustomCode:      105,
		ResponseMessage: "TEST",
		ErrorMessage:    "TEST",
		ResponseCode:    105,
	})

	err := Wrap(errors.New("TEST"), 105)

	newErr := Extract(err)

	if newErr.ErrorMessage != "TEST" {
		t.Errorf("codes message should be 'TEST', current codes message is %s", newErr.ErrorMessage)
	}

	if newErr.CustomCode != 105 {
		t.Errorf("codes code should be 105, current codes code is %d", newErr.CustomCode)
	}

	err = Wrap(errors.New("TEST"), 333)

	newErr = Extract(err)

	if newErr.CustomCode != 999 {
		t.Errorf("codes code should be 999, current codes code is %d", newErr.CustomCode)
	}

	//remove sample codes (buat 1 lgi fungsi buat remove codes / reset)
	Remove(105)
}
