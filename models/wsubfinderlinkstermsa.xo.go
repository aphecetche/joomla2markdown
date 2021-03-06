// Package toto contains the types for schema 'jlabo'.
package toto

// Code generated by xo. DO NOT EDIT.

import (
	"errors"
)

// WsubFinderLinksTermsa represents a row from 'jlabo.wsub_finder_links_termsa'.
type WsubFinderLinksTermsa struct {
	LinkID uint    `json:"link_id"` // link_id
	TermID uint    `json:"term_id"` // term_id
	Weight float32 `json:"weight"`  // weight

	// xo fields
	_exists, _deleted bool
}

// Exists determines if the WsubFinderLinksTermsa exists in the database.
func (wflt *WsubFinderLinksTermsa) Exists() bool {
	return wflt._exists
}

// Deleted provides information if the WsubFinderLinksTermsa has been deleted from the database.
func (wflt *WsubFinderLinksTermsa) Deleted() bool {
	return wflt._deleted
}

// Insert inserts the WsubFinderLinksTermsa to the database.
func (wflt *WsubFinderLinksTermsa) Insert(db XODB) error {
	var err error

	// if already exist, bail
	if wflt._exists {
		return errors.New("insert failed: already exists")
	}

	// sql insert query, primary key must be provided
	const sqlstr = `INSERT INTO jlabo.wsub_finder_links_termsa (` +
		`link_id, term_id, weight` +
		`) VALUES (` +
		`?, ?, ?` +
		`)`

	// run query
	XOLog(sqlstr, wflt.LinkID, wflt.TermID, wflt.Weight)
	_, err = db.Exec(sqlstr, wflt.LinkID, wflt.TermID, wflt.Weight)
	if err != nil {
		return err
	}

	// set existence
	wflt._exists = true

	return nil
}

// Update updates the WsubFinderLinksTermsa in the database.
func (wflt *WsubFinderLinksTermsa) Update(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !wflt._exists {
		return errors.New("update failed: does not exist")
	}

	// if deleted, bail
	if wflt._deleted {
		return errors.New("update failed: marked for deletion")
	}

	// sql query with composite primary key
	const sqlstr = `UPDATE jlabo.wsub_finder_links_termsa SET ` +
		`weight = ?` +
		` WHERE link_id = ? AND term_id = ?`

	// run query
	XOLog(sqlstr, wflt.Weight, wflt.LinkID, wflt.TermID)
	_, err = db.Exec(sqlstr, wflt.Weight, wflt.LinkID, wflt.TermID)
	return err
}

// Save saves the WsubFinderLinksTermsa to the database.
func (wflt *WsubFinderLinksTermsa) Save(db XODB) error {
	if wflt.Exists() {
		return wflt.Update(db)
	}

	return wflt.Insert(db)
}

// Delete deletes the WsubFinderLinksTermsa from the database.
func (wflt *WsubFinderLinksTermsa) Delete(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !wflt._exists {
		return nil
	}

	// if deleted, bail
	if wflt._deleted {
		return nil
	}

	// sql query with composite primary key
	const sqlstr = `DELETE FROM jlabo.wsub_finder_links_termsa WHERE link_id = ? AND term_id = ?`

	// run query
	XOLog(sqlstr, wflt.LinkID, wflt.TermID)
	_, err = db.Exec(sqlstr, wflt.LinkID, wflt.TermID)
	if err != nil {
		return err
	}

	// set deleted
	wflt._deleted = true

	return nil
}

// WsubFinderLinksTermsasByLinkIDTermIDWeight retrieves a row from 'jlabo.wsub_finder_links_termsa' as a WsubFinderLinksTermsa.
//
// Generated from index 'idx_link_term_weight'.
func WsubFinderLinksTermsasByLinkIDTermIDWeight(db XODB, linkID uint, termID uint, weight float32) ([]*WsubFinderLinksTermsa, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`link_id, term_id, weight ` +
		`FROM jlabo.wsub_finder_links_termsa ` +
		`WHERE link_id = ? AND term_id = ? AND weight = ?`

	// run query
	XOLog(sqlstr, linkID, termID, weight)
	q, err := db.Query(sqlstr, linkID, termID, weight)
	if err != nil {
		return nil, err
	}
	defer q.Close()

	// load results
	res := []*WsubFinderLinksTermsa{}
	for q.Next() {
		wflt := WsubFinderLinksTermsa{
			_exists: true,
		}

		// scan
		err = q.Scan(&wflt.LinkID, &wflt.TermID, &wflt.Weight)
		if err != nil {
			return nil, err
		}

		res = append(res, &wflt)
	}

	return res, nil
}

// WsubFinderLinksTermsasByTermIDWeight retrieves a row from 'jlabo.wsub_finder_links_termsa' as a WsubFinderLinksTermsa.
//
// Generated from index 'idx_term_weight'.
func WsubFinderLinksTermsasByTermIDWeight(db XODB, termID uint, weight float32) ([]*WsubFinderLinksTermsa, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`link_id, term_id, weight ` +
		`FROM jlabo.wsub_finder_links_termsa ` +
		`WHERE term_id = ? AND weight = ?`

	// run query
	XOLog(sqlstr, termID, weight)
	q, err := db.Query(sqlstr, termID, weight)
	if err != nil {
		return nil, err
	}
	defer q.Close()

	// load results
	res := []*WsubFinderLinksTermsa{}
	for q.Next() {
		wflt := WsubFinderLinksTermsa{
			_exists: true,
		}

		// scan
		err = q.Scan(&wflt.LinkID, &wflt.TermID, &wflt.Weight)
		if err != nil {
			return nil, err
		}

		res = append(res, &wflt)
	}

	return res, nil
}

// WsubFinderLinksTermsaByTermID retrieves a row from 'jlabo.wsub_finder_links_termsa' as a WsubFinderLinksTermsa.
//
// Generated from index 'wsub_finder_links_termsa_term_id_pkey'.
func WsubFinderLinksTermsaByTermID(db XODB, termID uint) (*WsubFinderLinksTermsa, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`link_id, term_id, weight ` +
		`FROM jlabo.wsub_finder_links_termsa ` +
		`WHERE term_id = ?`

	// run query
	XOLog(sqlstr, termID)
	wflt := WsubFinderLinksTermsa{
		_exists: true,
	}

	err = db.QueryRow(sqlstr, termID).Scan(&wflt.LinkID, &wflt.TermID, &wflt.Weight)
	if err != nil {
		return nil, err
	}

	return &wflt, nil
}
