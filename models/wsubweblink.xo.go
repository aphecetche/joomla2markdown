// Package toto contains the types for schema 'jlabo'.
package toto

// Code generated by xo. DO NOT EDIT.

import (
	"errors"
	"time"
)

// WsubWeblink represents a row from 'jlabo.wsub_weblinks'.
type WsubWeblink struct {
	ID             uint      `json:"id"`               // id
	Catid          int       `json:"catid"`            // catid
	Title          string    `json:"title"`            // title
	Alias          string    `json:"alias"`            // alias
	URL            string    `json:"url"`              // url
	Description    string    `json:"description"`      // description
	Hits           int       `json:"hits"`             // hits
	State          bool      `json:"state"`            // state
	CheckedOut     int       `json:"checked_out"`      // checked_out
	CheckedOutTime time.Time `json:"checked_out_time"` // checked_out_time
	Ordering       int       `json:"ordering"`         // ordering
	Access         int       `json:"access"`           // access
	Params         string    `json:"params"`           // params
	Language       string    `json:"language"`         // language
	Created        time.Time `json:"created"`          // created
	CreatedBy      uint      `json:"created_by"`       // created_by
	CreatedByAlias string    `json:"created_by_alias"` // created_by_alias
	Modified       time.Time `json:"modified"`         // modified
	ModifiedBy     uint      `json:"modified_by"`      // modified_by
	Metakey        string    `json:"metakey"`          // metakey
	Metadesc       string    `json:"metadesc"`         // metadesc
	Metadata       string    `json:"metadata"`         // metadata
	Featured       int8      `json:"featured"`         // featured
	Xreference     string    `json:"xreference"`       // xreference
	PublishUp      time.Time `json:"publish_up"`       // publish_up
	PublishDown    time.Time `json:"publish_down"`     // publish_down
	Version        uint      `json:"version"`          // version
	Images         string    `json:"images"`           // images

	// xo fields
	_exists, _deleted bool
}

// Exists determines if the WsubWeblink exists in the database.
func (ww *WsubWeblink) Exists() bool {
	return ww._exists
}

// Deleted provides information if the WsubWeblink has been deleted from the database.
func (ww *WsubWeblink) Deleted() bool {
	return ww._deleted
}

// Insert inserts the WsubWeblink to the database.
func (ww *WsubWeblink) Insert(db XODB) error {
	var err error

	// if already exist, bail
	if ww._exists {
		return errors.New("insert failed: already exists")
	}

	// sql insert query, primary key provided by autoincrement
	const sqlstr = `INSERT INTO jlabo.wsub_weblinks (` +
		`catid, title, alias, url, description, hits, state, checked_out, checked_out_time, ordering, access, params, language, created, created_by, created_by_alias, modified, modified_by, metakey, metadesc, metadata, featured, xreference, publish_up, publish_down, version, images` +
		`) VALUES (` +
		`?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?` +
		`)`

	// run query
	XOLog(sqlstr, ww.Catid, ww.Title, ww.Alias, ww.URL, ww.Description, ww.Hits, ww.State, ww.CheckedOut, ww.CheckedOutTime, ww.Ordering, ww.Access, ww.Params, ww.Language, ww.Created, ww.CreatedBy, ww.CreatedByAlias, ww.Modified, ww.ModifiedBy, ww.Metakey, ww.Metadesc, ww.Metadata, ww.Featured, ww.Xreference, ww.PublishUp, ww.PublishDown, ww.Version, ww.Images)
	res, err := db.Exec(sqlstr, ww.Catid, ww.Title, ww.Alias, ww.URL, ww.Description, ww.Hits, ww.State, ww.CheckedOut, ww.CheckedOutTime, ww.Ordering, ww.Access, ww.Params, ww.Language, ww.Created, ww.CreatedBy, ww.CreatedByAlias, ww.Modified, ww.ModifiedBy, ww.Metakey, ww.Metadesc, ww.Metadata, ww.Featured, ww.Xreference, ww.PublishUp, ww.PublishDown, ww.Version, ww.Images)
	if err != nil {
		return err
	}

	// retrieve id
	id, err := res.LastInsertId()
	if err != nil {
		return err
	}

	// set primary key and existence
	ww.ID = uint(id)
	ww._exists = true

	return nil
}

