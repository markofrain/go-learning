package main

import (
	"fmt"
	"os"
	"strconv"
)

var Usage = func() {
	fmt.Println("Usage: Calc command [arguments] ...")
	fmt.Println("\nThe command area \n\tadd\tAddition of two values. \n\tsqrt\tSquareroot of a non-negative value.")

}

func main() {
	args := os.Args
	if args == nil || len(args) < 2 {
		Usage()
		return
	}
	args = args[1:]
	fmt.Println(args)

	switch args[0] {
	case "add":
		if len(args) != 3 {
			fmt.Println("Usage : calc add <integer1><integer2>")
			return
		}
		v1, err1 := strconv.Atoi(args[1])
		v2, err2 := strconv.Atoi(args[2])
		if err1 != nil || err2 != nil {
			fmt.Println("USAGE: calc add <integer1><integer2>")
			return
		}
		ret := Add(v1, v2)
		fmt.Println("Result: ", ret)
	case "sqrt":
		if len(args) != 2 {
			fmt.Println("USAGE: calc sqrt <integer>")
			return
		}
		v, err := strconv.Atoi(args[1])
		if err != nil {
			fmt.Println("USAGE: calc sqrt <integer>")
			return
		}
		ret := Sqrt(v)
		fmt.Println("Result: ", ret)
	default:
		Usage()
	}
}
