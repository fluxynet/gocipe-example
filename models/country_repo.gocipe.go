// generated by gocipe 02fef3d117f1029d4142b6b7ae2d1ea6f313fd8a2f44e25333775a308c8afb37; DO NOT EDIT

package models

import (
	"context"
	"database/sql"
	"strconv"
	"strings"

	"github.com/fluxynet/gocipe-example/util"
	"github.com/gobuffalo/uuid"
)

// Get returns a single Country from database by primary key
func (repo CountryRepository) Get(ctx context.Context, id string) (Country, error) {
	var (
		rows   *sql.Rows
		err    error
		entity Country
	)

	if err = util.CheckContext(ctx); err != nil {
		return entity, err
	}

	rows, err = repo.db.Query("SELECT id, name, continent FROM countries WHERE id = $1 ORDER BY id ASC", id)
	if err != nil {
		return entity, err
	}

	defer rows.Close()
	if rows.Next() {
		if err = util.CheckContext(ctx); err != nil {
			return entity, err
		}

		err = rows.Scan(&entity.ID, &entity.Name, &entity.Continent)
		if err != nil {
			return entity, err
		}

	}

	return entity, nil
}

// List returns a slice containing Country records
func (repo CountryRepository) List(ctx context.Context, filters []ListFilter, offset, limit int) ([]Country, error) {
	var (
		list     []Country
		segments []string
		values   []interface{}
		err      error
		rows     *sql.Rows
	)

	query := "SELECT id, name, continent FROM countries"

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

		var entity Country

		err = rows.Scan(entity.ID, &entity.Name, &entity.Continent)
		if err != nil {
			return nil, err
		}

		list = append(list, entity)
	}

	return list, nil
}

// Delete deletes a Country record from database and sets id to nil
func (repo CountryRepository) Delete(ctx context.Context, entity Country, tx *sql.Tx, autocommit bool) (Country, error) {
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

	stmt, err = tx.Prepare("DELETE FROM countries WHERE id = $1")
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

// DeleteMany deletes many Country records from database using filter
func (repo CountryRepository) DeleteMany(ctx context.Context, filters []ListFilter, tx *sql.Tx, autocommit bool) error {
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

	query := "DELETE FROM countries"

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

// Save either inserts or updates a Country record based on whether or not id is nil
func (repo CountryRepository) Save(ctx context.Context, entity Country, tx *sql.Tx, autocommit bool) (Country, error) {
	if entity.ID == "" {
		return CountryRepo.Insert(ctx, entity, tx, autocommit)
	}
	return CountryRepo.Update(ctx, entity, tx, autocommit)
}

// Insert performs an SQL insert for Country record and update instance with inserted id.
func (repo CountryRepository) Insert(ctx context.Context, entity Country, tx *sql.Tx, autocommit bool) (Country, error) {
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

	stmt, err = tx.Prepare("INSERT INTO countries (id, name, continent) VALUES ($1, $2, $3)")
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

	_, err = stmt.Exec(entity.ID, entity.Name, entity.Continent)
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

// Update Will execute an SQLUpdate Statement for Country in the database. Prefer using Save instead of Update directly.
func (repo CountryRepository) Update(ctx context.Context, entity Country, tx *sql.Tx, autocommit bool) (Country, error) {
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

	stmt, err = tx.Prepare("UPDATE countries SET name = $1, continent = $2 WHERE id = $1")
	if err != nil {
		return entity, err
	}

	if err = util.CheckContext(ctx); err != nil {
		tx.Rollback()
		return entity, err
	}
	_, err = stmt.Exec(entity.Name, entity.Continent)
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

// Merge performs an SQL merge for Country record.
func (repo CountryRepository) Merge(ctx context.Context, entity Country, tx *sql.Tx, autocommit bool) (Country, error) {
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
		return CountryRepo.Insert(ctx, entity, tx, autocommit)
	}

	stmt, err = tx.Prepare(`INSERT INTO countries (id, name, continent) VALUES ($1, $2, $3) 
	ON CONFLICT (id) DO UPDATE SET name = $2, continent = $3`)
	if err != nil {
		return entity, err
	}

	if err = util.CheckContext(ctx); err != nil {
		return entity, err
	}

	if err = util.CheckContext(ctx); err != nil {
		return entity, err
	}
	_, err = stmt.Exec(entity.ID, entity.Name, entity.Continent)
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

// LoadCitizens is a helper function to load related Citizen entities
func (repo CountryRepository) LoadCitizens(ctx context.Context, entities ...Country) error {
	var (
		err         error
		placeholder string
		values      []interface{}
		indices     = make(map[string][]*Country)
	)

	if err = util.CheckContext(ctx); err != nil {
		return err
	}

	c := 1
	for _, entity := range entities {
		placeholder += "$" + strconv.Itoa(c) + ","
		indices[entity.ID] = append(indices[entity.ID], &entity)
		values = append(values, entity.ID)
		c++
	}
	placeholder = strings.TrimRight(placeholder, ",")

	rows, err := repo.db.Query("SELECT country_id, id FROM citizens WHERE country_id IN ("+placeholder+")", values...)
	if err != nil {
		return err
	}

	if err = util.CheckContext(ctx); err != nil {
		return err
	}

	for rows.Next() {
		var (
			thisID string
			thatID string
		)
		err = rows.Scan(&thisID, &thatID)
		if err != nil {
			return err
		}

		for _, ent := range indices[thisID] {
			ent.Citizen = append(ent.Citizen, thatID)
		}

		if err = util.CheckContext(ctx); err != nil {
			return err
		}
	}

	return nil
} // LoadTags is a helper function to load related Tags entities
func (repo CountryRepository) LoadTags(ctx context.Context, entities ...Country) error {
	var (
		err         error
		placeholder string
		values      []interface{}
		indices     = make(map[string][]*Country)
	)

	if err = util.CheckContext(ctx); err != nil {
		return err
	}

	c := 1
	for _, entity := range entities {
		placeholder += "$" + strconv.Itoa(c) + ","
		indices[entity.ID] = append(indices[entity.ID], &entity)
		values = append(values, entity.ID)
		c++
	}
	placeholder = strings.TrimRight(placeholder, ",")

	rows, err := repo.db.Query(`
		SELECT j.country_id, t.id, t.name FROM tags t 
		INNER JOIN countries_tags j ON t.id = j.tag_id
		WHERE j.country_id IN (`+placeholder+`)
	`, values...)
	if err != nil {
		return err
	}

	if err = util.CheckContext(ctx); err != nil {
		return err
	}

	for rows.Next() {
		var (
			thisID     string
			entity     Country
			thatEntity Tag
		)

		err = rows.Scan(&thisID, entity.ID, &entity.Name)
		if err != nil {
			return err
		}

		for _, ent := range indices[thisID] {
			ent.Tags = append(ent.Tags, &thatEntity)
		}

		if err = util.CheckContext(ctx); err != nil {
			return err
		}
	}

	return nil
}
