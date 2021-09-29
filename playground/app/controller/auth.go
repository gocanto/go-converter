package controller

import (
	"fmt"
	"github.com/voyago/converter/environment"
)

func Store() {
	env, err := environment.Make("converter")

	fmt.Println(env, err)
	fmt.Println(env.Items())
}
