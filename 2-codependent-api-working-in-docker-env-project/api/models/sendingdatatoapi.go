package models

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

//DataFromAPI2 struct to store response from api2
type DataFromAPI2 struct {
	Message string `json:"message"`
}

//SendDataToAPI to send the json to another api
func SendDataToAPI(scrapeddataOBJ ScrapedData) string {

	postBody, _ := json.Marshal(scrapeddataOBJ)

	requestBody := bytes.NewBuffer(postBody)

	resp, err := http.Post("http://host.docker.internal:8081/getapidata", "application/json", requestBody)
	if err != nil {
		log.Println(err.Error())
		return "err"
	}

	defer resp.Body.Close()

	//Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err.Error())
		return "err2"
	}

	var DataFromAPI2OBj DataFromAPI2

	json.Unmarshal(body, &DataFromAPI2OBj)

	log.Println(DataFromAPI2OBj.Message)

	return DataFromAPI2OBj.Message
}
