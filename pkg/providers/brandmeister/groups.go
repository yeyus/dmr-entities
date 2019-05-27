package brandmeister

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func GetGroups() (data map[int]string, err error) {
	resp, err := http.Get("https://api.brandmeister.network/v1.0/groups/")
	if err != nil {
		log.Printf("[bm][GetGroups] ERROR: failure fetching Brandmeister groups: %v", err)
		return
	}
	defer resp.Body.Close()

	log.Printf("[bm][GetGroups] downloaded BM talkgroups with size %d bytes", resp.ContentLength)
	if resp.StatusCode != 200 {
		b, _ := ioutil.ReadAll(resp.Body)
		log.Printf("[bm][GetGroups] BM talkgroups list download failed with status = %d, body => %s", resp.StatusCode, string(b))
		return nil, fmt.Errorf("talkgroup list download failed with status = %d, body => %s", resp.StatusCode, string(b))
	}

	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		log.Printf("[bm][GetGroups] ERROR: failure parsing Bradmeister groups: %v", err)
		return
	}

	return
}
