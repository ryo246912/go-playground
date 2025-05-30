package models

// Code generated by xo. DO NOT EDIT.

import (
	"context"
	"database/sql"
	"time"

	"github.com/uptrace/bun"
)

// Film represents a row from 'sakila.film'.
type Film struct {
	// モデル名とテーブル名を対応付け
	bun.BaseModel `bun:"table:film,alias:f"`

	FilmID             uint16         `json:"film_id" bun:"film_id,pk,autoincrement"`              // film_id
	Title              string         `json:"title"`                // title
	Description        sql.NullString `json:"description"`          // description
	ReleaseYear        sql.NullInt64  `json:"release_year"`         // release_year
	LanguageID         uint8          `json:"language_id"`          // language_id
	OriginalLanguageID sql.NullInt64  `json:"original_language_id"` // original_language_id
	RentalDuration     uint8          `json:"rental_duration"`      // rental_duration
	RentalRate         float64        `json:"rental_rate"`          // rental_rate
	Length             sql.NullInt64  `json:"length"`               // length
	ReplacementCost    float64        `json:"replacement_cost"`     // replacement_cost
	Rating             NullRating     `json:"rating"`               // rating
	SpecialFeatures    []byte         `json:"special_features"`     // special_features
	LastUpdate         time.Time      `json:"last_update"`          // last_update
	// xo fields
	_exists, _deleted bool
	// relation
	Categoris   []*FilmCategory  `bun:"rel:has-many,join:film_id=film_id"` // 1対多の関係
}

// Exists returns true when the [Film] exists in the database.
func (f *Film) Exists() bool {
	return f._exists
}

// Deleted returns true when the [Film] has been marked for deletion
// from the database.
func (f *Film) Deleted() bool {
	return f._deleted
}

// Insert inserts the [Film] to the database.
func (f *Film) Insert(ctx context.Context, db DB) error {
	switch {
	case f._exists: // already exists
		return logerror(&ErrInsertFailed{ErrAlreadyExists})
	case f._deleted: // deleted
		return logerror(&ErrInsertFailed{ErrMarkedForDeletion})
	}
	// insert (primary key generated and returned by database)
	const sqlstr = `INSERT INTO sakila.film (` +
		`title, description, release_year, language_id, original_language_id, rental_duration, rental_rate, length, replacement_cost, rating, special_features, last_update` +
		`) VALUES (` +
		`?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?` +
		`)`
	// run
	logf(sqlstr, f.Title, f.Description, f.ReleaseYear, f.LanguageID, f.OriginalLanguageID, f.RentalDuration, f.RentalRate, f.Length, f.ReplacementCost, f.Rating, f.SpecialFeatures, f.LastUpdate)
	res, err := db.ExecContext(ctx, sqlstr, f.Title, f.Description, f.ReleaseYear, f.LanguageID, f.OriginalLanguageID, f.RentalDuration, f.RentalRate, f.Length, f.ReplacementCost, f.Rating, f.SpecialFeatures, f.LastUpdate)
	if err != nil {
		return logerror(err)
	}
	// retrieve id
	id, err := res.LastInsertId()
	if err != nil {
		return logerror(err)
	} // set primary key
	f.FilmID = uint16(id)
	// set exists
	f._exists = true
	return nil
}

// Update updates a [Film] in the database.
func (f *Film) Update(ctx context.Context, db DB) error {
	switch {
	case !f._exists: // doesn't exist
		return logerror(&ErrUpdateFailed{ErrDoesNotExist})
	case f._deleted: // deleted
		return logerror(&ErrUpdateFailed{ErrMarkedForDeletion})
	}
	// update with primary key
	const sqlstr = `UPDATE sakila.film SET ` +
		`title = ?, description = ?, release_year = ?, language_id = ?, original_language_id = ?, rental_duration = ?, rental_rate = ?, length = ?, replacement_cost = ?, rating = ?, special_features = ?, last_update = ? ` +
		`WHERE film_id = ?`
	// run
	logf(sqlstr, f.Title, f.Description, f.ReleaseYear, f.LanguageID, f.OriginalLanguageID, f.RentalDuration, f.RentalRate, f.Length, f.ReplacementCost, f.Rating, f.SpecialFeatures, f.LastUpdate, f.FilmID)
	if _, err := db.ExecContext(ctx, sqlstr, f.Title, f.Description, f.ReleaseYear, f.LanguageID, f.OriginalLanguageID, f.RentalDuration, f.RentalRate, f.Length, f.ReplacementCost, f.Rating, f.SpecialFeatures, f.LastUpdate, f.FilmID); err != nil {
		return logerror(err)
	}
	return nil
}

