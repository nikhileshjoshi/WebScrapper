package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"log"
	"time"
)

type beerType struct {
	beer_id   string
	beer_name string
	brewery   string
	location  string
	style     string
	size      string
	abv       string
	ibus      string
}

func main() {
	t := time.Now()

	doc, err := goquery.NewDocument("http://craftcans.com/db.php?search=all&sort=beerid&ord=desc&view=text")
	if err != nil {
		log.Fatal(err)
	}

	var beers []beerType

	doc.Find("tr").Each(func(i int, s *goquery.Selection) {
		tds := s.Children()

		if tds.Size() == 8 {
			beer := beerType{tds.Slice(0, 1).Text(), tds.Slice(1, 2).Text(), tds.Slice(2, 3).Text(), tds.Slice(3, 4).Text(),
				tds.Slice(4, 5).Text(), tds.Slice(5, 6).Text(), tds.Slice(6, 7).Text(), tds.Slice(7, 8).Text()}
			beers = append(beers, beer)

		}
	})

	fmt.Println(beers)
	fmt.Println(time.Since(t))
}
