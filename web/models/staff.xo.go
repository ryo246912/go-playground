package models

// Code generated by xo. DO NOT EDIT.

import (
	"context"
	"database/sql"
	"time"
)

// Staff represents a row from 'sakila.staff'.
type Staff struct {
	StaffID    uint8          `json:"staff_id"`    // staff_id
	FirstName  string         `json:"first_name"`  // first_name
	LastName   string         `json:"last_name"`   // last_name
	AddressID  uint16         `json:"address_id"`  // address_id
	Picture    []byte         `json:"picture"`     // picture
	Email      sql.NullString `json:"email"`       // email
	StoreID    uint8          `json:"store_id"`    // store_id
	Active     bool           `json:"active"`      // active
	Username   string         `json:"username"`    // username
	Password   sql.NullString `json:"password"`    // password
	LastUpdate time.Time      `json:"last_update"` // last_update
	// xo fields
	_exists, _deleted bool
}

// Exists returns true when the [Staff] exists in the database.
func (s *Staff) Exists() bool {
	return s._exists
}

// Deleted returns true when the [Staff] has been marked for deletion
// from the database.
func (s *Staff) Deleted() bool {
	return s._deleted
}

// Insert inserts the [Staff] to the database.
func (s *Staff) Insert(ctx context.Context, db DB) error {
	switch {
	case s._exists: // already exists
		return logerror(&ErrInsertFailed{ErrAlreadyExists})
	case s._deleted: // deleted
		return logerror(&ErrInsertFailed{ErrMarkedForDeletion})
	}
	// insert (primary key generated and returned by database)
	const sqlstr = `INSERT INTO sakila.staff (` +
		`first_name, last_name, address_id, picture, email, store_id, active, username, password, last_update` +
		`) VALUES (` +
		`?, ?, ?, ?, ?, ?, ?, ?, ?, ?` +
		`)`
	// run
	logf(sqlstr, s.FirstName, s.LastName, s.AddressID, s.Picture, s.Email, s.StoreID, s.Active, s.Username, s.Password, s.LastUpdate)
	res, err := db.ExecContext(ctx, sqlstr, s.FirstName, s.LastName, s.AddressID, s.Picture, s.Email, s.StoreID, s.Active, s.Username, s.Password, s.LastUpdate)
	if err != nil {
		return logerror(err)
	}
	// retrieve id
	id, err := res.LastInsertId()
	if err != nil {
		return logerror(err)
	} // set primary key
	s.StaffID = uint8(id)
	// set exists
	s._exists = true
	return nil
}

// Update updates a [Staff] in the database.
func (s *Staff) Update(ctx context.Context, db DB) error {
	switch {
	case !s._exists: // doesn't exist
		return logerror(&ErrUpdateFailed{ErrDoesNotExist})
	case s._deleted: // deleted
		return logerror(&ErrUpdateFailed{ErrMarkedForDeletion})
	}
	// update with primary key
	const sqlstr = `UPDATE sakila.staff SET ` +
		`first_name = ?, last_name = ?, address_id = ?, picture = ?, email = ?, store_id = ?, active = ?, username = ?, password = ?, last_update = ? ` +
		`WHERE staff_id = ?`
	// run
	logf(sqlstr, s.FirstName, s.LastName, s.AddressID, s.Picture, s.Email, s.StoreID, s.Active, s.Username, s.Password, s.LastUpdate, s.StaffID)
	if _, err := db.ExecContext(ctx, sqlstr, s.FirstName, s.LastName, s.AddressID, s.Picture, s.Email, s.StoreID, s.Active, s.Username, s.Password, s.LastUpdate, s.StaffID); err != nil {
		return logerror(err)
	}
	return nil
}

// Save saves the [Staff] to the database.
func (s *Staff) Save(ctx context.Context, db DB) error {
	if s.Exists() {
		return s.Update(ctx, db)
	}
	return s.Insert(ctx, db)
}

// Upsert performs an upsert for [Staff].
func (s *Staff) Upsert(ctx context.Context, db DB) error {
	switch {
	case s._deleted: // deleted
		return logerror(&ErrUpsertFailed{ErrMarkedForDeletion})
	}
	// upsert
	const sqlstr = `INSERT INTO sakila.staff (` +
		`staff_id, first_name, last_name, address_id, picture, email, store_id, active, username, password, last_update` +
		`) VALUES (` +
		`?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?` +
		`)` +
		` ON DUPLICATE KEY UPDATE ` +
		`first_name = VALUES(first_name), last_name = VALUES(last_name), address_id = VALUES(address_id), picture = VALUES(picture), email = VALUES(email), store_id = VALUES(store_id), active = VALUES(active), username = VALUES(username), password = VALUES(password), last_update = VALUES(last_update)`
	// run
	logf(sqlstr, s.StaffID, s.FirstName, s.LastName, s.AddressID, s.Picture, s.Email, s.StoreID, s.Active, s.Username, s.Password, s.LastUpdate)
	if _, err := db.ExecContext(ctx, sqlstr, s.StaffID, s.FirstName, s.LastName, s.AddressID, s.Picture, s.Email, s.StoreID, s.Active, s.Username, s.Password, s.LastUpdate); err != nil {
		return logerror(err)
	}
	// set exists
	s._exists = true
	return nil
}

