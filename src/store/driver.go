package store

import (
	"errors"
	"fmt"
	"github.com/voyago/converter/src/store/handler"
)

type Driver struct {
	Id string
}

func (current Driver) Build() (handler.Handler, error) {
	for _, seed := range current.allowed() {
		if seed == current.Id {
			return current.resolve()
		}
	}

	return current.trow()
}

func (current Driver) allowed() []string {
	return []string{
		"currencylayer",
	}
}

func (current Driver) resolve() (handler.Handler, error) {
	switch current.Id {
	case "currencylayer":
		return handler.CurrencyLayer{ApiKey: ""}, nil
	}

	return current.trow()
}

func (current Driver) trow() (handler.Handler, error) {
	return handler.Null{}, errors.New(fmt.Sprintf("The given driver [%string] is invalid", current.Id))
}

