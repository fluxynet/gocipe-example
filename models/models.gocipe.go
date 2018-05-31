// generated by gocipe 02fef3d117f1029d4142b6b7ae2d1ea6f313fd8a2f44e25333775a308c8afb37; DO NOT EDIT

package models

import "database/sql"

var (
	db *sql.DB

	CapitalRepo CapitalRepository
	CitizenRepo CitizenRepository
	CountryRepo CountryRepository
	TagRepo     TagRepository
)

const (
	OperationMerge  byte = 'M'
	OperationInsert byte = 'I'
	OperationUpdate byte = 'U'
)

// Init is responsible to initialize all repositories
func Init(database *sql.DB) {
	db = database

	CapitalRepo = CapitalRepository{db: database}
	CitizenRepo = CitizenRepository{db: database}
	CountryRepo = CountryRepository{db: database}
	TagRepo = TagRepository{db: database}
}

// StartTransaction initiates a database transaction
func StartTransaction() (*sql.Tx, error) {
	return db.Begin()
}

//ListFilter represents a filter to apply during listing (crud)
type ListFilter struct {
	Field     string
	Operation string
	Value     interface{}
}

// CapitalRepository encapsulates operations that may be performed on the entity Capital
type CapitalRepository struct {
	db *sql.DB
}

// CitizenRepository encapsulates operations that may be performed on the entity Citizen
type CitizenRepository struct {
	db *sql.DB
}

// CountryRepository encapsulates operations that may be performed on the entity Country
type CountryRepository struct {
	db *sql.DB
}

// TagRepository encapsulates operations that may be performed on the entity Tag
type TagRepository struct {
	db *sql.DB
}
