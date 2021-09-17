package main

import (
	"fmt"
	"github.com/joho/godotenv"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	myEnv, _ := godotenv.Read()

	fmt.Printf("Welcome to the [%s] library... \n", myEnv["LIBRARY_NAME"])
}
