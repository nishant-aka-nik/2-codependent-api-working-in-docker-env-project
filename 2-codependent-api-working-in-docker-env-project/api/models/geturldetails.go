package models

import (
	"log"
	"strconv"
	"strings"

	"github.com/gocolly/colly"
)

//RequestData struct to store request data
type RequestData struct {
	URL string `json:"url"`
}

//ScrapedData full data
type ScrapedData struct {
	URL     string  `json:"url"`
	Product Product `json:"product"`
}

//Product - scrapped data from url
type Product struct {
	Name         string `json:"name"`
	ImageURL     string `json:"imageURL"`
	Description  string `json:"description"`
	Price        string `json:"price"`
	TotalReviews int64  `json:"totalReviews"`
}

//Geturldetails function to scrape relevant details from URL
func Geturldetails(requestdata RequestData) (ScrapedData, string) {

	var statusChecker string
	var scrapeddataOBJ ScrapedData
	var productOBJ Product

	c := colly.NewCollector()

	// On every a element which has href attribute call callback
	c.OnHTML("#ppd", func(e *colly.HTMLElement) {
		Name := e.ChildText("#productTitle")
		productOBJ.Name = Name
		log.Println(Name)

		imageURL := e.ChildAttr("#landingImage", "data-old-hires")
		productOBJ.ImageURL = imageURL
		log.Println(imageURL)

		var description string = ""
		e.ForEach("#feature-bullets > ul > li", func(_ int, f *colly.HTMLElement) {

			description = description + strings.TrimSpace(f.Text)
		})

		description = strings.ReplaceAll(description, "\n", "")
		description = strings.ReplaceAll(description, "Make sure this fitsby entering your model number.", "")

		productOBJ.Description = description
		log.Println(description)

		price := e.ChildText("#priceblock_ourprice")
		productOBJ.Price = price
		log.Println(price)

		totalReviews := e.ChildText("#acrCustomerReviewText")

		totalReviews = strings.TrimFunc(totalReviews, func(c rune) bool {
			return c <= 48 || c >= 57
		})

		totalReviews = strings.ReplaceAll(totalReviews, ",", "")

		if totalReviews, err := strconv.ParseInt(totalReviews, 10, 32); err == nil {
			productOBJ.TotalReviews = totalReviews
			log.Println(totalReviews)

		} else {
			log.Println(err)

		}

		scrapeddataOBJ.Product = productOBJ
		scrapeddataOBJ.URL = requestdata.URL

		statusChecker = "successful"

	})

	// Before making a request print "Visiting ..."
	c.OnRequest(func(r *colly.Request) {
		log.Println("Visiting", r.URL.String())
	})

	// Set error handler
	c.OnError(func(r *colly.Response, err error) {
		log.Println("Request URL:", r.Request.URL, "failed with response:", r, "\nError:", err)
		statusChecker = "error"
	})

	c.Visit(requestdata.URL)

	return scrapeddataOBJ, statusChecker

}
