package recoverexample

import "fmt"

func mayPanic() {
	defer func() {
		if r := recover(); r != nil {

			fmt.Println("Recovered. Error:\n", r)
		}
	}()
	panic("a problem")
}

func RecoverExample() {

	mayPanic()

	fmt.Println("After mayPanic()")
}
