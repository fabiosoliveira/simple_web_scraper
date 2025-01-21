package scrap

import (
	"fmt"
	"io"
	"net/http"
)

type Collector struct {
	caracters       []rune
	tag             string
	htmlCallbacks   map[string]func(text string)
	scrapedCallback func(r *http.Response)
}

func NewCollector() *Collector {
	return &Collector{
		caracters: make([]rune, 0),
	}
}

func (c *Collector) Visit(url string) {
	res, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	defer res.Body.Close()

	for {
		buf := make([]byte, 1024)
		n, err := res.Body.Read(buf)
		if err == io.EOF {
			if n != 0 {
				n, str := c.read(buf[:n])
				if n > 5 {
					// fmt.Println(str)
					c.htmlCallbacks[c.tag](str)
				}
			}
			break
		}
		if err != nil {
			panic(err)
		}

		n, str := c.read(buf[:n])
		if n > 5 {
			// fmt.Println(str)
			c.htmlCallbacks[c.tag](str)
		}
	}

	c.scrapedCallback(res)
}

func (c *Collector) lenght() int {
	return len(c.caracters)
}

func (c *Collector) read(buf []byte) (int, string) {
	str := string(buf)
	for _, r := range str {
		if c.lenght() == 0 && r != '<' {
			continue
		}

		if c.lenght() > (len(c.tag)+1) && string((c.caracters)[:(len(c.tag)+1)]) != fmt.Sprintf("<%s", c.tag) {
			c.caracters = nil
			continue
		}

		if c.lenght() > (len(c.tag)+3) && string((c.caracters)[c.lenght()-(len(c.tag)+3):]) == fmt.Sprintf("</%s>", c.tag) {
			return c.lenght(), string(c.caracters)
		}

		c.caracters = append(c.caracters, r)
	}

	return c.lenght(), string(c.caracters)
}

func (c *Collector) OnScraped(callback func(r *http.Response)) {
	c.scrapedCallback = callback
}

func (c *Collector) OnHTML(querySelector string, callback func(text string)) {
	c.htmlCallbacks[querySelector] = callback
	c.tag = querySelector
}