// Save saves the [Film] to the database.
func (f *Film) Save(ctx context.Context, db DB) error {
	if f.Exists() {
		return f.Update(ctx, db)
	}
	return f.Insert(ctx, db)
}

// Upsert performs an upsert for [Film].
func (f *Film) Upsert(ctx context.Context, db DB) error {
	switch {
	case f._deleted: // deleted
		return logerror(&ErrUpsertFailed{ErrMarkedForDeletion})
	}
	// upsert
	const sqlstr = `INSERT INTO sakila.film (` +
		`film_id, title, description, release_year, language_id, original_language_id, rental_duration, rental_rate, length, replacement_cost, rating, special_features, last_update` +
		`) VALUES (` +
		`?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?` +
		`)` +
		` ON DUPLICATE KEY UPDATE ` +
		`title = VALUES(title), description = VALUES(description), release_year = VALUES(release_year), language_id = VALUES(language_id), original_language_id = VALUES(original_language_id), rental_duration = VALUES(rental_duration), rental_rate = VALUES(rental_rate), length = VALUES(length), replacement_cost = VALUES(replacement_cost), rating = VALUES(rating), special_features = VALUES(special_features), last_update = VALUES(last_update)`
	// run
	logf(sqlstr, f.FilmID, f.Title, f.Description, f.ReleaseYear, f.LanguageID, f.OriginalLanguageID, f.RentalDuration, f.RentalRate, f.Length, f.ReplacementCost, f.Rating, f.SpecialFeatures, f.LastUpdate)
	if _, err := db.ExecContext(ctx, sqlstr, f.FilmID, f.Title, f.Description, f.ReleaseYear, f.LanguageID, f.OriginalLanguageID, f.RentalDuration, f.RentalRate, f.Length, f.ReplacementCost, f.Rating, f.SpecialFeatures, f.LastUpdate); err != nil {
		return logerror(err)
	}
	// set exists
	f._exists = true
	return nil
}

// Delete deletes the [Film] from the database.
func (f *Film) Delete(ctx context.Context, db DB) error {
	switch {
	case !f._exists: // doesn't exist
		return nil
	case f._deleted: // deleted
		return nil
	}
	// delete with single primary key
	const sqlstr = `DELETE FROM sakila.film ` +
		`WHERE film_id = ?`
	// run
	logf(sqlstr, f.FilmID)
	if _, err := db.ExecContext(ctx, sqlstr, f.FilmID); err != nil {
		return logerror(err)
	}
	// set deleted
	f._deleted = true
	return nil
}

// FilmByFilmID retrieves a row from 'sakila.film' as a [Film].
//
// Generated from index 'film_film_id_pkey'.
func FilmByFilmID(ctx context.Context, db DB, filmID uint16) (*Film, error) {
	// query
	const sqlstr = `SELECT ` +
		`film_id, title, description, release_year, language_id, original_language_id, rental_duration, rental_rate, length, replacement_cost, rating, special_features, last_update ` +
		`FROM sakila.film ` +
		`WHERE film_id = ?`
	// run
	logf(sqlstr, filmID)
	f := Film{
		_exists: true,
	}
	if err := db.QueryRowContext(ctx, sqlstr, filmID).Scan(&f.FilmID, &f.Title, &f.Description, &f.ReleaseYear, &f.LanguageID, &f.OriginalLanguageID, &f.RentalDuration, &f.RentalRate, &f.Length, &f.ReplacementCost, &f.Rating, &f.SpecialFeatures, &f.LastUpdate); err != nil {
		return nil, logerror(err)
	}
	return &f, nil
}

