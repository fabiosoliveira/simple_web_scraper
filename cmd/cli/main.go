package main

import (
	"fmt"
	"log"

	"github.com/gocolly/colly"
)

func main() {
	// instantiate a new collector object
	c := colly.NewCollector()

	// OnError callback
	c.OnError(func(_ *colly.Response, err error) {
		log.Println("Something went wrong:", err)
	})

	c.OnHTML("ol li", func(e *colly.HTMLElement) {
		fmt.Println(e.ChildAttr("a", "href"))
		// product.Image = e.ChildAttr("img", "src")
		// product.Name = e.ChildText(".product-name")
		fmt.Println(e.ChildText("h2"))
		fmt.Println(e.ChildText("div.poly-price__current span:nth-child(1)"))
		fmt.Println()
	})

	// OnScraped callback
	c.OnScraped(func(r *colly.Response) {
		fmt.Println("Finished scraping:", r.Request.URL.String())
	})

	// open the target URL
	c.Visit("https://lista.mercadolivre.com.br/bau-bauleto-givi-e27")
}
