package main

import (
	"fmt"
	"log"
	"strings"
	"sync"

	"webCrawler/internal/collector"
	"webCrawler/internal/db"
	"webCrawler/internal/router"

	"github.com/gocolly/colly/v2"
)

var startUrl = "https://lwn.net/"
var visited = make(map[string]bool)
var visitedMutex sync.Mutex

func main() {
	r := router.SetupRouter()

	collector, err := collector.NewCollector(
		colly.AllowedDomains("www.lwn.net", "lwn.net"),
		colly.Async(true),
	)
	if err != nil {
		log.Fatal(err)
	}

	collector.OnHTML("a[href]", func(h *colly.HTMLElement) {
		link := h.Attr("href")
		absUrl := h.Request.AbsoluteURL(link)

		visitedMutex.Lock()
		if !visited[absUrl] {
			visited[absUrl] = true
			collector.Visit(absUrl)
		}
		visitedMutex.Unlock()
	})

	collector.OnHTML("div.ArticleText", func(h *colly.HTMLElement) {
		test := h.Text
		title, formattedText := formatText(test)
		url := h.Request.URL.String()

		insertData(title, url, formattedText)
	})

	collector.OnRequest(func(r *colly.Request) {
		fmt.Println("Посещено: ", r.URL.String())
	})

	err = collector.Visit(startUrl)
	if err != nil {
		log.Println("Error visiting URL:", err)
	}

	r.Run(":8080")
	collector.Wait()

	fmt.Println(visited, "ЭТО ПОСЕЩЕННЫЙ СПИСОК")
}

func formatText(text string) (string, string) {
	formattedText := strings.Join(strings.Fields(text), " ")

	words := strings.Fields(formattedText)
	title := ""
	for i := 0; i < 5 && i < len(words); i++ {
		title += words[i] + " "
	}
	title = strings.TrimSpace(title)

	return title, formattedText
}

func insertData(title, url, text string) {
	_, err := db.DB.Exec("INSERT INTO pages (title, url, text) VALUES ($1, $2, $3)", title, url, text)
	if err != nil {
		log.Println("Error inserting data into db ", err)
	}
}
