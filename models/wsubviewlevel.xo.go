// Package toto contains the types for schema 'jlabo'.
package toto

// Code generated by xo. DO NOT EDIT.

import (
	"errors"
)

// WsubViewlevel represents a row from 'jlabo.wsub_viewlevels'.
type WsubViewlevel struct {
	ID       uint   `json:"id"`       // id
	Title    string `json:"title"`    // title
	Ordering int    `json:"ordering"` // ordering
	Rules    string `json:"rules"`    // rules

	// xo fields
	_exists, _deleted bool
}

// Exists determines if the WsubViewlevel exists in the database.
func (wv *WsubViewlevel) Exists() bool {
	return wv._exists
}

// Deleted provides information if the WsubViewlevel has been deleted from the database.
func (wv *WsubViewlevel) Deleted() bool {
	return wv._deleted
}

// Insert inserts the WsubViewlevel to the database.
func (wv *WsubViewlevel) Insert(db XODB) error {
	var err error

	// if already exist, bail
	if wv._exists {
		return errors.New("insert failed: already exists")
	}

	// sql insert query, primary key provided by autoincrement
	const sqlstr = `INSERT INTO jlabo.wsub_viewlevels (` +
		`title, ordering, rules` +
		`) VALUES (` +
		`?, ?, ?` +
		`)`

	// run query
	XOLog(sqlstr, wv.Title, wv.Ordering, wv.Rules)
	res, err := db.Exec(sqlstr, wv.Title, wv.Ordering, wv.Rules)
	if err != nil {
		return err
	}

	// retrieve id
	id, err := res.LastInsertId()
	if err != nil {
		return err
	}

	// set primary key and existence
	wv.ID = uint(id)
	wv._exists = true

	return nil
}

// Update updates the WsubViewlevel in the database.
func (wv *WsubViewlevel) Update(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !wv._exists {
		return errors.New("update failed: does not exist")
	}

	// if deleted, bail
	if wv._deleted {
		return errors.New("update failed: marked for deletion")
	}

	// sql query
	const sqlstr = `UPDATE jlabo.wsub_viewlevels SET ` +
		`title = ?, ordering = ?, rules = ?` +
		` WHERE id = ?`

	// run query
	XOLog(sqlstr, wv.Title, wv.Ordering, wv.Rules, wv.ID)
	_, err = db.Exec(sqlstr, wv.Title, wv.Ordering, wv.Rules, wv.ID)
	return err
}

// Save saves the WsubViewlevel to the database.
func (wv *WsubViewlevel) Save(db XODB) error {
	if wv.Exists() {
		return wv.Update(db)
	}

	return wv.Insert(db)
}

// Delete deletes the WsubViewlevel from the database.
func (wv *WsubViewlevel) Delete(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !wv._exists {
		return nil
	}

	// if deleted, bail
	if wv._deleted {
		return nil
	}

	// sql query
	const sqlstr = `DELETE FROM jlabo.wsub_viewlevels WHERE id = ?`

	// run query
	XOLog(sqlstr, wv.ID)
	_, err = db.Exec(sqlstr, wv.ID)
	if err != nil {
		return err
	}

	// set deleted
	wv._deleted = true

	return nil
}

// WsubViewlevelByTitle retrieves a row from 'jlabo.wsub_viewlevels' as a WsubViewlevel.
//
// Generated from index 'idx_assetgroup_title_lookup'.
func WsubViewlevelByTitle(db XODB, title string) (*WsubViewlevel, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`id, title, ordering, rules ` +
		`FROM jlabo.wsub_viewlevels ` +
		`WHERE title = ?`

	// run query
	XOLog(sqlstr, title)
	wv := WsubViewlevel{
		_exists: true,
	}

	err = db.QueryRow(sqlstr, title).Scan(&wv.ID, &wv.Title, &wv.Ordering, &wv.Rules)
	if err != nil {
		return nil, err
	}

	return &wv, nil
}

// WsubViewlevelByID retrieves a row from 'jlabo.wsub_viewlevels' as a WsubViewlevel.
//
// Generated from index 'wsub_viewlevels_id_pkey'.
func WsubViewlevelByID(db XODB, id uint) (*WsubViewlevel, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`id, title, ordering, rules ` +
		`FROM jlabo.wsub_viewlevels ` +
		`WHERE id = ?`

	// run query
	XOLog(sqlstr, id)
	wv := WsubViewlevel{
		_exists: true,
	}

	err = db.QueryRow(sqlstr, id).Scan(&wv.ID, &wv.Title, &wv.Ordering, &wv.Rules)
	if err != nil {
		return nil, err
	}

	return &wv, nil
}