// Update updates the WsubWeblink in the database.
func (ww *WsubWeblink) Update(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !ww._exists {
		return errors.New("update failed: does not exist")
	}

	// if deleted, bail
	if ww._deleted {
		return errors.New("update failed: marked for deletion")
	}

	// sql query
	const sqlstr = `UPDATE jlabo.wsub_weblinks SET ` +
		`catid = ?, title = ?, alias = ?, url = ?, description = ?, hits = ?, state = ?, checked_out = ?, checked_out_time = ?, ordering = ?, access = ?, params = ?, language = ?, created = ?, created_by = ?, created_by_alias = ?, modified = ?, modified_by = ?, metakey = ?, metadesc = ?, metadata = ?, featured = ?, xreference = ?, publish_up = ?, publish_down = ?, version = ?, images = ?` +
		` WHERE id = ?`

	// run query
	XOLog(sqlstr, ww.Catid, ww.Title, ww.Alias, ww.URL, ww.Description, ww.Hits, ww.State, ww.CheckedOut, ww.CheckedOutTime, ww.Ordering, ww.Access, ww.Params, ww.Language, ww.Created, ww.CreatedBy, ww.CreatedByAlias, ww.Modified, ww.ModifiedBy, ww.Metakey, ww.Metadesc, ww.Metadata, ww.Featured, ww.Xreference, ww.PublishUp, ww.PublishDown, ww.Version, ww.Images, ww.ID)
	_, err = db.Exec(sqlstr, ww.Catid, ww.Title, ww.Alias, ww.URL, ww.Description, ww.Hits, ww.State, ww.CheckedOut, ww.CheckedOutTime, ww.Ordering, ww.Access, ww.Params, ww.Language, ww.Created, ww.CreatedBy, ww.CreatedByAlias, ww.Modified, ww.ModifiedBy, ww.Metakey, ww.Metadesc, ww.Metadata, ww.Featured, ww.Xreference, ww.PublishUp, ww.PublishDown, ww.Version, ww.Images, ww.ID)
	return err
}

// Save saves the WsubWeblink to the database.
func (ww *WsubWeblink) Save(db XODB) error {
	if ww.Exists() {
		return ww.Update(db)
	}

	return ww.Insert(db)
}

// Delete deletes the WsubWeblink from the database.
func (ww *WsubWeblink) Delete(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !ww._exists {
		return nil
	}

	// if deleted, bail
	if ww._deleted {
		return nil
	}

	// sql query
	const sqlstr = `DELETE FROM jlabo.wsub_weblinks WHERE id = ?`

	// run query
	XOLog(sqlstr, ww.ID)
	_, err = db.Exec(sqlstr, ww.ID)
	if err != nil {
		return err
	}

	// set deleted
	ww._deleted = true

	return nil
}

// WsubWeblinksByAccess retrieves a row from 'jlabo.wsub_weblinks' as a WsubWeblink.
//
// Generated from index 'idx_access'.
func WsubWeblinksByAccess(db XODB, access int) ([]*WsubWeblink, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`id, catid, title, alias, url, description, hits, state, checked_out, checked_out_time, ordering, access, params, language, created, created_by, created_by_alias, modified, modified_by, metakey, metadesc, metadata, featured, xreference, publish_up, publish_down, version, images ` +
		`FROM jlabo.wsub_weblinks ` +
		`WHERE access = ?`

	// run query
	XOLog(sqlstr, access)
	q, err := db.Query(sqlstr, access)
	if err != nil {
		return nil, err
	}
	defer q.Close()

	// load results
	res := []*WsubWeblink{}
	for q.Next() {
		ww := WsubWeblink{
			_exists: true,
		}

		// scan
		err = q.Scan(&ww.ID, &ww.Catid, &ww.Title, &ww.Alias, &ww.URL, &ww.Description, &ww.Hits, &ww.State, &ww.CheckedOut, &ww.CheckedOutTime, &ww.Ordering, &ww.Access, &ww.Params, &ww.Language, &ww.Created, &ww.CreatedBy, &ww.CreatedByAlias, &ww.Modified, &ww.ModifiedBy, &ww.Metakey, &ww.Metadesc, &ww.Metadata, &ww.Featured, &ww.Xreference, &ww.PublishUp, &ww.PublishDown, &ww.Version, &ww.Images)
		if err != nil {
			return nil, err
		}

		res = append(res, &ww)
	}

	return res, nil
}

