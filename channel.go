package main

import (
	"fmt"
	"net/http"
)

func checkLink(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "may be down!")
		c <- "Site may be down"
		return
	}

	fmt.Println(link, "is up!")
	c <- "It's up"
}
