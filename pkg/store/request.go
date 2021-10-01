package store

import "github.com/voyago/converter/environment"

const CurrencyLayerDriverName = "currency-layer"

type Request struct {
	env    environment.Env
	driver string
	source string
}

func NewCurrencyLayerRequest(env environment.Env, source string) Request {
	return Request{
		env:    env,
		source: source,
		driver: CurrencyLayerDriverName,
	}
}

func (current Request) GetEnv() environment.Env {
	return current.env
}

func (current Request) GetDriver() string {
	return current.driver
}

func (current Request) GetSource() string {
	return current.source
}
