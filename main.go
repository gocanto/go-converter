package main

import "github.com/voyago/converter/environment"

func main() {
	environment.Live(".env")

	//env = environment.Env{}

	//env, _ := environment.Make()
	//
	//fmt.Println("main", env.Items)
	//resp, _ := http.Get(currencyLayer.ApiEndpoint + "?source=SGD&access_key=" + env.Get("CONVERTER_CURRENCY_LAYER_KEY"))
	//
	//defer func(Body io.ReadCloser) {
	//	err := Body.Close()
	//	if err != nil {
	//		fmt.Println("error closing handler")
	//	}
	//}(resp.Body)
	//
	//body, _ := io.ReadAll(resp.Body)
	//
	//fmt.Println(string(body))
}
