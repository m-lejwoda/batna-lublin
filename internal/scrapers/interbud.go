package scrappers

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/PuerkitoBio/goquery"
)

func GetInterbudFlats() {
	BASE_URL := "https://interbud.com.pl/inwestycje/zielony-felin/"
	fmt.Println("Interbud")
	resp, err := http.Get(BASE_URL)
	if err != nil {
		log.Fatalf("Can't get Data")
	}
	defer resp.Body.Close()
	document, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Fatalf("Problem with go query")
	}
	selected := document.Find("col-lg-6 col-xl-4 local-item mb-4")
	selected.Each(func(i int, singleSelected *goquery.Selection) {
		customID := ""
		price := 0
		pricePerMeter := 0
		status := "Wolne"
		localNumber := ""
		buildingNumber := ""
		area := ""
		rooms := 0
		floor := "parter"
		balcony := "5.0"
		customID, _ = singleSelected.Attr("id")
		status, _ = singleSelected.Attr("data-status")
		area, _ = singleSelected.Attr("data-footage")
		buildingNumber, _ = singleSelected.Attr("data-building")
		roomsStr, _ := singleSelected.Attr("data-rooms")
		rooms, _ = strconv.Atoi(roomsStr)
		floor, _ = singleSelected.Attr("data-level")
		url, _ := singleSelected.Find("fw-400 fs-xxl-23 lh-xxl-34 pt-xxl-8 text-black").First().Attr("href")
		singleResp, _ := http.Get(url)
		defer singleResp.Body.Close()
		singleDocument, err := goquery.NewDocumentFromReader(singleResp.Body)
		if err != nil {
			log.Fatalf("Can't get document")
		}
		fmt.Println("customId", customID)
		fmt.Println("PricePerMeter", pricePerMeter)
		fmt.Println("localNumber", localNumber)
		fmt.Println("buildingNumber", buildingNumber)
		fmt.Println("area", area)
		fmt.Println("rooms", rooms)
		fmt.Println("floor", floor)
		fmt.Println("balcony", balcony)
		fmt.Println("status", status)
		fmt.Println("price", price)

		infoSelected := singleDocument.Find("col-lg-6 ps-lg-5 d-flex flex-column justify-content-between").First().Text()
		fmt.Println("infoSelected", infoSelected)

	})
}
