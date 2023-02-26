package errorSimple

import (
	"fmt"

	"github.com/pkg/errors"
)

func aFunc_() error {
	err := errors.New("A Mini Error !")
	if err != nil {
		return errors.Wrap(err, "aFunc_")
	}
	return nil
}

func anotherFunc_() error {
	err := aFunc_()
	if err != nil {
		return errors.Wrap(err, "anotherFunc_")
	}
	return nil
}

func StartClient2() {
	err := anotherFunc_()
	if err != nil {
		fmt.Printf("stack trace: %+v\n", errors.Cause(err))
		return
	}
	fmt.Println("ok")
}
