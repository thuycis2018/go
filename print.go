package main

import (
	"fmt"
	"io"
	"os"
)

func printFile() {
	// get argument typed in the terminal, e.g. my_products
	f, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	io.Copy(os.Stdout, f)
}
