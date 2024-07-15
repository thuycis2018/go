package main

import (
	"fmt"
)

func main() {
	// products := newProductList()
	// products.saveToFile("my_products")
	products := getProductListFromFile("my_products")
	products.getRandomProductList()
	fmt.Println(products)

	// interface
	t := triangle{base: 10, height: 10}
	s := square{sideLength: 10}
	printArea(t)
	printArea(s)

	// get argument typed in the terminal, e.g. my_products
	// printFile()

	//channel
	links := []string{
		"http://google.com",
		"http://golang.org",
	}

	c := make(chan string)

	for _, link := range links {
		go checkLink(link, c)
	}
	for i := 0; i < len(links); i++ {
		fmt.Println(<-c)
	}
}
