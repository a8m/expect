package matchers

import (
	"errors"
	"fmt"
	"reflect"
)

type ReceiveMatcher struct {
	ReceiveTo interface{}
}

func Receive() ReceiveMatcher {
	return ReceiveMatcher{}
}

func ReceiveTo(arg interface{}) ReceiveMatcher {
	return ReceiveMatcher{
		ReceiveTo: arg,
	}
}

func (m ReceiveMatcher) Match(actual interface{}) error {
	t := reflect.TypeOf(actual)
	if t.Kind() != reflect.Chan || t.ChanDir() == reflect.SendDir {
		return fmt.Errorf("%s is not a readable channel", t.String())
	}

	v := reflect.ValueOf(actual)
	rxValue, ok := v.TryRecv()

	if !ok {
		return errors.New("did not receive")
	}

	if m.ReceiveTo == nil {
		return nil
	}

	outType := reflect.TypeOf(m.ReceiveTo)
	if outType.Kind() != reflect.Ptr {
		return fmt.Errorf("%s is not a pointer type", outType.String())
	}

	if !reflect.TypeOf(rxValue.Interface()).AssignableTo(outType.Elem()) {
		return fmt.Errorf("can not assigned %s to %s",
			reflect.TypeOf(rxValue.Interface()).String(),
			reflect.TypeOf(m.ReceiveTo).String(),
		)
	}

	outValue := reflect.ValueOf(m.ReceiveTo)
	outValue.Elem().Set(rxValue)

	return nil
}
