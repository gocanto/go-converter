package main

import (
	"fmt"
	"github.com/joho/godotenv"
)

func main() {
	myEnv, _ := godotenv.Read()

	fmt.Printf("Welcome to the [%s] library... \n", myEnv["LIBRARY_NAME"])
}
