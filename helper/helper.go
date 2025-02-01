package helper

import (
	"errors"
)

func ConverErrors(einterface interface{}, defaultErr string) error {
	switch einterface.(type) {
	case string:
		if msg, ok := interfaceToString(einterface); ok {
			return errors.New(msg)
		}
	case []string:
		if msg, ok := interfaceToStringList(einterface); ok && len(msg) > 0 {
			return errors.New(msg[0])
		}
	}

	return errors.New(defaultErr)
}

func interfaceToString(i interface{}) (string, bool) {
	str, ok := i.(string)
	return str, ok
}

func interfaceToStringList(i interface{}) ([]string, bool) {
	strList, ok := i.([]string)
	return strList, ok
}
