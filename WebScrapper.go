package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"log"
	"net/http"
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

	//beers := make([]beer,100)
	var beers []beerType
	index := 0

	doc.Find("tr").Each(func(i int, s *goquery.Selection) {
		//fmt.Println(i, len(s.Contents().Nodes) ,s.Contents().Text())
		tds := s.Children()

		if tds.Size() == 8 {
			index++
			beer := beerType{tds.Slice(0, 1).Text(), tds.Slice(1, 2).Text(), tds.Slice(2, 3).Text(), tds.Slice(3, 4).Text(),
				tds.Slice(4, 5).Text(), tds.Slice(5, 6).Text(), tds.Slice(6, 7).Text(), tds.Slice(7, 8).Text()}
			/*beers[index].beer_id = tds.Slice(0,1).Text()
			beers[index].beer_name = tds.Slice(1,2).Text()
			beers[index].brewery = tds.Slice(2,3).Text()
			beers[index].location = tds.Slice(3,4).Text()
			beers[index].style = tds.Slice(4,5).Text()
			beers[index].size = tds.Slice(5,6).Text()
			beers[index].abv = tds.Slice(6,7).Text()
			beers[index].ibus = tds.Slice(7,8).Text()*/
			beers = append(beers, beer)

		}
	})

	fmt.Println(beers)

}
