package controllers

import (
	"encoding/json"
	"net/http"
	"sellerapi/models"
	"strings"
)

//ResponseData struct to store response data
type ResponseData struct {
	Message         string              `json:"message"`
	MessageFromAPI2 string              `json:"messagefromapi2,omitempty" `
	Scrapeddata     *models.ScrapedData `json:"scrapeddata,omitempty" `
}

//Geturl to get url to scrape data from
func Geturl(w http.ResponseWriter, r *http.Request) {

	// Validating the request method
	if r.Method != "POST" {
		http.Error(w, http.StatusText(500), 500)
		w.Header().Set("Content-Type", "application/json")
	}

	var requestdata models.RequestData

	//decoding the request data to struct
	err := json.NewDecoder(r.Body).Decode(&requestdata)

	//erro handling of decoded json
	if err != nil {
		var responsedata ResponseData
		responsedata.Message = "Error in decoding JSON"
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(responsedata)
		return
	}

	//error handling the request data
	if requestdata.URL == "" {
		var responsedata ResponseData
		responsedata.Message = "Invalid JSON request"
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(responsedata)
		return
	}

	//Calling Geturldetails to scrape the data from amazon
	scrapeddataOBJ, statusChecker := models.Geturldetails(requestdata)

	if statusChecker == "error" || statusChecker == "" {
		var responsedata ResponseData
		responsedata.Message = "Unable to fetch data from URL, Please try again"
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(responsedata)
		return
	}

	//calling SendDataToAPI to call and send scarped data to API2
	responseFormAPI2 := models.SendDataToAPI(scrapeddataOBJ)
	if responseFormAPI2 == "err" {
		var responsedata ResponseData
		responsedata.Message = "Error while fetching data from API2 or API2 is not running"
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(responsedata)
		return
	} else if responseFormAPI2 == "err2" {
		var responsedata ResponseData
		responsedata.Message = "Error while decoding body of API2"
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(responsedata)
		return
	} else if strings.Contains(responseFormAPI2, "Successfully") {
		var responsedata ResponseData
		responsedata.Message = "Successfully called and created/updated data from API2 to mongoDB "
		responsedata.MessageFromAPI2 = responseFormAPI2
		responsedata.Scrapeddata = &scrapeddataOBJ
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(200)
		json.NewEncoder(w).Encode(responsedata)
		return
	} else {
		var responsedata ResponseData
		responsedata.Message = "Error in creating/updating record in mongoDB from API2 "
		responsedata.MessageFromAPI2 = responseFormAPI2
		responsedata.Scrapeddata = &scrapeddataOBJ
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(responsedata)
		return
	}

}
