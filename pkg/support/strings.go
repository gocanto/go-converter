package support

import (
	"strconv"
	"strings"
)

type Strings struct {
	Value string
}

func (current *Strings) RightPad(padStr string, times int) {
	padCountInt := 1 + ((times - len(padStr)) / len(padStr))

	(*current).Value = current.Value + strings.Repeat(padStr, padCountInt)
}

func (current Strings) ToInt64() (int64, error)  {
	result, err := strconv.ParseInt(current.Value, 10, 64)

	if err != nil {
		return 1, err
	}

	return result, nil
}
