package scrap

import "fmt"

type Product struct {
	Link  string
	Name  string
	Value float64
}

func (p Product) String() string {
	return fmt.Sprintf("Link: %s\nName: %s\nValue: %f", p.Link, p.Name, p.Value)
}
