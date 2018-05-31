// generated by gocipe 02fef3d117f1029d4142b6b7ae2d1ea6f313fd8a2f44e25333775a308c8afb37; DO NOT EDIT

package models

import (
	"context"
	"database/sql"
	"strconv"
	"strings"
	"time"

	"github.com/fluxynet/gocipe-example/util"
	"github.com/gobuffalo/uuid"
	"github.com/golang/protobuf/ptypes"
)

// Get returns a single Citizen from database by primary key
func (repo CitizenRepository) Get(ctx context.Context, id string) (Citizen, error) {
	var (
		rows   *sql.Rows
		err    error
		entity Citizen
	)

	if err = util.CheckContext(ctx); err != nil {
		return entity, err
	}

	rows, err = repo.db.Query("SELECT id, surname, othernames, gender, dob, country_id FROM citizens WHERE id = $1 ORDER BY id ASC", id)
	if err != nil {
		return entity, err
	}

	defer rows.Close()
	if rows.Next() {
		if err = util.CheckContext(ctx); err != nil {
			return entity, err
		}

		var dob time.Time

		err = rows.Scan(&entity.ID, &entity.Surname, &entity.OtherNames, &entity.Gender, &dob, &entity.CountryID)
		if err != nil {
			return entity, err
		}

		entity.DOB, _ = ptypes.TimestampProto(dob)

	}

	return entity, nil
}

// List returns a slice containing Citizen records
func (repo CitizenRepository) List(ctx context.Context, filters []ListFilter, offset, limit int) ([]Citizen, error) {
	var (
		list     []Citizen
		segments []string
		values   []interface{}
		err      error
		rows     *sql.Rows
	)

	query := "SELECT id, surname, othernames, gender, dob FROM citizens"

	if err = util.CheckContext(ctx); err != nil {
		return nil, err
	}

	for i, filter := range filters {
		segments = append(segments, filter.Field+" "+filter.Operation+" $"+strconv.Itoa(i+1))
		values = append(values, filter.Value)
	}

	if len(segments) != 0 {
		query += " WHERE " + strings.Join(segments, " AND ")
	}

	if limit > -1 {
		query += " LIMIT " + strconv.Itoa(limit)
	}

	if offset > -1 {
		query += " OFFSET " + strconv.Itoa(limit)
	}

	query += " ORDER BY id ASC"

	rows, err = repo.db.Query(query, values...)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	for rows.Next() {
		if err = util.CheckContext(ctx); err != nil {
			return nil, err
		}

		var entity Citizen
		var dob time.Time

		err = rows.Scan(entity.ID, &entity.Surname, &entity.OtherNames, &entity.Gender, &dob)
		if err != nil {
			return nil, err
		}

		entity.DOB, _ = ptypes.TimestampProto(dob)

		list = append(list, entity)
	}

	return list, nil
}

// Delete deletes a Citizen record from database and sets id to nil
func (repo CitizenRepository) Delete(ctx context.Context, entity Citizen, tx *sql.Tx, autocommit bool) (Citizen, error) {
	var (
		err  error
		stmt *sql.Stmt
	)
	id := entity.ID

	if tx == nil {
		if err = util.CheckContext(ctx); err != nil {
			return entity, err
		}

		tx, err = repo.db.Begin()
		if err != nil {
			return entity, err
		}
	}

	if err = util.CheckContext(ctx); err != nil {
		return entity, err
	}

	stmt, err = tx.Prepare("DELETE FROM citizens WHERE id = $1")
	if err != nil {
		return entity, err
	}

	if err = util.CheckContext(ctx); err != nil {
		tx.Rollback()
		return entity, err
	}

	_, err = stmt.Exec(id)
	if err == nil {
		entity.ID = ""
	} else {
		tx.Rollback()
		return entity, err
	}

	if err = util.CheckContext(ctx); err != nil {
		tx.Rollback()
		return entity, err
	}

	if autocommit {
		err = tx.Commit()
	}

	return entity, nil
}

// DeleteMany deletes many Citizen records from database using filter
func (repo CitizenRepository) DeleteMany(ctx context.Context, filters []ListFilter, tx *sql.Tx, autocommit bool) error {
	var (
		err      error
		stmt     *sql.Stmt
		segments []string
		values   []interface{}
	)

	if tx == nil {
		if err = util.CheckContext(ctx); err != nil {
			return err
		}

		tx, err = repo.db.Begin()
		if err != nil {
			return err
		}
	}

	query := "DELETE FROM citizens"

	if err = util.CheckContext(ctx); err != nil {
		tx.Rollback()
		return err
	}

	for i, filter := range filters {
		segments = append(segments, filter.Field+" "+filter.Operation+" $"+strconv.Itoa(i+1))
		values = append(values, filter.Value)
	}

	if len(segments) != 0 {
		query += " WHERE " + strings.Join(segments, " AND ")
	}

	stmt, err = repo.db.Prepare(query)
	if err != nil {
		return err
	}

	if err = util.CheckContext(ctx); err != nil {
		tx.Rollback()
		return err
	}

	_, err = stmt.Exec(values...)
	if err != nil {
		tx.Rollback()
		return err
	}

	if err = util.CheckContext(ctx); err != nil {
		tx.Rollback()
		return err
	}

	if autocommit {
		err = tx.Commit()
	}

	return err
}

// Save either inserts or updates a Citizen record based on whether or not id is nil
func (repo CitizenRepository) Save(ctx context.Context, entity Citizen, tx *sql.Tx, autocommit bool) (Citizen, error) {
	if entity.ID == "" {
		return CitizenRepo.Insert(ctx, entity, tx, autocommit)
	}
	return CitizenRepo.Update(ctx, entity, tx, autocommit)
}

