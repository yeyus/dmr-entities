package brandmeister

import (
	"encoding/json"
	"log"
	"net/http"
)

func GetBrandmeisterGroups() (data map[int]string) {
	resp, err := http.Get("https://api.brandmeister.network/v1.0/groups/")
	if err != nil {
		log.Printf("[ERROR] Failure fetching Brandmeister groups: %v", err)
		return
	}
	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		log.Printf("[ERROR] Failure parsing Bradmeister groups: %v", err)
		return
	}

	log.Printf("Parsed %d groups", len(data))

	return
}
