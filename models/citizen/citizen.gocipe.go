// generated by gocipe 2cdecd62f8cc1e3e1b51a96e61f1519e5d31bc3eb806037239eeb7707acf58b5; DO NOT EDIT

package citizen

import (
	"time"

	"github.com/fluxynet/gocipe-example/models/country"
)

// Citizen A human being belonging to a country
type Citizen struct {
	ID         *string          `json:"id"`
	Surname    *string          `json:""`
	OtherNames *string          `json:""`
	Gender     *string          `json:""`
	DOB        *time.Time       `json:""`
	Country    *country.Country `json:""`
}

// New returns an instance of Citizen
func New() *Citizen {
	return &Citizen{
		ID:         new(string),
		Surname:    new(string),
		OtherNames: new(string),
		Gender:     new(string),
		DOB:        new(time.Time),
		Country:    new(country.Country),
	}
}
