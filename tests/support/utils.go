package support

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

func ParseJson(file string, target interface{}) error {
	content, err := os.Open(file)

	if err != nil {
		return err
	}

	byteValue, err := ioutil.ReadAll(content)

	if err != nil {
		return err
	}

	if err = json.Unmarshal(byteValue, &target); err != nil {
		return err
	}

	return nil
}
