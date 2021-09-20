package lib

import (
	"github.com/joho/godotenv"
	"runtime"
	"strings"
)

func Env(key string) interface{} {
	_, fileName, _, _ := runtime.Caller(0)

	baseDir := strings.Split(fileName, "/src/env.go")[0]

	env, _ := godotenv.Read(baseDir + "/.env")

	return env[key]
}
