package dmrmarc

import (
	"encoding/json"
	"log"
	"net/http"
)

type RepeatersResponse struct {
	Repeaters []Repeater `json:"repeaters"`
}

type Repeater struct {
	Country     string `json:"country"`
	Locator     string `json:"locator"`
	Callsign    string `json:"callsing"`
	MapInfo     string `json:"map_info"`
	State       string `json:"state"`
	ColorCode   string `json:"color_code"`
	RadioId     string `json:"radio_id"`
	Frequency   string `json:"frequency"`
	Map         string `json:"map"`
	Trustee     string `json:"trustee"`
	TSLinked    string `json:"ts_linked"`
	City        string `json:"city"`
	Assigned    string `json:"peer"`
	IPSCNetwork string `json:"ipsc_network"`
	Offset      string `json:"offset"`
	Remarks     string `json:"remarks"`
}

func FetchRepeaters() []Repeater {
	resp, err := http.Get("http://www.dmr-marc.net/cgi-bin/trbo-database/datadump.cgi?table=repeaters&format=json&header=0")
	if err != nil {
		log.Panicf("ERROR: Failure fetching DMR-MARC repeaters, %v", err)
	}
	defer resp.Body.Close()

	var r RepeatersResponse

	err = json.NewDecoder(resp.Body).Decode(&r)
	if err != nil {
		log.Panicf("ERROR: Failure parsing DMR-MARC repeaters, %v", err)
	}

	return r.Repeaters
}
