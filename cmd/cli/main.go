package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/fabiosoliveira/simple_web_scraper/pkg/scrap"
)

func main() {
	ch := make(chan scrap.Product)

	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <search>")
		return
	}

	search := strings.Join(os.Args[1:], " ")

	go scrap.GetMercadoLivreProducts(search, ch)

	for product := range ch {
		fmt.Println()
		fmt.Println(product)
	}

	fmt.Printf("Search: %s\n", search)
}
