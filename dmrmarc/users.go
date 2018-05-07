package dmrmarc

import (
	"encoding/json"
	"log"
	"net/http"
)

type UsersResponse struct {
	Users []Users `json:"users"`
}

type Users struct {
	Country  string `json:"country"`
	Callsign string `json:"callsign"`
	Name     string `json:"name"`
	RadioId  string `json:"radio_id"`
	Surname  string `json:"surname"`
	State    string `json:"state"`
	City     string `json:"city"`
	Remarks  string `json:"remarks"`
}

func FetchUsers() []Users {
	resp, err := http.Get("https://ham-digital.org/status/users.json")
	if err != nil {
		log.Panicf("ERROR: Failure fetching DMR-MARC users, %v", err)
	}
	defer resp.Body.Close()

	var u UsersResponse

	err = json.NewDecoder(resp.Body).Decode(&u)
	if err != nil {
		log.Panicf("ERROR: Failure parsing DMR-MARC users, %v", err)
	}

	return u.Users
}
