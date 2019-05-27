package radioid

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gocarina/gocsv"
)

type NXDNUser struct {
	SystemId uint16 `csv:"RADIO_ID"`
	Callsign string `csv:"CALLSIGN"`
	Name     string `csv:"FIRST_NAME"`
	Surname  string `csv:"LAST_NAME"`
	City     string `csv:"CITY"`
	State    string `csv:"STATE"`
	Country  string `csv:"COUNTRY"`
	Remarks  string `csv:"REMARKS"`
}

func FetchUsers() ([]NXDNUser, error) {
	resp, err := http.Get("https://www.radioid.net/static/nxdn.csv")
	if err != nil {
		log.Printf("[radioid][FetchUsers] ERROR: failure fetching RadioID NXDN users, %v", err)
		return nil, err
	}
	defer resp.Body.Close()

	log.Printf("[radioid][FetchUsers] downloaded BM talkgroups with size %d bytes", resp.ContentLength)
	if resp.StatusCode != 200 {
		b, _ := ioutil.ReadAll(resp.Body)
		log.Printf("[radioid][FetchUsers] ERROR: nxdn user list download failed with status = %d, body => %s", resp.StatusCode, string(b))
		return nil, fmt.Errorf("nxdn us list download failed with status = %d, body => %s", resp.StatusCode, string(b))
	}

	var users []NXDNUser
	err = gocsv.Unmarshal(resp.Body, &users)
	if err != nil {
		log.Printf("ERROR: Failure parsing RadioID NXDN users, %v", err)
		return users, err
	}

	return users, err
}
