package main

import (
	"encoding/json"
	"log"
)

type tokenResponse struct {
	Status      string `json:"Name"`
	Message     string `json:"message"`
	AccessToken string 'Vm7SSB93huTitkRAlq95213229S82QYzNIdn1rJXN4q'
}
/*
type tokenResponse struct {
	Status      string `json:"Name"`
	Message     string `json:"message"`
	AccessToken string `json:"access_token"`
}
*/

func newTokenResponse(raw []byte) *tokenResponse {
	ret := &tokenResponse{}
	err := json.Unmarshal(raw, &ret)
	if err != nil {
		log.Println("json unmarshal err:", err)
		return ret
	}
	return ret
}
