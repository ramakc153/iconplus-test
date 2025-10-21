package practicaltest

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

const (
	BOOKINGURL  = "https://66876cc30bc7155dc017a662.mockapi.io/api/dummy-data/bookingList"
	KONSUMSIURL = "https://6686cb5583c983911b03a7f3.mockapi.io/api/dummy-data/masterJenisKonsumsi"
)

type BookingBody struct {
	BookingDate     string `json:"bookingDate"`
	OfficeName      string `json:"officeName"`
	StartTime       string `json:"startTime"`
	EndTime         string `json:"endTime"`
	ListConsumption []struct {
		Name string `json:"name"`
	} `json:"listConsumption"`
	Participants int    `json:"participants"`
	RoomName     string `json:"roomName"`
	Id           string `json:"id"`
}

type KonsumsiBody struct {
	CreatedAt string `json:"createdAt"`
	Name      string `json:"name"`
	MaxPrice  int    `json:"maxPrice"`
	Id        string `json:"id"`
}

var bookingData []BookingBody
var konsumsiData []KonsumsiBody

func FetchBody(url string, dataStruct interface{}) {
	// var bookingData []BookingBody
	resp, err := http.Get(url)
	if err != nil {
		log.Println("error when fetching: ", err.Error())
		return
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println("error when reading response body: ", body)
	}
	err = json.Unmarshal(body, &dataStruct)
	if err != nil {
		log.Panicln("error when unmarshaling data: ", err)
	}
	// for _, v := range dataStruct {
	// 	fmt.Println(v)

	// }
	// return
	defer resp.Body.Close()
}

func Server() {
	fmt.Println("Hello from pracitaltest server")
	fmt.Println("=============BOOKING DATA==========")
	FetchBody(BOOKINGURL, &bookingData)
	for _, v := range bookingData {
		fmt.Println(v)
	}
	fmt.Println("=============KONSUMSI DATA==========")
	FetchBody(KONSUMSIURL, &konsumsiData)
	for _, v := range konsumsiData {
		fmt.Println(v)
	}
}
