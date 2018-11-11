// Package toto contains the types for schema 'jlabo'.
package toto

// Code generated by xo. DO NOT EDIT.

import (
	"errors"
)

// WsubFinderTaxonomy represents a row from 'jlabo.wsub_finder_taxonomy'.
type WsubFinderTaxonomy struct {
	ID       uint   `json:"id"`        // id
	ParentID uint   `json:"parent_id"` // parent_id
	Title    string `json:"title"`     // title
	State    bool   `json:"state"`     // state
	Access   bool   `json:"access"`    // access
	Ordering bool   `json:"ordering"`  // ordering

	// xo fields
	_exists, _deleted bool
}

// Exists determines if the WsubFinderTaxonomy exists in the database.
func (wft *WsubFinderTaxonomy) Exists() bool {
	return wft._exists
}

// Deleted provides information if the WsubFinderTaxonomy has been deleted from the database.
func (wft *WsubFinderTaxonomy) Deleted() bool {
	return wft._deleted
}

// Insert inserts the WsubFinderTaxonomy to the database.
func (wft *WsubFinderTaxonomy) Insert(db XODB) error {
	var err error

	// if already exist, bail
	if wft._exists {
		return errors.New("insert failed: already exists")
	}

	// sql insert query, primary key provided by autoincrement
	const sqlstr = `INSERT INTO jlabo.wsub_finder_taxonomy (` +
		`parent_id, title, state, access, ordering` +
		`) VALUES (` +
		`?, ?, ?, ?, ?` +
		`)`

	// run query
	XOLog(sqlstr, wft.ParentID, wft.Title, wft.State, wft.Access, wft.Ordering)
	res, err := db.Exec(sqlstr, wft.ParentID, wft.Title, wft.State, wft.Access, wft.Ordering)
	if err != nil {
		return err
	}

	// retrieve id
	id, err := res.LastInsertId()
	if err != nil {
		return err
	}

	// set primary key and existence
	wft.ID = uint(id)
	wft._exists = true

	return nil
}

// Update updates the WsubFinderTaxonomy in the database.
func (wft *WsubFinderTaxonomy) Update(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !wft._exists {
		return errors.New("update failed: does not exist")
	}

	// if deleted, bail
	if wft._deleted {
		return errors.New("update failed: marked for deletion")
	}

	// sql query
	const sqlstr = `UPDATE jlabo.wsub_finder_taxonomy SET ` +
		`parent_id = ?, title = ?, state = ?, access = ?, ordering = ?` +
		` WHERE id = ?`

	// run query
	XOLog(sqlstr, wft.ParentID, wft.Title, wft.State, wft.Access, wft.Ordering, wft.ID)
	_, err = db.Exec(sqlstr, wft.ParentID, wft.Title, wft.State, wft.Access, wft.Ordering, wft.ID)
	return err
}

// Save saves the WsubFinderTaxonomy to the database.
func (wft *WsubFinderTaxonomy) Save(db XODB) error {
	if wft.Exists() {
		return wft.Update(db)
	}

	return wft.Insert(db)
}

// Delete deletes the WsubFinderTaxonomy from the database.
func (wft *WsubFinderTaxonomy) Delete(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !wft._exists {
		return nil
	}

	// if deleted, bail
	if wft._deleted {
		return nil
	}

	// sql query
	const sqlstr = `DELETE FROM jlabo.wsub_finder_taxonomy WHERE id = ?`

	// run query
	XOLog(sqlstr, wft.ID)
	_, err = db.Exec(sqlstr, wft.ID)
	if err != nil {
		return err
	}

	// set deleted
	wft._deleted = true

	return nil
}

// WsubFinderTaxonomiesByAccess retrieves a row from 'jlabo.wsub_finder_taxonomy' as a WsubFinderTaxonomy.
//
// Generated from index 'access'.
func WsubFinderTaxonomiesByAccess(db XODB, access bool) ([]*WsubFinderTaxonomy, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`id, parent_id, title, state, access, ordering ` +
		`FROM jlabo.wsub_finder_taxonomy ` +
		`WHERE access = ?`

	// run query
	XOLog(sqlstr, access)
	q, err := db.Query(sqlstr, access)
	if err != nil {
		return nil, err
	}
	defer q.Close()

	// load results
	res := []*WsubFinderTaxonomy{}
	for q.Next() {
		wft := WsubFinderTaxonomy{
			_exists: true,
		}

		// scan
		err = q.Scan(&wft.ID, &wft.ParentID, &wft.Title, &wft.State, &wft.Access, &wft.Ordering)
		if err != nil {
			return nil, err
		}

		res = append(res, &wft)
	}

	return res, nil
}

