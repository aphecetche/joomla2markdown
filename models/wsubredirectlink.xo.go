// Package toto contains the types for schema 'jlabo'.
package toto

// Code generated by xo. DO NOT EDIT.

import (
	"database/sql"
	"errors"
	"time"
)

// WsubRedirectLink represents a row from 'jlabo.wsub_redirect_links'.
type WsubRedirectLink struct {
	ID           uint           `json:"id"`            // id
	OldURL       string         `json:"old_url"`       // old_url
	NewURL       sql.NullString `json:"new_url"`       // new_url
	Referer      string         `json:"referer"`       // referer
	Comment      string         `json:"comment"`       // comment
	Hits         uint           `json:"hits"`          // hits
	Published    int8           `json:"published"`     // published
	CreatedDate  time.Time      `json:"created_date"`  // created_date
	ModifiedDate time.Time      `json:"modified_date"` // modified_date
	Header       int16          `json:"header"`        // header

	// xo fields
	_exists, _deleted bool
}

// Exists determines if the WsubRedirectLink exists in the database.
func (wrl *WsubRedirectLink) Exists() bool {
	return wrl._exists
}

// Deleted provides information if the WsubRedirectLink has been deleted from the database.
func (wrl *WsubRedirectLink) Deleted() bool {
	return wrl._deleted
}

// Insert inserts the WsubRedirectLink to the database.
func (wrl *WsubRedirectLink) Insert(db XODB) error {
	var err error

	// if already exist, bail
	if wrl._exists {
		return errors.New("insert failed: already exists")
	}

	// sql insert query, primary key provided by autoincrement
	const sqlstr = `INSERT INTO jlabo.wsub_redirect_links (` +
		`old_url, new_url, referer, comment, hits, published, created_date, modified_date, header` +
		`) VALUES (` +
		`?, ?, ?, ?, ?, ?, ?, ?, ?` +
		`)`

	// run query
	XOLog(sqlstr, wrl.OldURL, wrl.NewURL, wrl.Referer, wrl.Comment, wrl.Hits, wrl.Published, wrl.CreatedDate, wrl.ModifiedDate, wrl.Header)
	res, err := db.Exec(sqlstr, wrl.OldURL, wrl.NewURL, wrl.Referer, wrl.Comment, wrl.Hits, wrl.Published, wrl.CreatedDate, wrl.ModifiedDate, wrl.Header)
	if err != nil {
		return err
	}

	// retrieve id
	id, err := res.LastInsertId()
	if err != nil {
		return err
	}

	// set primary key and existence
	wrl.ID = uint(id)
	wrl._exists = true

	return nil
}

// Update updates the WsubRedirectLink in the database.
func (wrl *WsubRedirectLink) Update(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !wrl._exists {
		return errors.New("update failed: does not exist")
	}

	// if deleted, bail
	if wrl._deleted {
		return errors.New("update failed: marked for deletion")
	}

	// sql query
	const sqlstr = `UPDATE jlabo.wsub_redirect_links SET ` +
		`old_url = ?, new_url = ?, referer = ?, comment = ?, hits = ?, published = ?, created_date = ?, modified_date = ?, header = ?` +
		` WHERE id = ?`

	// run query
	XOLog(sqlstr, wrl.OldURL, wrl.NewURL, wrl.Referer, wrl.Comment, wrl.Hits, wrl.Published, wrl.CreatedDate, wrl.ModifiedDate, wrl.Header, wrl.ID)
	_, err = db.Exec(sqlstr, wrl.OldURL, wrl.NewURL, wrl.Referer, wrl.Comment, wrl.Hits, wrl.Published, wrl.CreatedDate, wrl.ModifiedDate, wrl.Header, wrl.ID)
	return err
}

// Save saves the WsubRedirectLink to the database.
func (wrl *WsubRedirectLink) Save(db XODB) error {
	if wrl.Exists() {
		return wrl.Update(db)
	}

	return wrl.Insert(db)
}

// Delete deletes the WsubRedirectLink from the database.
func (wrl *WsubRedirectLink) Delete(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !wrl._exists {
		return nil
	}

	// if deleted, bail
	if wrl._deleted {
		return nil
	}

	// sql query
	const sqlstr = `DELETE FROM jlabo.wsub_redirect_links WHERE id = ?`

	// run query
	XOLog(sqlstr, wrl.ID)
	_, err = db.Exec(sqlstr, wrl.ID)
	if err != nil {
		return err
	}

	// set deleted
	wrl._deleted = true

	return nil
}

// WsubRedirectLinksByModifiedDate retrieves a row from 'jlabo.wsub_redirect_links' as a WsubRedirectLink.
//
// Generated from index 'idx_link_modifed'.
func WsubRedirectLinksByModifiedDate(db XODB, modifiedDate time.Time) ([]*WsubRedirectLink, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`id, old_url, new_url, referer, comment, hits, published, created_date, modified_date, header ` +
		`FROM jlabo.wsub_redirect_links ` +
		`WHERE modified_date = ?`

	// run query
	XOLog(sqlstr, modifiedDate)
	q, err := db.Query(sqlstr, modifiedDate)
	if err != nil {
		return nil, err
	}
	defer q.Close()

	// load results
	res := []*WsubRedirectLink{}
	for q.Next() {
		wrl := WsubRedirectLink{
			_exists: true,
		}

		// scan
		err = q.Scan(&wrl.ID, &wrl.OldURL, &wrl.NewURL, &wrl.Referer, &wrl.Comment, &wrl.Hits, &wrl.Published, &wrl.CreatedDate, &wrl.ModifiedDate, &wrl.Header)
		if err != nil {
			return nil, err
		}

		res = append(res, &wrl)
	}

	return res, nil
}

// WsubRedirectLinkByOldURL retrieves a row from 'jlabo.wsub_redirect_links' as a WsubRedirectLink.
//
// Generated from index 'idx_link_old'.
func WsubRedirectLinkByOldURL(db XODB, oldURL string) (*WsubRedirectLink, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`id, old_url, new_url, referer, comment, hits, published, created_date, modified_date, header ` +
		`FROM jlabo.wsub_redirect_links ` +
		`WHERE old_url = ?`

	// run query
	XOLog(sqlstr, oldURL)
	wrl := WsubRedirectLink{
		_exists: true,
	}

	err = db.QueryRow(sqlstr, oldURL).Scan(&wrl.ID, &wrl.OldURL, &wrl.NewURL, &wrl.Referer, &wrl.Comment, &wrl.Hits, &wrl.Published, &wrl.CreatedDate, &wrl.ModifiedDate, &wrl.Header)
	if err != nil {
		return nil, err
	}

	return &wrl, nil
}

// WsubRedirectLinkByID retrieves a row from 'jlabo.wsub_redirect_links' as a WsubRedirectLink.
//
// Generated from index 'wsub_redirect_links_id_pkey'.
func WsubRedirectLinkByID(db XODB, id uint) (*WsubRedirectLink, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`id, old_url, new_url, referer, comment, hits, published, created_date, modified_date, header ` +
		`FROM jlabo.wsub_redirect_links ` +
		`WHERE id = ?`

	// run query
	XOLog(sqlstr, id)
	wrl := WsubRedirectLink{
		_exists: true,
	}

	err = db.QueryRow(sqlstr, id).Scan(&wrl.ID, &wrl.OldURL, &wrl.NewURL, &wrl.Referer, &wrl.Comment, &wrl.Hits, &wrl.Published, &wrl.CreatedDate, &wrl.ModifiedDate, &wrl.Header)
	if err != nil {
		return nil, err
	}

	return &wrl, nil
}
