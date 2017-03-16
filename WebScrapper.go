package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"log"
	"net/http"
)

type beer struct {
	beer_id   int
	beer_name string
	brewery   string
	location  string
	style     string
	size      string
	abv       string
	ibus      string
}

func main() {
	resp, err := http.Get("http://craftcans.com/db.php?search=all&sort=beerid&ord=desc&view=text")
	if err != nil {
		log.Fatal(err)
	}

	b := resp.Body
	defer b.Close()

	doc, err := goquery.NewDocument("http://craftcans.com/db.php?search=all&sort=beerid&ord=desc&view=text")
	if err != nil {
		log.Fatal(err)
	}

	doc.Find("tr").Each(func(i int, s *goquery.Selection) {
		//fmt.Println(i, len(s.Contents().Nodes) ,s.Contents().Text())
		if len(s.Contents().Nodes) == 8 {
			fmt.Println(s.Children().First().Text(), s.Children().Next().Text(), s.Children().Next().Text())
		}
		/*s.First().Find("td").Each(func(i int, s *goquery.Selection){
			fmt.
		})*/
	})

}
