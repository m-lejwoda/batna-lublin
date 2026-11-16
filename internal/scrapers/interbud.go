package scrappers

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

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
	fmt.Println("Open document")
	selected := document.Find(".col-lg-6.col-xl-4.local-item.mb-4")
	selected.Each(func(i int, singleSelected *goquery.Selection) {
		customID := ""
		price := 0
		pricePerMeter := 0
		status := "Wolne"
		localNumber := ""
		buildingNumber := ""
		area := ""
		rooms := 0
		basement := 0.0
		floor := "parter"
		balcony := "5.0"
		customID, _ = singleSelected.Attr("id")
		status, _ = singleSelected.Attr("data-status")
		area, _ = singleSelected.Attr("data-footage")
		buildingNumber, _ = singleSelected.Attr("data-building")
		roomsStr, _ := singleSelected.Attr("data-rooms")
		rooms, _ = strconv.Atoi(roomsStr)
		floor, _ = singleSelected.Attr("data-level")
		url, _ := singleSelected.Find(".fw-400.fs-xxl-23.lh-xxl-34.pt-xxl-8.text-black").First().Attr("href")
		fmt.Println("url", url)
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

		//infoSelected := singleDocument.Find(".col-lg-6.ps-lg-5.d-flex.flex-column.justify-content-between").First().Text()
		infoSel := singleDocument.Find(".d-flex.justify-content-between.pb-1.fs-xxl-23.flex-column.flex-md-row.align-items-center.align-md-start")
		fmt.Println("infoSel")
		infoSel.Each(func(i int, singleRow *goquery.Selection) {
			fmt.Println("test")
			fmt.Println("singleRow", singleRow.Find("span.fw-300").First().Text())
			rowName := singleRow.Find("span.fw-300").First().Text()
			switch rowName {
			case "Pokoje":
				rooms, _ = strconv.Atoi(singleRow.Find(".color-2.fw-500").First().Text())
			case "Metraż":
				area = singleRow.Find(".color-2.fw-500").First().Text()

			case "Piwnica":
				basement, _ = strconv.ParseFloat(singleRow.Find(".color-1.fw-500").First().Text(), 64)
				fmt.Println("basement", basement)

			case "Mieszkanie":
				prices_string := strings.Split(singleRow.Find(".color-2.fw-500.text-nowrap").First().Text(), " (")
				price, _ = strconv.Atoi(strings.ReplaceAll(strings.TrimSuffix(prices_string[0], " zł"), " ", ""))
				pricePerMeter, _ = strconv.Atoi(strings.ReplaceAll(strings.TrimSuffix(prices_string[1], " zł/m2)"), " ", ""))
				fmt.Println("price1", price)
				fmt.Println("pricePerMeter1", pricePerMeter)
			}
		})
	})
}
