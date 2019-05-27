package dmrmarc

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
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

func FetchUsers() ([]Users, error) {
	resp, err := http.Get("https://ham-digital.org/status/users.json")
	if err != nil {
		log.Printf("[dmr-marc][FetchUsers] ERROR: Failure fetching DMR-MARC users, %v", err)
		return nil, err
	}
	defer resp.Body.Close()

	log.Printf("[dmr-marc][FetchUsers] downloaded DMR-MARC users with size %d bytes", resp.ContentLength)
	if resp.StatusCode != 200 {
		b, _ := ioutil.ReadAll(resp.Body)
		log.Printf("[dmr-marc][FetchUsers] user list download failed with status = %d, body => %s", resp.StatusCode, string(b))
		return nil, fmt.Errorf("user list download failed with status = %d, body => %s", resp.StatusCode, string(b))
	}

	var u UsersResponse
	err = json.NewDecoder(resp.Body).Decode(&u)
	if err != nil {
		log.Printf("[dmr-marc][FetchUsers] ERROR: Failure parsing DMR-MARC users, %v", err)
		return nil, err
	}

	return u.Users, nil
}
