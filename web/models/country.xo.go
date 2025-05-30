package models

// Code generated by xo. DO NOT EDIT.

import (
	"context"
	"time"
)

// Country represents a row from 'sakila.country'.
type Country struct {
	CountryID  uint16    `json:"country_id"`  // country_id
	Country    string    `json:"country"`     // country
	LastUpdate time.Time `json:"last_update"` // last_update
	// xo fields
	_exists, _deleted bool
}

// Exists returns true when the [Country] exists in the database.
func (c *Country) Exists() bool {
	return c._exists
}

// Deleted returns true when the [Country] has been marked for deletion
// from the database.
func (c *Country) Deleted() bool {
	return c._deleted
}

// Insert inserts the [Country] to the database.
func (c *Country) Insert(ctx context.Context, db DB) error {
	switch {
	case c._exists: // already exists
		return logerror(&ErrInsertFailed{ErrAlreadyExists})
	case c._deleted: // deleted
		return logerror(&ErrInsertFailed{ErrMarkedForDeletion})
	}
	// insert (primary key generated and returned by database)
	const sqlstr = `INSERT INTO sakila.country (` +
		`country, last_update` +
		`) VALUES (` +
		`?, ?` +
		`)`
	// run
	logf(sqlstr, c.Country, c.LastUpdate)
	res, err := db.ExecContext(ctx, sqlstr, c.Country, c.LastUpdate)
	if err != nil {
		return logerror(err)
	}
	// retrieve id
	id, err := res.LastInsertId()
	if err != nil {
		return logerror(err)
	} // set primary key
	c.CountryID = uint16(id)
	// set exists
	c._exists = true
	return nil
}

// Update updates a [Country] in the database.
func (c *Country) Update(ctx context.Context, db DB) error {
	switch {
	case !c._exists: // doesn't exist
		return logerror(&ErrUpdateFailed{ErrDoesNotExist})
	case c._deleted: // deleted
		return logerror(&ErrUpdateFailed{ErrMarkedForDeletion})
	}
	// update with primary key
	const sqlstr = `UPDATE sakila.country SET ` +
		`country = ?, last_update = ? ` +
		`WHERE country_id = ?`
	// run
	logf(sqlstr, c.Country, c.LastUpdate, c.CountryID)
	if _, err := db.ExecContext(ctx, sqlstr, c.Country, c.LastUpdate, c.CountryID); err != nil {
		return logerror(err)
	}
	return nil
}

// Save saves the [Country] to the database.
func (c *Country) Save(ctx context.Context, db DB) error {
	if c.Exists() {
		return c.Update(ctx, db)
	}
	return c.Insert(ctx, db)
}

// Upsert performs an upsert for [Country].
func (c *Country) Upsert(ctx context.Context, db DB) error {
	switch {
	case c._deleted: // deleted
		return logerror(&ErrUpsertFailed{ErrMarkedForDeletion})
	}
	// upsert
	const sqlstr = `INSERT INTO sakila.country (` +
		`country_id, country, last_update` +
		`) VALUES (` +
		`?, ?, ?` +
		`)` +
		` ON DUPLICATE KEY UPDATE ` +
		`country = VALUES(country), last_update = VALUES(last_update)`
	// run
	logf(sqlstr, c.CountryID, c.Country, c.LastUpdate)
	if _, err := db.ExecContext(ctx, sqlstr, c.CountryID, c.Country, c.LastUpdate); err != nil {
		return logerror(err)
	}
	// set exists
	c._exists = true
	return nil
}

// Delete deletes the [Country] from the database.
func (c *Country) Delete(ctx context.Context, db DB) error {
	switch {
	case !c._exists: // doesn't exist
		return nil
	case c._deleted: // deleted
		return nil
	}
	// delete with single primary key
	const sqlstr = `DELETE FROM sakila.country ` +
		`WHERE country_id = ?`
	// run
	logf(sqlstr, c.CountryID)
	if _, err := db.ExecContext(ctx, sqlstr, c.CountryID); err != nil {
		return logerror(err)
	}
	// set deleted
	c._deleted = true
	return nil
}

// CountryByCountryID retrieves a row from 'sakila.country' as a [Country].
//
// Generated from index 'country_country_id_pkey'.
func CountryByCountryID(ctx context.Context, db DB, countryID uint16) (*Country, error) {
	// query
	const sqlstr = `SELECT ` +
		`country_id, country, last_update ` +
		`FROM sakila.country ` +
		`WHERE country_id = ?`
	// run
	logf(sqlstr, countryID)
	c := Country{
		_exists: true,
	}
	if err := db.QueryRowContext(ctx, sqlstr, countryID).Scan(&c.CountryID, &c.Country, &c.LastUpdate); err != nil {
		return nil, logerror(err)
	}
	return &c, nil
}