// Delete deletes the [Staff] from the database.
func (s *Staff) Delete(ctx context.Context, db DB) error {
	switch {
	case !s._exists: // doesn't exist
		return nil
	case s._deleted: // deleted
		return nil
	}
	// delete with single primary key
	const sqlstr = `DELETE FROM sakila.staff ` +
		`WHERE staff_id = ?`
	// run
	logf(sqlstr, s.StaffID)
	if _, err := db.ExecContext(ctx, sqlstr, s.StaffID); err != nil {
		return logerror(err)
	}
	// set deleted
	s._deleted = true
	return nil
}

// StaffByAddressID retrieves a row from 'sakila.staff' as a [Staff].
//
// Generated from index 'idx_fk_address_id'.
func StaffByAddressID(ctx context.Context, db DB, addressID uint16) ([]*Staff, error) {
	// query
	const sqlstr = `SELECT ` +
		`staff_id, first_name, last_name, address_id, picture, email, store_id, active, username, password, last_update ` +
		`FROM sakila.staff ` +
		`WHERE address_id = ?`
	// run
	logf(sqlstr, addressID)
	rows, err := db.QueryContext(ctx, sqlstr, addressID)
	if err != nil {
		return nil, logerror(err)
	}
	defer rows.Close()
	// process
	var res []*Staff
	for rows.Next() {
		s := Staff{
			_exists: true,
		}
		// scan
		if err := rows.Scan(&s.StaffID, &s.FirstName, &s.LastName, &s.AddressID, &s.Picture, &s.Email, &s.StoreID, &s.Active, &s.Username, &s.Password, &s.LastUpdate); err != nil {
			return nil, logerror(err)
		}
		res = append(res, &s)
	}
	if err := rows.Err(); err != nil {
		return nil, logerror(err)
	}
	return res, nil
}

// StaffByStoreID retrieves a row from 'sakila.staff' as a [Staff].
//
// Generated from index 'idx_fk_store_id'.
func StaffByStoreID(ctx context.Context, db DB, storeID uint8) ([]*Staff, error) {
	// query
	const sqlstr = `SELECT ` +
		`staff_id, first_name, last_name, address_id, picture, email, store_id, active, username, password, last_update ` +
		`FROM sakila.staff ` +
		`WHERE store_id = ?`
	// run
	logf(sqlstr, storeID)
	rows, err := db.QueryContext(ctx, sqlstr, storeID)
	if err != nil {
		return nil, logerror(err)
	}
	defer rows.Close()
	// process
	var res []*Staff
	for rows.Next() {
		s := Staff{
			_exists: true,
		}
		// scan
		if err := rows.Scan(&s.StaffID, &s.FirstName, &s.LastName, &s.AddressID, &s.Picture, &s.Email, &s.StoreID, &s.Active, &s.Username, &s.Password, &s.LastUpdate); err != nil {
			return nil, logerror(err)
		}
		res = append(res, &s)
	}
	if err := rows.Err(); err != nil {
		return nil, logerror(err)
	}
	return res, nil
}

// StaffByStaffID retrieves a row from 'sakila.staff' as a [Staff].
//
// Generated from index 'staff_staff_id_pkey'.
func StaffByStaffID(ctx context.Context, db DB, staffID uint8) (*Staff, error) {
	// query
	const sqlstr = `SELECT ` +
		`staff_id, first_name, last_name, address_id, picture, email, store_id, active, username, password, last_update ` +
		`FROM sakila.staff ` +
		`WHERE staff_id = ?`
	// run
	logf(sqlstr, staffID)
	s := Staff{
		_exists: true,
	}
	if err := db.QueryRowContext(ctx, sqlstr, staffID).Scan(&s.StaffID, &s.FirstName, &s.LastName, &s.AddressID, &s.Picture, &s.Email, &s.StoreID, &s.Active, &s.Username, &s.Password, &s.LastUpdate); err != nil {
		return nil, logerror(err)
	}
	return &s, nil
}

// Address returns the Address associated with the [Staff]'s (AddressID).
//
// Generated from foreign key 'fk_staff_address'.
func (s *Staff) Address(ctx context.Context, db DB) (*Address, error) {
	return AddressByAddressID(ctx, db, s.AddressID)
}

// Store returns the Store associated with the [Staff]'s (StoreID).
//
// Generated from foreign key 'fk_staff_store'.
func (s *Staff) Store(ctx context.Context, db DB) (*Store, error) {
	return StoreByStoreID(ctx, db, s.StoreID)
}
