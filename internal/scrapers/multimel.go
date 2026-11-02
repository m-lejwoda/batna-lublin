package scrappers

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func GetMultimelFlats() {
	BASE_URL := "https://multimel-nieruchomosci.pl"
	resp, err := http.Get("https://multimel-nieruchomosci.pl/mieszkania/dominow-rynek")
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		log.Fatalf("Problem with website")
	}
	document, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Fatalf("Problem with goquery")
	}
	selected := document.Find(".mm-card")
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
		price, _ = strconv.Atoi(strings.ReplaceAll(strings.TrimSuffix(singleSelected.Find(".mm-price-tag > span").First().Text(), " zł"), " ", ""))
		fmt.Println("price", price)
		pricePerMeter, _ = strconv.Atoi(strings.ReplaceAll(strings.TrimSuffix(strings.TrimSpace(singleSelected.Find(".mm-price-tag-per-m").First().Text()), " zł / m2"), " ", ""))
		fmt.Println("pricePerMeter", pricePerMeter)
		status = strings.TrimSpace(singleSelected.Find(".mm-status").Text())
		fmt.Println("status", status)
		href, exists := singleSelected.Find(".mm-media").First().Attr("href")
		fmt.Println("href", href)
		if !exists {
			log.Fatalf("Href not exists")
		}
		elementURL := strings.Split(href, "/")
		customID = elementURL[len(elementURL)-1]
		//if !exists {
		//	log.Fatalf("No href")
		//}
		fmt.Println(BASE_URL + href)
		flatPageResp, err := http.Get(BASE_URL + href)
		if err != nil {
			log.Fatalf("Cant get flat url")
		}
		defer flatPageResp.Body.Close()
		if err != nil {
			log.Fatalf("Error to get flatPage website")
		}
		documentFlat, err := goquery.NewDocumentFromReader(flatPageResp.Body)
		flatPage := documentFlat.Find(".mm-item")
		flatPage.Find("ul.mm-details-list").Find("li").Each(func(i int, singleDetailFlat *goquery.Selection) {
			sideValue := singleDetailFlat.Find("strong").Text()
			switch i {
			case 0:
				localNumber = sideValue
			case 1:
				buildingNumber = sideValue
			case 2:
				area = strings.TrimSuffix(sideValue, " m²")
			case 3:
				rooms, _ = strconv.Atoi(sideValue)
			case 4:
				floor = sideValue
			}
		})
		fmt.Println("customId", customID)
		fmt.Println("PricePerMeter", pricePerMeter)
		fmt.Println("localNumber", localNumber)
		fmt.Println("buildingNumber", buildingNumber)
		fmt.Println("area", area)
		fmt.Println("rooms", rooms)
		fmt.Println("floor", floor)
		fmt.Println(flatPage.Find(".mm-feature-box").Last().Find("strong").Text())
		balcony = strings.ReplaceAll(strings.TrimSuffix(strings.TrimSpace(flatPage.Find(".mm-feature-box").Last().Find("strong").Text()), "m²"), " ", "")
		fmt.Println("balcony", balcony)
		layoutLink, exists := documentFlat.Find(".btn-outline-primary.mm-ask-btn").First().Attr("href")
		if !exists {
			log.Fatalf("No href")
		}
		layoutUrl := BASE_URL + layoutLink
		fmt.Println("layoutUrl", layoutUrl)
		layout, err := http.Get(layoutUrl)
		if err != nil {
			log.Fatalf("Problem with Get")
		}
		defer layout.Body.Close()

		out, err := os.Create(customID + ".pdf")
		if err != nil {
			log.Fatalf("Cant create file")
		}
		defer out.Close()
		_, err = io.Copy(out, layout.Body)

		if err != nil {
			log.Fatalf("Can't save file")
		}
		fmt.Println("File Saved")
		os.Exit(1)
	})

}
