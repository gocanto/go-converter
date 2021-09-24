package environment

import (
	"github.com/joho/godotenv"
	"runtime"
	"strings"
)

type Env struct {
	Items    map[string]string
	filePath string
}

func Make() (*Env, error) {
	filePath := parseFileName("env")
	env, err := godotenv.Read(filePath)

	if err != nil {
		return nil, err
	}

	return &Env{Items: env, filePath: filePath}, nil
}

func MakeWith(file string) (*Env, error) {
	filePath := parseFileName(file)
	env, err := godotenv.Read(filePath)

	if err != nil {
		return nil, err
	}

	return &Env{Items: env, filePath: filePath}, nil
}

func (current Env) Get(key string) string {
	return current.Items[key]
}

func (current Env) IsLive() bool {
	parts := strings.Split(current.filePath, "/.")[1]

	return parts == "env"
}

func (current Env) IsTest() bool {
	parts := strings.Split(current.filePath, "/.")[1]

	return parts == "testing"
}

func parseFileName(file string) string {
	_, fileName, _, _ := runtime.Caller(0)

	baseDir := strings.Split(fileName, "/env.go")[0]

	return baseDir + "/." + file
}
