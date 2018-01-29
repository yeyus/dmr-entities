package model

type EntityTypeEnum string

const (
	EntityHam       EntityTypeEnum = "ham"
	EntityTalkgroup EntityTypeEnum = "talkgroup"
)

type SystemTypeEnum string

const (
	SystemTypeDMR SystemTypeEnum = "DMR"
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
