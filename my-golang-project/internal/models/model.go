package models

import "time"

type Page struct {
	Title  string
	Data   interface{}
	Footer Footer
}

type Footer struct {
	EshtablishedYear     int
	EshtablishedYearUrdu int
	CopyRightStartYear   int
	Mobile               string
	PrimaryEmail         string
	SecondaryEmail       string
	Address              string
	AddressUrdu          string
	About                string
	AboutUrdu            string
}

func (f Footer) GetYear() int {
	// Get the current year
	return time.Now().Year()
}

type Activity struct {
	ID              string
	Description     string
	DescriptionUrdu string
	Images          []string
}

type ActivityPageData struct {
	Activities []Activity
}

type ContactDetails struct {
	Title          map[string]string
	Address        map[string]string
	PrimaryEmail   string
	SecondaryEmail string
	Phone          string
	FacebookURL    string
	TwitterURL     string
	YouTubeURL     string
}
