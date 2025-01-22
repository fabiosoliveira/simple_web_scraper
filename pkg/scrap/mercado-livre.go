package scrap

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/gocolly/colly"
)

func GetMercadoLivreProducts(seach string, ch chan<- Product) {
	// instantiate a new collector object
	c := colly.NewCollector()

	// OnError callback
	c.OnError(func(_ *colly.Response, err error) {
		log.Println("Something went wrong:", err)
	})

	c.OnHTML("ol li", func(e *colly.HTMLElement) {
		Link := e.ChildAttr("a", "href")
		Name := e.ChildText("h2")
		Value := sanitizerFloat(e.ChildText("div.poly-price__current span:nth-child(1)"))

		ch <- Product{
			Link,
			Name,
			Value,
		}
	})

	// OnScraped callback
	c.OnScraped(func(r *colly.Response) {
		fmt.Println("Finished scraping:", r.Request.URL.String())

		// for _, product := range products {
		// 	fmt.Println("\n\n", product)
		// }
		close(ch)
	})

	// open the target URL
	c.Visit("https://www.mercadolivre.com.br/jm/search?as_word=" + seach)
}

func sanitizerFloat(value string) float64 {
	newValue := strings.ReplaceAll((value), "R$", "")
	newValue = strings.ReplaceAll((newValue), ",", ".")
	number, err := strconv.ParseFloat(newValue, 64)

	if err != nil {
		number = 0
	}

	return number
}
