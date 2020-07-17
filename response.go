package main

import (
	"encoding/json"
	"log"
)

type tokenResponse struct {
	Status      string `json:"Name"`
	Message     string `json:"message"`
	AccessToken string `json:"access_token"`
}

json_string := `
    {
        "Name": "200",
        "message": "access_token is issued",
        "access_token": "Vm7SSB93huTitkRAlq95213229S82QYzNIdn1rJXN4q"
    }`

emp1 := new(tokenResponse)

json.Unmarshal([]byte(json_string), emp1)

/*
type tokenResponse struct {
	Status      string `json:"Name"`
	Message     string `json:"message"`
	AccessToken string `json:"access_token"`
}
*/

func newTokenResponse(raw []byte) *tokenResponse {
	ret := &tokenResponse{}
	ret := &emp1{}
	err := json.Unmarshal(raw, &ret)
	if err != nil {
		log.Println("json unmarshal err:", err)
		return ret
	}
	return ret
}
