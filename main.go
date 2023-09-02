package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"

	git "github.com/BE2047401/ala-ala/web-scraping"
	"github.com/gocolly/colly"
)

func main() {
	fmt.Println("ALa ala")

	git.WebScraping()

	fileName := "data-scraping1.csv"
	file, err := os.Create(fileName)
	if err != nil {
		log.Fatalf("Could not create file, err: %q", err)
		return
	}

	defer file.Close()

	write := csv.NewWriter(file)
	defer write.Flush()

	c := colly.NewCollector()
	c.OnHTML("table#customers", func(e *colly.HTMLElement) {
		e.ForEach("tr", func(_ int, el *colly.HTMLElement) {
			write.Write([]string{
				el.ChildText("td:nth-child(1)"),
				el.ChildText("td:nth-child(2)"),
				el.ChildText("td:nth-child(3)"),
			})
		})
		fmt.Println("Scarapping Complete")
	})
	c.Visit("https://www.w3schools.com/html/html_tables.asp")

	// c := colly.NewCollector(

	// 	colly.AllowedDomains("en.wikipedia.org"),
	// )

	// c.OnHTML(".mw-parser-output", func(e *colly.HTMLElement) {

	// 	links := e.ChildAttrs("a", "href")
	// 	fmt.Println(links)
	// })
	// c.Visit("https://en.wikipedia.org/wiki/Web_scraping")
}
