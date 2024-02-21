package error

import (
	"errors"
	"fmt"
	"testing"
)

func TestNewBadRequestError(t *testing.T) {
	var err error
	err = NewBadRequestError(errors.New("myerr"))
	fmt.Println(err)
}
