package examples

import (
	"fmt"
	"github.com/voyago/converter/environment"
	"github.com/voyago/converter/pkg/store/handler/currencyLayer"
	"io"
	"net/http"
)

func Fetch() {
	env, _ := environment.Make("converter")
	resp, _ := http.Get(currencyLayer.ApiEndpoint + "?source=SGD&access_key=" + env.Get(currencyLayer.ApiKeyName))

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Println("error closing handler")
		}
	}(resp.Body)

	body, _ := io.ReadAll(resp.Body)

	fmt.Println(string(body))
}