// WsubWeblinksByCatid retrieves a row from 'jlabo.wsub_weblinks' as a WsubWeblink.
//
// Generated from index 'idx_catid'.
func WsubWeblinksByCatid(db XODB, catid int) ([]*WsubWeblink, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`id, catid, title, alias, url, description, hits, state, checked_out, checked_out_time, ordering, access, params, language, created, created_by, created_by_alias, modified, modified_by, metakey, metadesc, metadata, featured, xreference, publish_up, publish_down, version, images ` +
		`FROM jlabo.wsub_weblinks ` +
		`WHERE catid = ?`

	// run query
	XOLog(sqlstr, catid)
	q, err := db.Query(sqlstr, catid)
	if err != nil {
		return nil, err
	}
	defer q.Close()

	// load results
	res := []*WsubWeblink{}
	for q.Next() {
		ww := WsubWeblink{
			_exists: true,
		}

		// scan
		err = q.Scan(&ww.ID, &ww.Catid, &ww.Title, &ww.Alias, &ww.URL, &ww.Description, &ww.Hits, &ww.State, &ww.CheckedOut, &ww.CheckedOutTime, &ww.Ordering, &ww.Access, &ww.Params, &ww.Language, &ww.Created, &ww.CreatedBy, &ww.CreatedByAlias, &ww.Modified, &ww.ModifiedBy, &ww.Metakey, &ww.Metadesc, &ww.Metadata, &ww.Featured, &ww.Xreference, &ww.PublishUp, &ww.PublishDown, &ww.Version, &ww.Images)
		if err != nil {
			return nil, err
		}

		res = append(res, &ww)
	}

	return res, nil
}

// WsubWeblinksByCheckedOut retrieves a row from 'jlabo.wsub_weblinks' as a WsubWeblink.
//
// Generated from index 'idx_checkout'.
func WsubWeblinksByCheckedOut(db XODB, checkedOut int) ([]*WsubWeblink, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`id, catid, title, alias, url, description, hits, state, checked_out, checked_out_time, ordering, access, params, language, created, created_by, created_by_alias, modified, modified_by, metakey, metadesc, metadata, featured, xreference, publish_up, publish_down, version, images ` +
		`FROM jlabo.wsub_weblinks ` +
		`WHERE checked_out = ?`

	// run query
	XOLog(sqlstr, checkedOut)
	q, err := db.Query(sqlstr, checkedOut)
	if err != nil {
		return nil, err
	}
	defer q.Close()

	// load results
	res := []*WsubWeblink{}
	for q.Next() {
		ww := WsubWeblink{
			_exists: true,
		}

		// scan
		err = q.Scan(&ww.ID, &ww.Catid, &ww.Title, &ww.Alias, &ww.URL, &ww.Description, &ww.Hits, &ww.State, &ww.CheckedOut, &ww.CheckedOutTime, &ww.Ordering, &ww.Access, &ww.Params, &ww.Language, &ww.Created, &ww.CreatedBy, &ww.CreatedByAlias, &ww.Modified, &ww.ModifiedBy, &ww.Metakey, &ww.Metadesc, &ww.Metadata, &ww.Featured, &ww.Xreference, &ww.PublishUp, &ww.PublishDown, &ww.Version, &ww.Images)
		if err != nil {
			return nil, err
		}

		res = append(res, &ww)
	}

	return res, nil
}

// WsubWeblinksByCreatedBy retrieves a row from 'jlabo.wsub_weblinks' as a WsubWeblink.
//
// Generated from index 'idx_createdby'.
func WsubWeblinksByCreatedBy(db XODB, createdBy uint) ([]*WsubWeblink, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`id, catid, title, alias, url, description, hits, state, checked_out, checked_out_time, ordering, access, params, language, created, created_by, created_by_alias, modified, modified_by, metakey, metadesc, metadata, featured, xreference, publish_up, publish_down, version, images ` +
		`FROM jlabo.wsub_weblinks ` +
		`WHERE created_by = ?`

	// run query
	XOLog(sqlstr, createdBy)
	q, err := db.Query(sqlstr, createdBy)
	if err != nil {
		return nil, err
	}
	defer q.Close()

	// load results
	res := []*WsubWeblink{}
	for q.Next() {
		ww := WsubWeblink{
			_exists: true,
		}

		// scan
		err = q.Scan(&ww.ID, &ww.Catid, &ww.Title, &ww.Alias, &ww.URL, &ww.Description, &ww.Hits, &ww.State, &ww.CheckedOut, &ww.CheckedOutTime, &ww.Ordering, &ww.Access, &ww.Params, &ww.Language, &ww.Created, &ww.CreatedBy, &ww.CreatedByAlias, &ww.Modified, &ww.ModifiedBy, &ww.Metakey, &ww.Metadesc, &ww.Metadata, &ww.Featured, &ww.Xreference, &ww.PublishUp, &ww.PublishDown, &ww.Version, &ww.Images)
		if err != nil {
			return nil, err
		}

		res = append(res, &ww)
	}

	return res, nil
}