// WsubFinderTaxonomiesByParentIDStateAccess retrieves a row from 'jlabo.wsub_finder_taxonomy' as a WsubFinderTaxonomy.
//
// Generated from index 'idx_parent_published'.
func WsubFinderTaxonomiesByParentIDStateAccess(db XODB, parentID uint, state bool, access bool) ([]*WsubFinderTaxonomy, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`id, parent_id, title, state, access, ordering ` +
		`FROM jlabo.wsub_finder_taxonomy ` +
		`WHERE parent_id = ? AND state = ? AND access = ?`

	// run query
	XOLog(sqlstr, parentID, state, access)
	q, err := db.Query(sqlstr, parentID, state, access)
	if err != nil {
		return nil, err
	}
	defer q.Close()

	// load results
	res := []*WsubFinderTaxonomy{}
	for q.Next() {
		wft := WsubFinderTaxonomy{
			_exists: true,
		}

		// scan
		err = q.Scan(&wft.ID, &wft.ParentID, &wft.Title, &wft.State, &wft.Access, &wft.Ordering)
		if err != nil {
			return nil, err
		}

		res = append(res, &wft)
	}

	return res, nil
}

// WsubFinderTaxonomiesByOrdering retrieves a row from 'jlabo.wsub_finder_taxonomy' as a WsubFinderTaxonomy.
//
// Generated from index 'ordering'.
func WsubFinderTaxonomiesByOrdering(db XODB, ordering bool) ([]*WsubFinderTaxonomy, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`id, parent_id, title, state, access, ordering ` +
		`FROM jlabo.wsub_finder_taxonomy ` +
		`WHERE ordering = ?`

	// run query
	XOLog(sqlstr, ordering)
	q, err := db.Query(sqlstr, ordering)
	if err != nil {
		return nil, err
	}
	defer q.Close()

	// load results
	res := []*WsubFinderTaxonomy{}
	for q.Next() {
		wft := WsubFinderTaxonomy{
			_exists: true,
		}

		// scan
		err = q.Scan(&wft.ID, &wft.ParentID, &wft.Title, &wft.State, &wft.Access, &wft.Ordering)
		if err != nil {
			return nil, err
		}

		res = append(res, &wft)
	}

	return res, nil
}

// WsubFinderTaxonomiesByParentID retrieves a row from 'jlabo.wsub_finder_taxonomy' as a WsubFinderTaxonomy.
//
// Generated from index 'parent_id'.
func WsubFinderTaxonomiesByParentID(db XODB, parentID uint) ([]*WsubFinderTaxonomy, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`id, parent_id, title, state, access, ordering ` +
		`FROM jlabo.wsub_finder_taxonomy ` +
		`WHERE parent_id = ?`

	// run query
	XOLog(sqlstr, parentID)
	q, err := db.Query(sqlstr, parentID)
	if err != nil {
		return nil, err
	}
	defer q.Close()

	// load results
	res := []*WsubFinderTaxonomy{}
	for q.Next() {
		wft := WsubFinderTaxonomy{
			_exists: true,
		}

		// scan
		err = q.Scan(&wft.ID, &wft.ParentID, &wft.Title, &wft.State, &wft.Access, &wft.Ordering)
		if err != nil {
			return nil, err
		}

		res = append(res, &wft)
	}

	return res, nil
}

// WsubFinderTaxonomiesByState retrieves a row from 'jlabo.wsub_finder_taxonomy' as a WsubFinderTaxonomy.
//
// Generated from index 'state'.
func WsubFinderTaxonomiesByState(db XODB, state bool) ([]*WsubFinderTaxonomy, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`id, parent_id, title, state, access, ordering ` +
		`FROM jlabo.wsub_finder_taxonomy ` +
		`WHERE state = ?`

	// run query
	XOLog(sqlstr, state)
	q, err := db.Query(sqlstr, state)
	if err != nil {
		return nil, err
	}
	defer q.Close()

	// load results
	res := []*WsubFinderTaxonomy{}
	for q.Next() {
		wft := WsubFinderTaxonomy{
			_exists: true,
		}

		// scan
		err = q.Scan(&wft.ID, &wft.ParentID, &wft.Title, &wft.State, &wft.Access, &wft.Ordering)
		if err != nil {
			return nil, err
		}

		res = append(res, &wft)
	}

	return res, nil
}

// WsubFinderTaxonomyByID retrieves a row from 'jlabo.wsub_finder_taxonomy' as a WsubFinderTaxonomy.
//
// Generated from index 'wsub_finder_taxonomy_id_pkey'.
func WsubFinderTaxonomyByID(db XODB, id uint) (*WsubFinderTaxonomy, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`id, parent_id, title, state, access, ordering ` +
		`FROM jlabo.wsub_finder_taxonomy ` +
		`WHERE id = ?`

	// run query
	XOLog(sqlstr, id)
	wft := WsubFinderTaxonomy{
		_exists: true,
	}

	err = db.QueryRow(sqlstr, id).Scan(&wft.ID, &wft.ParentID, &wft.Title, &wft.State, &wft.Access, &wft.Ordering)
	if err != nil {
		return nil, err
	}

	return &wft, nil
}
