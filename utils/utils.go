package utils

import "fmt"

func SafeExec(f func()) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("Panic occurred: ", err)

		}
	}()
	f()
}
