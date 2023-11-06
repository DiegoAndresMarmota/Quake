package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

const endpoint = "example.earthquake.net"




func getInfoEarthquake(methods, parameters float32) (*Data, error) {
	resp, err := http.Get(endpoint)
	if err != nil {
		return nil, err
	}
	
	var data Data
	
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return log.Fatal(err)
	}

	return &data, nil

}


func main() {
	data, err := getInfoEarthquake()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(data)
}