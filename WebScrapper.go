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
		tds := s.Children()
		if s.Contents().Size() == 8 {
			fmt.Println(tds.Slice(0,1).Text(), "--" ,tds.Slice(1,2).Text(),tds.Slice(2,3).Text(),tds.Slice(3,4).Text(),
													tds.Slice(4,5).Text(),tds.Slice(5,6).Text(),tds.Slice(6,7).Text(),tds.Slice(7,8).Text())
		}
		/*s.First().Find("td").Each(func(i int, s *goquery.Selection){
			fmt.
		})*/
	})

}
