package datautil_test

import (
	"errors"
	"testing"

	"github.com/alcionai/corso/internal/connector/datautil"
)

func TestCreateErrorList(t *testing.T) {
	listing := datautil.NewErrorList()
	if listing.GetLength() != 0 {
		t.Errorf("Incorrect initialization\n")
	}
}

func TestErrorListAdd(t *testing.T) {
	listing := datautil.NewErrorList()
	err1 := errors.New("Sample")
	err2 := errors.New("I have two")
	listing.AddError(&err1)
	listing.AddError(&err2)
	if listing.GetLength() != 2 {
		t.Errorf("Received: %d, Expected: 2\n", listing.GetLength())
	}
	t.Logf("Error Print: %s\n", listing.GetErrors())
}

func TestErrorListFormatCheck(t *testing.T) {
	listing := datautil.NewErrorList()
	emptyReturn := listing.GetErrors()
	if len(emptyReturn) != 0 {
		t.Errorf("Improper length\n")
	}
	err1 := errors.New("cauliflower")
	listing.AddError(&err1)
	if len(listing.GetErrors()) != 19 { // Err 0 cauliflower
		t.Errorf("Incorrect format")
	}
}
