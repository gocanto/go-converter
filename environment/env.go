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
	rootDir  string
}

func Make(rootDir string) (Env, error) {
	path := resolveFilePath(rootDir, FileName)

	return buildFrom(rootDir, path)
}

func MakeWith(rootDir string, file string) (Env, error) {
	path := resolveFilePath(rootDir, file)

	return buildFrom(rootDir, path)
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

func (current Env) GetRootDir() string {
	return current.rootDir
}

func resolveFilePath(directory string, fileName string) string {
	expression := regexp.MustCompile(`^(.*` + directory + `)`)

	cwd, _ := os.Getwd()
	path := expression.Find([]byte(cwd))

	return string(path) + `/` + fileName
}

func buildFrom(rootDir string, path string) (Env, error) {
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
	env.rootDir = rootDir

	return env, nil
}
