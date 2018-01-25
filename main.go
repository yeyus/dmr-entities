package main

import (
	"fmt"
	"github.com/yeyus/dmr-entities/brandmeister"
	//	"github.com/yeyus/dmr-entities/dmrmarc"
	//	"log"
)

type Entity struct {
	Id         uint32
	System     SystemTypeEnum
	Type       EntityTypeEnum
	Callsign   string
	Name       string
	Surname    string
	Country    string
	CountryISO string
	State      string
	City       string
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
	fmt.Printf("Groups: %v", brandmeister.GetBrandmeisterGroups())
}
