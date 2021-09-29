package environment

import (
	"github.com/joho/godotenv"
	"os"
	"regexp"
)

const (
	FileName     = ".env"
	EnvKey       = "CONVERTER_APP_ENV"
	EnvLiveValue = "production"
)

type Env struct {
	items    *map[string]string
	filePath string
}

func Make(rootDir string) (Env, error) {
	path := resolveFilePath(rootDir, FileName)

	return buildFrom(path)
}

func MakeWith(rootDir string, file string) (Env, error) {
	path := resolveFilePath(rootDir, file)

	return buildFrom(path)
}

func (current Env) Get(key string) string {
	items := *current.items
	value := items[key]

	if value == "" {
		value = os.Getenv(key)
	}

	return value
}

func (current Env) IsLive() bool {
	return current.Get(EnvKey) == EnvLiveValue
}

func (current Env) IsTest() bool {
	return !current.IsLive()
}

func (current Env) Items() map[string]string {
	return *current.items
}

func resolveFilePath(directory string, fileName string) string {
	expression := regexp.MustCompile(`^(.*` + directory + `)`)

	cwd, _ := os.Getwd()
	path := expression.Find([]byte(cwd))

	return string(path) + `/` + fileName
}

func buildFrom(path string) (Env, error) {
	env := Env{}

	if _, err := os.Stat(path); os.IsNotExist(err) {
		return env, err
	}

	if items, err := godotenv.Read(path); err != nil {
		return env, err
	} else {
		env.items = &items
	}

	env.filePath = path

	return env, nil
}
