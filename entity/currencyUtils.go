package converter

import (
	"errors"
	"fmt"
)

func getCodeFrom(abstract interface{})  (string, error) {
	switch value := abstract.(type) {
		case string:
			return value, nil
		case Currency:
			return value.Code, nil
		default:
			return "", errors.New(fmt.Sprintf("The given abstract [%v] is invalid", abstract))
		}
}
