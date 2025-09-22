package main

import (
	"fmt"
	"log"

	"github.com/kuromittsu/util_password"
)

func hash() {

	fmt.Println("=========================================================")
	fmt.Println("START EXAMPLE HASH")
	fmt.Println("=========================================================")

	t := "password"
	fmt.Printf("t: %v\n", t)

	fmt.Println("hashing ....")

	hash, err := util_password.PasswordHash(t)
	if err != nil {
		log.Fatalf("err | %v", err)
	}

	fmt.Printf("hash: %v\n", hash)

	fmt.Println("=== Comparing ===")

	cr := "password"
	cw := "passwordd"

	fmt.Printf("compare right: %v\n", cr)
	fmt.Printf("compare wrong: %v\n", cw)

	fmt.Println("=== Comparing right ===")

	err = util_password.PasswordCompare(hash, cr)
	if err != nil {
		fmt.Printf("comparing right | err: %v\n", err)
	} else {
		fmt.Println("comparing right success")
	}

	fmt.Println("=== Comparing wrong ===")

	err = util_password.PasswordCompare(hash, cw)
	if err != nil {
		fmt.Printf("comparing wrong | err: %v\n", err)
	} else {
		fmt.Println("comparing wrong success")
	}

	fmt.Println("=========================================================")
	fmt.Println("END EXAMPLE HASH")
	fmt.Println("=========================================================")
	fmt.Println()

}
