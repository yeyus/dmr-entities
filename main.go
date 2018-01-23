package main

import (
	"fmt"
	"github.com/yeyus/dmr-entitites/dmrmarc"
	"log"
	"net/http"
)

type Entity struct {
	Id       uint32
	System   SystemTypeEnum
	Type     EntityTypeEnum
	Callsign string
	Name     string
	Surname  string
	Country  string
	State    string
	City     string
}

type EntityTypeEnum string

const (
	EntityHam       EntityTypeEnum = "ham"
	EntityTalkgroup EntityTypeEnum = "talkgroup"
)

type SystemTypeEnum string

const (
	SystemTypeDMR SystemTypeEnum = "DMR"
)

func main() {

}

func GetBrandmeisterGroups() {
	resp, err := http.Get("https://api.brandmeister.network/v1.0/groups/")
	if err != nil {
		log.Printf("ERROR: Failure fetching Brandmeister groups")
	}
	defer resp.Body.Close()
}
