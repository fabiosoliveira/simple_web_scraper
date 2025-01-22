package scrap

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/gocolly/colly"
)

func GetMercadoLivreProducts(seach string, ch chan<- Product) {
	c := colly.NewCollector()

	c.OnError(func(_ *colly.Response, err error) {
		log.Println("Something went wrong:", err)
	})

	c.OnHTML("ol li", func(e *colly.HTMLElement) {
		Link := e.ChildAttr("a", "href")
		Name := e.ChildText("h3.poly-component__title-wrapper a")
		Value := sanitizerFloat(e.ChildText("div.poly-price__current span:nth-child(1)"))

		ch <- Product{
			Link,
			Name,
			Value,
		}
	})

	// OnScraped callback
	c.OnScraped(func(r *colly.Response) {
		fmt.Println("\nFinished scraping:", r.Request.URL.String())
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
