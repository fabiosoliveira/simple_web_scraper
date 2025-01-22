package main

import (
	"fmt"

	"github.com/fabiosoliveira/simple_web_scraper/pkg/scrap"
)

func main() {
	ch := make(chan scrap.Product)

	go scrap.GetMercadoLivreProducts("bau bauleto givi e27", ch)

	count := 0
	for range ch {
		count++
	}

	fmt.Println(count)
}
