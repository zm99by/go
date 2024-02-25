package main

import (
	"fmt"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {

	if len(os.Args[0:]) < 2 {
		fmt.Println("please add filename")
		os.Exit(1)
	}

	arg := os.Args[1]

	dat, err := os.ReadFile(arg)
	check(err)
	fmt.Print(string(dat))

}