// WsubWeblinksByFeaturedCatid retrieves a row from 'jlabo.wsub_weblinks' as a WsubWeblink.
//
// Generated from index 'idx_featured_catid'.
func WsubWeblinksByFeaturedCatid(db XODB, featured int8, catid int) ([]*WsubWeblink, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`id, catid, title, alias, url, description, hits, state, checked_out, checked_out_time, ordering, access, params, language, created, created_by, created_by_alias, modified, modified_by, metakey, metadesc, metadata, featured, xreference, publish_up, publish_down, version, images ` +
		`FROM jlabo.wsub_weblinks ` +
		`WHERE featured = ? AND catid = ?`

	// run query
	XOLog(sqlstr, featured, catid)
	q, err := db.Query(sqlstr, featured, catid)
	if err != nil {
		return nil, err
	}
	defer q.Close()

	// load results
	res := []*WsubWeblink{}
	for q.Next() {
		ww := WsubWeblink{
			_exists: true,
		}

		// scan
		err = q.Scan(&ww.ID, &ww.Catid, &ww.Title, &ww.Alias, &ww.URL, &ww.Description, &ww.Hits, &ww.State, &ww.CheckedOut, &ww.CheckedOutTime, &ww.Ordering, &ww.Access, &ww.Params, &ww.Language, &ww.Created, &ww.CreatedBy, &ww.CreatedByAlias, &ww.Modified, &ww.ModifiedBy, &ww.Metakey, &ww.Metadesc, &ww.Metadata, &ww.Featured, &ww.Xreference, &ww.PublishUp, &ww.PublishDown, &ww.Version, &ww.Images)
		if err != nil {
			return nil, err
		}

		res = append(res, &ww)
	}

	return res, nil
}

// WsubWeblinksByLanguage retrieves a row from 'jlabo.wsub_weblinks' as a WsubWeblink.
//
// Generated from index 'idx_language'.
func WsubWeblinksByLanguage(db XODB, language string) ([]*WsubWeblink, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`id, catid, title, alias, url, description, hits, state, checked_out, checked_out_time, ordering, access, params, language, created, created_by, created_by_alias, modified, modified_by, metakey, metadesc, metadata, featured, xreference, publish_up, publish_down, version, images ` +
		`FROM jlabo.wsub_weblinks ` +
		`WHERE language = ?`

	// run query
	XOLog(sqlstr, language)
	q, err := db.Query(sqlstr, language)
	if err != nil {
		return nil, err
	}
	defer q.Close()

	// load results
	res := []*WsubWeblink{}
	for q.Next() {
		ww := WsubWeblink{
			_exists: true,
		}

		// scan
		err = q.Scan(&ww.ID, &ww.Catid, &ww.Title, &ww.Alias, &ww.URL, &ww.Description, &ww.Hits, &ww.State, &ww.CheckedOut, &ww.CheckedOutTime, &ww.Ordering, &ww.Access, &ww.Params, &ww.Language, &ww.Created, &ww.CreatedBy, &ww.CreatedByAlias, &ww.Modified, &ww.ModifiedBy, &ww.Metakey, &ww.Metadesc, &ww.Metadata, &ww.Featured, &ww.Xreference, &ww.PublishUp, &ww.PublishDown, &ww.Version, &ww.Images)
		if err != nil {
			return nil, err
		}

		res = append(res, &ww)
	}

	return res, nil
}

