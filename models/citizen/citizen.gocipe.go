// generated by gocipe fde8287dc8147e321eb245af512f5e1d059b75b9c464933b1ceb87a401fd8de6; DO NOT EDIT

package citizen

import (
	"time"

	"github.com/fluxynet/gocipe-example/models/country"
)

// Citizen A human being belonging to a country
type Citizen struct {
	ID         *string          `json:"id"`
	Surname    *string          `json:"surname"`
	OtherNames *string          `json:"othernames"`
	Gender     *string          `json:"gender"`
	DOB        *time.Time       `json:"dob"`
	Country    *country.Country `json:"country"`
}

// New returns an instance of Citizen
func New() *Citizen {
	return &Citizen{
		ID:         new(string),
		Surname:    new(string),
		OtherNames: new(string),
		Gender:     new(string),
		DOB:        new(time.Time),
	}
}
