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

	arg := os.Args[1]

	if len(os.Args) < 1 {
		fmt.Println("please add filename")
		os.Exit(1)
	}

	dat, err := os.ReadFile(arg)
	check(err)
	fmt.Print(string(dat))

}
