package main

import (
	"os"
	"testing"
)

func TestNewProductList(t *testing.T) {
	p := newProductList()

	if len(p) != 9 {
		t.Errorf("Expected length of 9, but got %v", len(p))
	}
}

func TestSaveToFileandGetProductListFromFile(t *testing.T) {
	os.Remove("_producttesting")
	products := newProductList()
	products.saveToFile("_producttesting")

	loadedProducts := getProductListFromFile("_producttesting")

	if len(loadedProducts) != 9 {
		t.Errorf("Expected length of 9, but got %v", len(loadedProducts))
	}
	os.Remove("_producttesting")
}
