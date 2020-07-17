package main

import (
	"encoding/json"
	"log"
)

type tokenResponse struct {
	Status      string `json:"Name"`
	Message     string `json:"message"`
	AccessToken string `kvAZ86e7a8XykA2KBP0a0LBfWwDWOChN8lDSSTn4N0A`
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
