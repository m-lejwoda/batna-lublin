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
	fmt.Println("Get Flats")
	fmt.Println("Multimel_v2")
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
	numberOfElements := len(selected.Nodes)
	fmt.Println("numberOfElements", numberOfElements)
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
		balcony := 0
		fmt.Println(strings.TrimSuffix(singleSelected.Find(".mm-price-tag > span").First().Text(), " zł"))
		price, _ = strconv.Atoi(strings.ReplaceAll(strings.TrimSuffix(singleSelected.Find(".mm-price-tag > span").First().Text(), " zł"), " ", ""))
		fmt.Println("price", price)
		pricePerMeter, _ = strconv.Atoi(strings.TrimSuffix(singleSelected.Find("mm-price-tag-per-m").Find("#text").Text(), " zł / m"))
		status = singleSelected.Find("mm-status").Text()
		fmt.Println("status", status)
		fmt.Println(singleSelected.Find("mm-media").Text())
		href, exists := singleSelected.Find("mm-media").Attr("href")
		fmt.Println("Tu jest błąd")
		fmt.Println("href", href)
		elementURL := strings.Split(href, "/")
		customID = elementURL[len(elementURL)-1]
		fmt.Println("customId", customID)
		fmt.Println("localNumber", localNumber)
		fmt.Println("buildingNumber", buildingNumber)
		fmt.Println("area", area)
		fmt.Println("rooms", rooms)
		fmt.Println("floor", floor)
		fmt.Println("balcony", balcony)
		if !exists {
			log.Fatalf("No href")
		}
		fmt.Println("price", price)
		fmt.Println("pricePerMeter", pricePerMeter)
		fmt.Println("status", status)
		flatPageResp, err := http.Get(href)
		defer flatPageResp.Body.Close()
		if err != nil {
			log.Fatalf("Error to get flatPage website")
		}
		documentFlat, err := goquery.NewDocumentFromReader(flatPageResp.Body)
		flatPage := documentFlat.Find("mm-item")
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
		balcony, _ = strconv.Atoi(strings.TrimSuffix(flatPage.Find("mm-feature-box").Last().Find("strong").Text(), " m²"))
		fmt.Println(documentFlat.Find("btn-outline-primary mm-ask-btn"))
		layoutLink, exists := documentFlat.Find("btn-outline-primary mm-ask-btn").Attr("href")
		if !exists {
			log.Fatalf("No href")
		}
		layoutUrl := "https://multimel-nieruchomosci.pl/" + layoutLink
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

	})

}
