package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

// new type
type product_list []string

func newProductList() product_list {
	products := product_list{}

	productColors := []string{"Blue", "Red", "Gray"}
	productTypes := []string{"Mouse", "Keyboard", "Laptop"}

	for _, color := range productColors {
		for _, _type := range productTypes {
			products = append(products, color+" "+_type)
		}
	}

	return products
}

// receiver of type product_list, p is working instance of this type
func (p product_list) print() {
	for i, product := range p {
		fmt.Println(i, product)
	}
}

// return 2 values of type product_list
func geProducts(p product_list, size int) (product_list, product_list) {
	return p[:size], p[size:]
}

func (p product_list) toString() string {
	return strings.Join([]string(p), ",")
}

func (p product_list) saveToFile(filename string) error {
	return os.WriteFile(filename, []byte(p.toString()), 0666)
}

func getProductListFromFile(filename string) product_list {
	bs, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	s := strings.Split(string(bs), ",")
	return product_list(s)
}

func (p product_list) getRandomProductList() {
	// set up seed for random number
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	for i := range p {
		newPosition := r.Intn(len(p) - 1)
		p[i], p[newPosition] = p[newPosition], p[i]
	}
}
