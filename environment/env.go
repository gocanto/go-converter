package environment

import (
	"github.com/joho/godotenv"
	"runtime"
	"strings"
)

type Env struct {
	Items map[string]string
}

func Make() (*Env, error) {
	_, fileName, _, _ := runtime.Caller(0)
	baseDir := strings.Split(fileName, "/env.go")[0]

	env, err := godotenv.Read(baseDir + "/.env")

	if err != nil {
		return nil, err
	}

	return &Env{Items: env}, nil
}

func (current Env) Get(key string) string {
	return current.Items[key]
}

func (current Env) IsDev() bool  {
	value, exist := current.Items["CONVERTER_ENV"]

	return exist && value == "DEV"
}