// Insert performs an SQL insert for Citizen record and update instance with inserted id.
func (repo CitizenRepository) Insert(ctx context.Context, entity Citizen, tx *sql.Tx, autocommit bool) (Citizen, error) {
	var (
		id   string
		err  error
		stmt *sql.Stmt
	)

	if tx == nil {
		if err = util.CheckContext(ctx); err != nil {
			return entity, err
		}

		tx, err = repo.db.Begin()
		if err != nil {
			return entity, err
		}
	}
	dob, _ := ptypes.Timestamp(entity.DOB)

	stmt, err = tx.Prepare("INSERT INTO citizens (id, surname, othernames, gender, dob) VALUES ($1, $2, $3, $4, $5)")
	if err != nil {
		return entity, err
	}

	if err = util.CheckContext(ctx); err != nil {
		tx.Rollback()
		return entity, err
	}

	idUUID, err := uuid.NewV4()

	if err == nil {
		id = idUUID.String()
	} else {
		tx.Rollback()
		return entity, err
	}
	entity.ID = id

	if err = util.CheckContext(ctx); err != nil {
		tx.Rollback()
		return entity, err
	}

	_, err = stmt.Exec(entity.ID, entity.Surname, entity.OtherNames, entity.Gender, dob)
	if err != nil {
		tx.Rollback()
		return entity, err
	}

	if err = util.CheckContext(ctx); err != nil {
		tx.Rollback()
		return entity, err
	}

	if autocommit {
		err = tx.Commit()
	}

	return entity, nil
}

// Update Will execute an SQLUpdate Statement for Citizen in the database. Prefer using Save instead of Update directly.
func (repo CitizenRepository) Update(ctx context.Context, entity Citizen, tx *sql.Tx, autocommit bool) (Citizen, error) {
	var (
		err  error
		stmt *sql.Stmt
	)

	if tx == nil {
		if err = util.CheckContext(ctx); err != nil {
			return entity, err
		}

		tx, err = repo.db.Begin()
		if err != nil {
			return entity, err
		}
	}

	dob, _ := ptypes.Timestamp(entity.DOB)

	stmt, err = tx.Prepare("UPDATE citizens SET surname = $1, othernames = $2, gender = $3, dob = $4 WHERE id = $1")
	if err != nil {
		return entity, err
	}

	if err = util.CheckContext(ctx); err != nil {
		tx.Rollback()
		return entity, err
	}
	_, err = stmt.Exec(entity.Surname, entity.OtherNames, entity.Gender, dob)
	if err != nil {
		tx.Rollback()
		return entity, err
	}

	if err = util.CheckContext(ctx); err != nil {
		tx.Rollback()
		return entity, err
	}

	if autocommit {
		err = tx.Commit()
	}

	return entity, err
}

// Merge performs an SQL merge for Citizen record.
func (repo CitizenRepository) Merge(ctx context.Context, entity Citizen, tx *sql.Tx, autocommit bool) (Citizen, error) {
	var (
		err  error
		stmt *sql.Stmt
	)

	if tx == nil {
		if err = util.CheckContext(ctx); err != nil {
			return entity, err
		}

		tx, err = repo.db.Begin()
		if err != nil {
			return entity, err
		}
	}

	if entity.ID == "" {
		return CitizenRepo.Insert(ctx, entity, tx, autocommit)
	}

	dob, _ := ptypes.Timestamp(entity.DOB)

	stmt, err = tx.Prepare(`INSERT INTO citizens (id, surname, othernames, gender, dob) VALUES ($1, $2, $3, $4, $5) 
	ON CONFLICT (id) DO UPDATE SET surname = $2, othernames = $3, gender = $4, dob = $5`)
	if err != nil {
		return entity, err
	}

	if err = util.CheckContext(ctx); err != nil {
		return entity, err
	}

	if err = util.CheckContext(ctx); err != nil {
		return entity, err
	}
	_, err = stmt.Exec(entity.ID, entity.Surname, entity.OtherNames, entity.Gender, dob)
	if err != nil {
		tx.Rollback()
		return entity, err
	}

	if err = util.CheckContext(ctx); err != nil {
		return entity, err
	}

	if autocommit {
		err = tx.Commit()
	}

	return entity, err
}

// LoadCountries is a helper function to load related Country entities
func (repo CitizenRepository) LoadCountries(ctx context.Context, entities ...Citizen) error {
	var (
		err         error
		placeholder string
		values      []interface{}
		indices     = make(map[string][]*Citizen)
	)

	if err = util.CheckContext(ctx); err != nil {
		return err
	}

	c := 1
	for _, entity := range entities {
		placeholder += "$" + strconv.Itoa(c) + ","
		indices[entity.CountryID] = append(indices[entity.CountryID], &entity)
		values = append(values, entity.CountryID)
		c++
	}
	placeholder = strings.TrimRight(placeholder, ",")
	rows, err := repo.db.Query(`
		SELECT id, id, name, continent FROM countries WHERE id IN (`+placeholder+`)
	`, values...)
	if err != nil {
		return err
	}

	if err = util.CheckContext(ctx); err != nil {
		return err
	}

	for rows.Next() {
		var (
			thatID     string
			thatEntity Country
		)

		err = rows.Scan(&thatID, thatEntity.ID, &thatEntity.Name, &thatEntity.Continent)
		if err != nil {
			return err
		}

		for _, ent := range indices[thatID] {
			ent.Country = &thatEntity
		}

		if err = util.CheckContext(ctx); err != nil {
			return err
		}
	}

	return nil
}
