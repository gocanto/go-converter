package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"io"
	"net/http"
)

func main() {
	// 1 - create the request
	// 2 - invoke the handler

	myEnv, _ := godotenv.Read()

	fmt.Printf("Welcome to the [%s] library... \n", myEnv["LIBRARY_NAME"])

	resp, _ := http.Get("http://api.currencylayer.com/live?access_key=" + myEnv["CURRENCY_LAYER_KEY"])

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Println("error closing handler")
		}
	}(resp.Body)

	body, _ := io.ReadAll(resp.Body)

	fmt.Println(string(body))
}
