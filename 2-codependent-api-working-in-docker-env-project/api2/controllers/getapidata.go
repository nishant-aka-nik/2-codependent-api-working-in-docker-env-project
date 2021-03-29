package controllers

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"sellerapi2/models"

	"github.com/Jeffail/gabs"
)

//ResponseData struct to store response data
type ResponseData struct {
	Message string `json:"message"`
}

//GetAPIData ffs
func GetAPIData(w http.ResponseWriter, r *http.Request) {

	// Validating the request method
	if r.Method != "POST" {
		http.Error(w, http.StatusText(500), 500)
		w.Header().Set("Content-Type", "application/json")
	}

	//reading the request body
	responseData, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(err.Error())
	}

	//parsing the request JSON using gabs package
	JSONData, err := gabs.ParseJSON([]byte(responseData))
	if err != nil {
		log.Println(err)
	}

	//checking for error in the payload
	if value, _ := JSONData.Path("url").Data().(string); value == "" {
		var responsedata ResponseData
		responsedata.Message = "Invalid JSON request"
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(responsedata)
		return
	} else if value, _ := JSONData.Path("product.name").Data().(string); value == "" {
		var responsedata ResponseData
		responsedata.Message = "Invalid JSON request"
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(responsedata)
		return
	} else if value, _ := JSONData.Path("product.imageURL").Data().(string); value == "" {
		var responsedata ResponseData
		responsedata.Message = "Invalid JSON request"
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(responsedata)
		return
	} else if value, _ := JSONData.Path("product.description").Data().(string); value == "" {
		var responsedata ResponseData
		responsedata.Message = "Invalid JSON request"
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(responsedata)
		return
	} else if value, _ := JSONData.Path("product.price").Data().(string); value == "" {
		var responsedata ResponseData
		responsedata.Message = "Invalid JSON request"
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(responsedata)
		return
	} else if value, _ := JSONData.Path("product.totalReviews").Data().(float64); value == 0 {
		var responsedata ResponseData
		responsedata.Message = "Invalid JSON request"
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(responsedata)
		return
	}

	//calling WriteAPIData to write data to mongoDB
	status := models.WriteAPIData(JSONData)

	//creating appropriate response
	if status == "errorinConnection" {
		var responsedata ResponseData
		responsedata.Message = "Error occurred in connecting mongoDB"
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(500)
		json.NewEncoder(w).Encode(responsedata)
		return

	} else if status == "error" {
		var responsedata ResponseData
		responsedata.Message = "Error occurred in updating mongoDB"
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(500)
		json.NewEncoder(w).Encode(responsedata)
		return

	} else if status == "update" {
		var responsedata ResponseData
		responsedata.Message = "Successfully updated the record in mongoDB"
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(200)
		json.NewEncoder(w).Encode(responsedata)
		return
	} else if status == "success" {
		var responsedata ResponseData
		responsedata.Message = "Successfully created the record in mongoDB"
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(200)
		json.NewEncoder(w).Encode(responsedata)
		return
	}

}
