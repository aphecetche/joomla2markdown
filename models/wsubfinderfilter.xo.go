// Package toto contains the types for schema 'jlabo'.
package toto

// Code generated by xo. DO NOT EDIT.

import (
	"database/sql"
	"errors"
	"time"
)

// WsubFinderFilter represents a row from 'jlabo.wsub_finder_filters'.
type WsubFinderFilter struct {
	FilterID       uint           `json:"filter_id"`        // filter_id
	Title          string         `json:"title"`            // title
	Alias          string         `json:"alias"`            // alias
	State          bool           `json:"state"`            // state
	Created        time.Time      `json:"created"`          // created
	CreatedBy      uint           `json:"created_by"`       // created_by
	CreatedByAlias string         `json:"created_by_alias"` // created_by_alias
	Modified       time.Time      `json:"modified"`         // modified
	ModifiedBy     uint           `json:"modified_by"`      // modified_by
	CheckedOut     uint           `json:"checked_out"`      // checked_out
	CheckedOutTime time.Time      `json:"checked_out_time"` // checked_out_time
	MapCount       uint           `json:"map_count"`        // map_count
	Data           string         `json:"data"`             // data
	Params         sql.NullString `json:"params"`           // params

	// xo fields
	_exists, _deleted bool
}

// Exists determines if the WsubFinderFilter exists in the database.
func (wff *WsubFinderFilter) Exists() bool {
	return wff._exists
}

// Deleted provides information if the WsubFinderFilter has been deleted from the database.
func (wff *WsubFinderFilter) Deleted() bool {
	return wff._deleted
}

// Insert inserts the WsubFinderFilter to the database.
func (wff *WsubFinderFilter) Insert(db XODB) error {
	var err error

	// if already exist, bail
	if wff._exists {
		return errors.New("insert failed: already exists")
	}

	// sql insert query, primary key provided by autoincrement
	const sqlstr = `INSERT INTO jlabo.wsub_finder_filters (` +
		`title, alias, state, created, created_by, created_by_alias, modified, modified_by, checked_out, checked_out_time, map_count, data, params` +
		`) VALUES (` +
		`?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?` +
		`)`

	// run query
	XOLog(sqlstr, wff.Title, wff.Alias, wff.State, wff.Created, wff.CreatedBy, wff.CreatedByAlias, wff.Modified, wff.ModifiedBy, wff.CheckedOut, wff.CheckedOutTime, wff.MapCount, wff.Data, wff.Params)
	res, err := db.Exec(sqlstr, wff.Title, wff.Alias, wff.State, wff.Created, wff.CreatedBy, wff.CreatedByAlias, wff.Modified, wff.ModifiedBy, wff.CheckedOut, wff.CheckedOutTime, wff.MapCount, wff.Data, wff.Params)
	if err != nil {
		return err
	}

	// retrieve id
	id, err := res.LastInsertId()
	if err != nil {
		return err
	}

	// set primary key and existence
	wff.FilterID = uint(id)
	wff._exists = true

	return nil
}

// Update updates the WsubFinderFilter in the database.
func (wff *WsubFinderFilter) Update(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !wff._exists {
		return errors.New("update failed: does not exist")
	}

	// if deleted, bail
	if wff._deleted {
		return errors.New("update failed: marked for deletion")
	}

	// sql query
	const sqlstr = `UPDATE jlabo.wsub_finder_filters SET ` +
		`title = ?, alias = ?, state = ?, created = ?, created_by = ?, created_by_alias = ?, modified = ?, modified_by = ?, checked_out = ?, checked_out_time = ?, map_count = ?, data = ?, params = ?` +
		` WHERE filter_id = ?`

	// run query
	XOLog(sqlstr, wff.Title, wff.Alias, wff.State, wff.Created, wff.CreatedBy, wff.CreatedByAlias, wff.Modified, wff.ModifiedBy, wff.CheckedOut, wff.CheckedOutTime, wff.MapCount, wff.Data, wff.Params, wff.FilterID)
	_, err = db.Exec(sqlstr, wff.Title, wff.Alias, wff.State, wff.Created, wff.CreatedBy, wff.CreatedByAlias, wff.Modified, wff.ModifiedBy, wff.CheckedOut, wff.CheckedOutTime, wff.MapCount, wff.Data, wff.Params, wff.FilterID)
	return err
}

// Save saves the WsubFinderFilter to the database.
func (wff *WsubFinderFilter) Save(db XODB) error {
	if wff.Exists() {
		return wff.Update(db)
	}

	return wff.Insert(db)
}

// Delete deletes the WsubFinderFilter from the database.
func (wff *WsubFinderFilter) Delete(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !wff._exists {
		return nil
	}

	// if deleted, bail
	if wff._deleted {
		return nil
	}

	// sql query
	const sqlstr = `DELETE FROM jlabo.wsub_finder_filters WHERE filter_id = ?`

	// run query
	XOLog(sqlstr, wff.FilterID)
	_, err = db.Exec(sqlstr, wff.FilterID)
	if err != nil {
		return err
	}

	// set deleted
	wff._deleted = true

	return nil
}

// WsubFinderFilterByFilterID retrieves a row from 'jlabo.wsub_finder_filters' as a WsubFinderFilter.
//
// Generated from index 'wsub_finder_filters_filter_id_pkey'.
func WsubFinderFilterByFilterID(db XODB, filterID uint) (*WsubFinderFilter, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`filter_id, title, alias, state, created, created_by, created_by_alias, modified, modified_by, checked_out, checked_out_time, map_count, data, params ` +
		`FROM jlabo.wsub_finder_filters ` +
		`WHERE filter_id = ?`

	// run query
	XOLog(sqlstr, filterID)
	wff := WsubFinderFilter{
		_exists: true,
	}

	err = db.QueryRow(sqlstr, filterID).Scan(&wff.FilterID, &wff.Title, &wff.Alias, &wff.State, &wff.Created, &wff.CreatedBy, &wff.CreatedByAlias, &wff.Modified, &wff.ModifiedBy, &wff.CheckedOut, &wff.CheckedOutTime, &wff.MapCount, &wff.Data, &wff.Params)
	if err != nil {
		return nil, err
	}

	return &wff, nil
}
