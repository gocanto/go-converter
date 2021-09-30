package main

import (
	"fmt"
	"github.com/voyago/converter/environment"
	"github.com/voyago/converter/pkg/store"
)

func main() {
	env, _ := environment.Make("converter")

	request := store.NewCurrencyLayerRequest(env, "SGD")

	fmt.Println(request)
}
