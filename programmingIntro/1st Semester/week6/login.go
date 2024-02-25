package main

import (
	"fmt"
)

func main() {
	const id, pass string = "user", "userpw"
	var idattempt, passattempt string
	var attempts uint
	var success bool = false
	fmt.Printf("Enter Username and Password\n")
	for success != true {
		fmt.Scan(&idattempt, &passattempt)
		if idattempt == id && passattempt == pass {
			success = true
			fmt.Printf("%d Failed attempts\nLogin Success!", attempts)
		} else {
			attempts++
		}
	}
}
