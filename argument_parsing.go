package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println(len(os.Args))

	//var sep, s string = " ", "";

	/* Argument parsing using OS */
	/*
	for i:=1; i < len(os.Args); i++ {
		//fmt.Println(os.Args[i])
		s += os.Args[i] + sep;
	}
	*/
	//argument parsing using parse
	//var sepParse, sParse string = " ",  "";
	/*
	for _, s := range(os.Args) {
		fmt.Println("argument is : " + s)
	}
	*/
	//Argument parsing using strings.join
	fmt.Println(strings.Join(os.Args, " "))
}