// WsubWeblinksByState retrieves a row from 'jlabo.wsub_weblinks' as a WsubWeblink.
//
// Generated from index 'idx_state'.
func WsubWeblinksByState(db XODB, state bool) ([]*WsubWeblink, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`id, catid, title, alias, url, description, hits, state, checked_out, checked_out_time, ordering, access, params, language, created, created_by, created_by_alias, modified, modified_by, metakey, metadesc, metadata, featured, xreference, publish_up, publish_down, version, images ` +
		`FROM jlabo.wsub_weblinks ` +
		`WHERE state = ?`

	// run query
	XOLog(sqlstr, state)
	q, err := db.Query(sqlstr, state)
	if err != nil {
		return nil, err
	}
	defer q.Close()

	// load results
	res := []*WsubWeblink{}
	for q.Next() {
		ww := WsubWeblink{
			_exists: true,
		}

		// scan
		err = q.Scan(&ww.ID, &ww.Catid, &ww.Title, &ww.Alias, &ww.URL, &ww.Description, &ww.Hits, &ww.State, &ww.CheckedOut, &ww.CheckedOutTime, &ww.Ordering, &ww.Access, &ww.Params, &ww.Language, &ww.Created, &ww.CreatedBy, &ww.CreatedByAlias, &ww.Modified, &ww.ModifiedBy, &ww.Metakey, &ww.Metadesc, &ww.Metadata, &ww.Featured, &ww.Xreference, &ww.PublishUp, &ww.PublishDown, &ww.Version, &ww.Images)
		if err != nil {
			return nil, err
		}

		res = append(res, &ww)
	}

	return res, nil
}

// WsubWeblinksByXreference retrieves a row from 'jlabo.wsub_weblinks' as a WsubWeblink.
//
// Generated from index 'idx_xreference'.
func WsubWeblinksByXreference(db XODB, xreference string) ([]*WsubWeblink, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`id, catid, title, alias, url, description, hits, state, checked_out, checked_out_time, ordering, access, params, language, created, created_by, created_by_alias, modified, modified_by, metakey, metadesc, metadata, featured, xreference, publish_up, publish_down, version, images ` +
		`FROM jlabo.wsub_weblinks ` +
		`WHERE xreference = ?`

	// run query
	XOLog(sqlstr, xreference)
	q, err := db.Query(sqlstr, xreference)
	if err != nil {
		return nil, err
	}
	defer q.Close()

	// load results
	res := []*WsubWeblink{}
	for q.Next() {
		ww := WsubWeblink{
			_exists: true,
		}

		// scan
		err = q.Scan(&ww.ID, &ww.Catid, &ww.Title, &ww.Alias, &ww.URL, &ww.Description, &ww.Hits, &ww.State, &ww.CheckedOut, &ww.CheckedOutTime, &ww.Ordering, &ww.Access, &ww.Params, &ww.Language, &ww.Created, &ww.CreatedBy, &ww.CreatedByAlias, &ww.Modified, &ww.ModifiedBy, &ww.Metakey, &ww.Metadesc, &ww.Metadata, &ww.Featured, &ww.Xreference, &ww.PublishUp, &ww.PublishDown, &ww.Version, &ww.Images)
		if err != nil {
			return nil, err
		}

		res = append(res, &ww)
	}

	return res, nil
}

// WsubWeblinkByID retrieves a row from 'jlabo.wsub_weblinks' as a WsubWeblink.
//
// Generated from index 'wsub_weblinks_id_pkey'.
func WsubWeblinkByID(db XODB, id uint) (*WsubWeblink, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`id, catid, title, alias, url, description, hits, state, checked_out, checked_out_time, ordering, access, params, language, created, created_by, created_by_alias, modified, modified_by, metakey, metadesc, metadata, featured, xreference, publish_up, publish_down, version, images ` +
		`FROM jlabo.wsub_weblinks ` +
		`WHERE id = ?`

	// run query
	XOLog(sqlstr, id)
	ww := WsubWeblink{
		_exists: true,
	}

	err = db.QueryRow(sqlstr, id).Scan(&ww.ID, &ww.Catid, &ww.Title, &ww.Alias, &ww.URL, &ww.Description, &ww.Hits, &ww.State, &ww.CheckedOut, &ww.CheckedOutTime, &ww.Ordering, &ww.Access, &ww.Params, &ww.Language, &ww.Created, &ww.CreatedBy, &ww.CreatedByAlias, &ww.Modified, &ww.ModifiedBy, &ww.Metakey, &ww.Metadesc, &ww.Metadata, &ww.Featured, &ww.Xreference, &ww.PublishUp, &ww.PublishDown, &ww.Version, &ww.Images)
	if err != nil {
		return nil, err
	}

	return &ww, nil
}
