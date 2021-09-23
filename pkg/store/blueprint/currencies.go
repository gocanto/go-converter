package blueprint

import (
	"fmt"
	"github.com/voyago/converter/pkg/support"
	"runtime"
	"strings"
)

type Blueprint struct {
	items []map[string]interface{}
}

func MakeCurrenciesBlueprint() (Blueprint, error) {
	body := Blueprint{}
	var response []map[string]interface{}

	if err := support.ParseJson(body.SourcePath(), &response); err != nil {
		return Blueprint{}, err
	}

	return Blueprint{items: response}, nil
}

func (body Blueprint) SourcePath() string {
	_, fileName, _, _ := runtime.Caller(0)
	baseDir := strings.Split(fileName, "/pkg/store/blueprint/currencies.go")[0]

	fmt.Println("===============>", baseDir)

	return baseDir + "/resources/currencies.json"
}

func (body Blueprint) Items() []map[string]interface{} {
	return body.items
}
