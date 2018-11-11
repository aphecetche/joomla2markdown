// Package toto contains the types for schema 'jlabo'.
package toto

// Code generated by xo. DO NOT EDIT.

import (
	"database/sql"
	"errors"
)

// WsubUpdateSite represents a row from 'jlabo.wsub_update_sites'.
type WsubUpdateSite struct {
	UpdateSiteID       int            `json:"update_site_id"`       // update_site_id
	Name               sql.NullString `json:"name"`                 // name
	Type               sql.NullString `json:"type"`                 // type
	Location           string         `json:"location"`             // location
	Enabled            sql.NullInt64  `json:"enabled"`              // enabled
	LastCheckTimestamp sql.NullInt64  `json:"last_check_timestamp"` // last_check_timestamp
	ExtraQuery         sql.NullString `json:"extra_query"`          // extra_query

	// xo fields
	_exists, _deleted bool
}

// Exists determines if the WsubUpdateSite exists in the database.
func (wus *WsubUpdateSite) Exists() bool {
	return wus._exists
}

// Deleted provides information if the WsubUpdateSite has been deleted from the database.
func (wus *WsubUpdateSite) Deleted() bool {
	return wus._deleted
}

// Insert inserts the WsubUpdateSite to the database.
func (wus *WsubUpdateSite) Insert(db XODB) error {
	var err error

	// if already exist, bail
	if wus._exists {
		return errors.New("insert failed: already exists")
	}

	// sql insert query, primary key provided by autoincrement
	const sqlstr = `INSERT INTO jlabo.wsub_update_sites (` +
		`name, type, location, enabled, last_check_timestamp, extra_query` +
		`) VALUES (` +
		`?, ?, ?, ?, ?, ?` +
		`)`

	// run query
	XOLog(sqlstr, wus.Name, wus.Type, wus.Location, wus.Enabled, wus.LastCheckTimestamp, wus.ExtraQuery)
	res, err := db.Exec(sqlstr, wus.Name, wus.Type, wus.Location, wus.Enabled, wus.LastCheckTimestamp, wus.ExtraQuery)
	if err != nil {
		return err
	}

	// retrieve id
	id, err := res.LastInsertId()
	if err != nil {
		return err
	}

	// set primary key and existence
	wus.UpdateSiteID = int(id)
	wus._exists = true

	return nil
}

// Update updates the WsubUpdateSite in the database.
func (wus *WsubUpdateSite) Update(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !wus._exists {
		return errors.New("update failed: does not exist")
	}

	// if deleted, bail
	if wus._deleted {
		return errors.New("update failed: marked for deletion")
	}

	// sql query
	const sqlstr = `UPDATE jlabo.wsub_update_sites SET ` +
		`name = ?, type = ?, location = ?, enabled = ?, last_check_timestamp = ?, extra_query = ?` +
		` WHERE update_site_id = ?`

	// run query
	XOLog(sqlstr, wus.Name, wus.Type, wus.Location, wus.Enabled, wus.LastCheckTimestamp, wus.ExtraQuery, wus.UpdateSiteID)
	_, err = db.Exec(sqlstr, wus.Name, wus.Type, wus.Location, wus.Enabled, wus.LastCheckTimestamp, wus.ExtraQuery, wus.UpdateSiteID)
	return err
}

// Save saves the WsubUpdateSite to the database.
func (wus *WsubUpdateSite) Save(db XODB) error {
	if wus.Exists() {
		return wus.Update(db)
	}

	return wus.Insert(db)
}

// Delete deletes the WsubUpdateSite from the database.
func (wus *WsubUpdateSite) Delete(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !wus._exists {
		return nil
	}

	// if deleted, bail
	if wus._deleted {
		return nil
	}

	// sql query
	const sqlstr = `DELETE FROM jlabo.wsub_update_sites WHERE update_site_id = ?`

	// run query
	XOLog(sqlstr, wus.UpdateSiteID)
	_, err = db.Exec(sqlstr, wus.UpdateSiteID)
	if err != nil {
		return err
	}

	// set deleted
	wus._deleted = true

	return nil
}

// WsubUpdateSiteByUpdateSiteID retrieves a row from 'jlabo.wsub_update_sites' as a WsubUpdateSite.
//
// Generated from index 'wsub_update_sites_update_site_id_pkey'.
func WsubUpdateSiteByUpdateSiteID(db XODB, updateSiteID int) (*WsubUpdateSite, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`update_site_id, name, type, location, enabled, last_check_timestamp, extra_query ` +
		`FROM jlabo.wsub_update_sites ` +
		`WHERE update_site_id = ?`

	// run query
	XOLog(sqlstr, updateSiteID)
	wus := WsubUpdateSite{
		_exists: true,
	}

	err = db.QueryRow(sqlstr, updateSiteID).Scan(&wus.UpdateSiteID, &wus.Name, &wus.Type, &wus.Location, &wus.Enabled, &wus.LastCheckTimestamp, &wus.ExtraQuery)
	if err != nil {
		return nil, err
	}

	return &wus, nil
}