// FilmByLanguageID retrieves a row from 'sakila.film' as a [Film].
//
// Generated from index 'idx_fk_language_id'.
func FilmByLanguageID(ctx context.Context, db DB, languageID uint8) ([]*Film, error) {
	// query
	const sqlstr = `SELECT ` +
		`film_id, title, description, release_year, language_id, original_language_id, rental_duration, rental_rate, length, replacement_cost, rating, special_features, last_update ` +
		`FROM sakila.film ` +
		`WHERE language_id = ?`
	// run
	logf(sqlstr, languageID)
	rows, err := db.QueryContext(ctx, sqlstr, languageID)
	if err != nil {
		return nil, logerror(err)
	}
	defer rows.Close()
	// process
	var res []*Film
	for rows.Next() {
		f := Film{
			_exists: true,
		}
		// scan
		if err := rows.Scan(&f.FilmID, &f.Title, &f.Description, &f.ReleaseYear, &f.LanguageID, &f.OriginalLanguageID, &f.RentalDuration, &f.RentalRate, &f.Length, &f.ReplacementCost, &f.Rating, &f.SpecialFeatures, &f.LastUpdate); err != nil {
			return nil, logerror(err)
		}
		res = append(res, &f)
	}
	if err := rows.Err(); err != nil {
		return nil, logerror(err)
	}
	return res, nil
}

// FilmByOriginalLanguageID retrieves a row from 'sakila.film' as a [Film].
//
// Generated from index 'idx_fk_original_language_id'.
func FilmByOriginalLanguageID(ctx context.Context, db DB, originalLanguageID sql.NullInt64) ([]*Film, error) {
	// query
	const sqlstr = `SELECT ` +
		`film_id, title, description, release_year, language_id, original_language_id, rental_duration, rental_rate, length, replacement_cost, rating, special_features, last_update ` +
		`FROM sakila.film ` +
		`WHERE original_language_id = ?`
	// run
	logf(sqlstr, originalLanguageID)
	rows, err := db.QueryContext(ctx, sqlstr, originalLanguageID)
	if err != nil {
		return nil, logerror(err)
	}
	defer rows.Close()
	// process
	var res []*Film
	for rows.Next() {
		f := Film{
			_exists: true,
		}
		// scan
		if err := rows.Scan(&f.FilmID, &f.Title, &f.Description, &f.ReleaseYear, &f.LanguageID, &f.OriginalLanguageID, &f.RentalDuration, &f.RentalRate, &f.Length, &f.ReplacementCost, &f.Rating, &f.SpecialFeatures, &f.LastUpdate); err != nil {
			return nil, logerror(err)
		}
		res = append(res, &f)
	}
	if err := rows.Err(); err != nil {
		return nil, logerror(err)
	}
	return res, nil
}

// FilmByTitle retrieves a row from 'sakila.film' as a [Film].
//
// Generated from index 'idx_title'.
func FilmByTitle(ctx context.Context, db DB, title string) ([]*Film, error) {
	// query
	const sqlstr = `SELECT ` +
		`film_id, title, description, release_year, language_id, original_language_id, rental_duration, rental_rate, length, replacement_cost, rating, special_features, last_update ` +
		`FROM sakila.film ` +
		`WHERE title = ?`
	// run
	logf(sqlstr, title)
	rows, err := db.QueryContext(ctx, sqlstr, title)
	if err != nil {
		return nil, logerror(err)
	}
	defer rows.Close()
	// process
	var res []*Film
	for rows.Next() {
		f := Film{
			_exists: true,
		}
		// scan
		if err := rows.Scan(&f.FilmID, &f.Title, &f.Description, &f.ReleaseYear, &f.LanguageID, &f.OriginalLanguageID, &f.RentalDuration, &f.RentalRate, &f.Length, &f.ReplacementCost, &f.Rating, &f.SpecialFeatures, &f.LastUpdate); err != nil {
			return nil, logerror(err)
		}
		res = append(res, &f)
	}
	if err := rows.Err(); err != nil {
		return nil, logerror(err)
	}
	return res, nil
}

// Language returns the Language associated with the [Film]'s (LanguageID).
//
// Generated from foreign key 'fk_film_language'.
func (f *Film) Language(ctx context.Context, db DB) (*Language, error) {
	return LanguageByLanguageID(ctx, db, f.LanguageID)
}

// Language returns the Language associated with the [Film]'s (OriginalLanguageID).
//
// Generated from foreign key 'fk_film_language_original'.
func (f *Film) OriginalLanguage(ctx context.Context, db DB) (*Language, error) {
	return LanguageByLanguageID(ctx, db, uint8(f.OriginalLanguageID.Int64))
}
