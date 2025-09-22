package main

import (
	"fmt"

	"github.com/kuromittsu/util_password"
)

func validate() {

	fmt.Println("=========================================================")
	fmt.Println("START EXAMPLE VALIDATE")
	fmt.Println("=========================================================")

	testList := []string{
		"!Alfiras123",
		"!Alfiras1",
		"!wwwlfiras123",
		"Alfiras",
		"21313123",
	}

	for _, v := range testList {

		result := util_password.PasswordValidate(v, util_password.DefaultRules(), util_password.LANG_EN)

		fmt.Println("=========================================================")
		fmt.Println("using DefaultRules function")
		fmt.Println("=========================================================")

		fmt.Printf("result of %s: \n", v)

		fmt.Printf("is valid: %v \n", result.IsValid())

		fmt.Printf("invalid list (raw): %v \n", result.InvalidList().Get())

		invalidListStr, invalidListErr := result.InvalidList().GetJoined("")
		fmt.Printf("invalid list (joined | string): %v \n", invalidListStr)
		fmt.Printf("invalid list (joined | error): %v \n", invalidListErr)

		fmt.Println("=========================================================")
		fmt.Println()
	}

	fmt.Println("=========================================================")
	fmt.Println("END EXAMPLE VALIDATE")
	fmt.Println("=========================================================")
	fmt.Println()

}
