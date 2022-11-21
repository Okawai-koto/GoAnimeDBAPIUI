package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

func main() {

	resp, err := http.Get("https://www.turkanime.co/anime/ani-ni-tsukeru-kusuri-wa-nai-3")
	if err != nil {
		log.Fatalln(err)
	}

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	//We Read the response body on the line below.
	doc.Find("#animedetay").Each(func(i int, s *goquery.Selection) {
		// For each item found, get the title
		title := s.Find(".ozet").Text()
		fmt.Println(title)

	})

}
