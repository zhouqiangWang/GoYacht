package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Location struct {
	Lat float64
	Lng float64
}

type Geometry struct {
	Location      Location
	Location_type string
}

type AddressResult struct {
	Formatted_address string
	Geometry          Geometry
	PlaceId           string
}

type GeoResp struct {
	Results []AddressResult
	Status  string
}

type HttpResp struct {
	Url  string
	Resp *http.Response
	Err  error
	Body []byte
}

func AyncGet(url string) <-chan *HttpResp {
	ch := make(chan *HttpResp)

	go func(u string) {
		resp, err := http.Get(u)

		var body []byte
		if err == nil {
			defer resp.Body.Close()
			body, _ = io.ReadAll(resp.Body)

			// fmt.Printf("%+v", geoResp)
		}

		ch <- &HttpResp{u, resp, err, body}
	}(url)

	return ch
}

func main() {
	url1 := "https://maps.googleapis.com/maps/api/geocode/json?address=1600+Amphitheatre+Parkway%2C+Mountain+View%2C+CA&key=AIzaSyA92zQXh2pdnqjFfAptSxJMbHOjjXM2PjA"
	url2 := "https://maps.googleapis.com/maps/api/geocode/json?address=3473+N+1st+Street%2C+San+Jose%2C+CA&key=AIzaSyA92zQXh2pdnqjFfAptSxJMbHOjjXM2PjA"

	ch1, ch2 := AyncGet(url1), AyncGet(url2)
	addr1, addr2 := <-ch1, <-ch2

	var geoResp1 GeoResp
	var geoResp2 GeoResp

	err := json.Unmarshal(addr1.Body, &geoResp1)
	if err != nil {
		fmt.Println("error: ", err)
	}
	err = json.Unmarshal(addr2.Body, &geoResp2)
	if err != nil {
		fmt.Println("error: ", err)
	}

	loc1 := geoResp1.Results[0].Geometry.Location
	loc2 := geoResp2.Results[0].Geometry.Location

	fmt.Println(loc1.Lat)
	fmt.Println(loc2)

	DisUrl := fmt.Sprintf("https://maps.googleapis.com/maps/api/distancematrix/json?origins=%f,%f&destinations=side_of_road:%f,%f&key=AIzaSyA92zQXh2pdnqjFfAptSxJMbHOjjXM2PjA", loc1.Lat, loc1.Lng, loc2.Lat, loc2.Lng)

	ch3 := AyncGet(DisUrl)

	resp3 := <-ch3

	fmt.Println(string(resp3.Body))

	fmt.Println("hello world")
}
