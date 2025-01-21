package main

import (
	"fmt"
	"net/http"

	"github.com/fabiosoliveira/simple_web_scraper/pkg/scrap"
)

func main() {
	c := scrap.NewCollector()

	c.OnHTML("h2", func(text string) {
		fmt.Println(text)
	})

	c.OnScraped(func(r *http.Response) {
		fmt.Println(r.StatusCode)
	})

	c.Visit("https://lista.mercadolivre.com.br/bau-bauleto-givi-e27")
}
