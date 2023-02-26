package errorSimple

import (
	"fmt"
	"github.com/pkg/errors"
)

func aFunc() error {
	err := errors.New("inFunc error")
	if err != nil {
		return errors.Wrap(err, "aFunc")
	}
	return nil
}

func anotherFunc() error {
	err := aFunc()
	if err != nil {
		return errors.Wrap(err, "anotherFunc")
	}
	return nil
}

func StartExample() {
    err := anotherFunc()
	if err != nil {
		fmt.Printf("main error: %+v\n",
			errors.Cause(err))
		return
	}
	fmt.Println("ok")
}

