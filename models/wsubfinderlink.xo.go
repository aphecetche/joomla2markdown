// Package toto contains the types for schema 'jlabo'.
package toto

// Code generated by xo. DO NOT EDIT.

import (
	"database/sql"
	"errors"
	"time"
)

// WsubFinderLink represents a row from 'jlabo.wsub_finder_links'.
type WsubFinderLink struct {
	LinkID           uint           `json:"link_id"`            // link_id
	URL              string         `json:"url"`                // url
	Route            string         `json:"route"`              // route
	Title            sql.NullString `json:"title"`              // title
	Description      sql.NullString `json:"description"`        // description
	Indexdate        time.Time      `json:"indexdate"`          // indexdate
	Md5sum           sql.NullString `json:"md5sum"`             // md5sum
	Published        bool           `json:"published"`          // published
	State            sql.NullInt64  `json:"state"`              // state
	Access           sql.NullInt64  `json:"access"`             // access
	Language         string         `json:"language"`           // language
	PublishStartDate time.Time      `json:"publish_start_date"` // publish_start_date
	PublishEndDate   time.Time      `json:"publish_end_date"`   // publish_end_date
	StartDate        time.Time      `json:"start_date"`         // start_date
	EndDate          time.Time      `json:"end_date"`           // end_date
	ListPrice        float64        `json:"list_price"`         // list_price
	SalePrice        float64        `json:"sale_price"`         // sale_price
	TypeID           int            `json:"type_id"`            // type_id
	Object           []byte         `json:"object"`             // object

	// xo fields
	_exists, _deleted bool
}

// Exists determines if the WsubFinderLink exists in the database.
func (wfl *WsubFinderLink) Exists() bool {
	return wfl._exists
}

// Deleted provides information if the WsubFinderLink has been deleted from the database.
func (wfl *WsubFinderLink) Deleted() bool {
	return wfl._deleted
}

// Insert inserts the WsubFinderLink to the database.
func (wfl *WsubFinderLink) Insert(db XODB) error {
	var err error

	// if already exist, bail
	if wfl._exists {
		return errors.New("insert failed: already exists")
	}

	// sql insert query, primary key provided by autoincrement
	const sqlstr = `INSERT INTO jlabo.wsub_finder_links (` +
		`url, route, title, description, indexdate, md5sum, published, state, access, language, publish_start_date, publish_end_date, start_date, end_date, list_price, sale_price, type_id, object` +
		`) VALUES (` +
		`?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?` +
		`)`

	// run query
	XOLog(sqlstr, wfl.URL, wfl.Route, wfl.Title, wfl.Description, wfl.Indexdate, wfl.Md5sum, wfl.Published, wfl.State, wfl.Access, wfl.Language, wfl.PublishStartDate, wfl.PublishEndDate, wfl.StartDate, wfl.EndDate, wfl.ListPrice, wfl.SalePrice, wfl.TypeID, wfl.Object)
	res, err := db.Exec(sqlstr, wfl.URL, wfl.Route, wfl.Title, wfl.Description, wfl.Indexdate, wfl.Md5sum, wfl.Published, wfl.State, wfl.Access, wfl.Language, wfl.PublishStartDate, wfl.PublishEndDate, wfl.StartDate, wfl.EndDate, wfl.ListPrice, wfl.SalePrice, wfl.TypeID, wfl.Object)
	if err != nil {
		return err
	}

	// retrieve id
	id, err := res.LastInsertId()
	if err != nil {
		return err
	}

	// set primary key and existence
	wfl.LinkID = uint(id)
	wfl._exists = true

	return nil
}

// Update updates the WsubFinderLink in the database.
func (wfl *WsubFinderLink) Update(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !wfl._exists {
		return errors.New("update failed: does not exist")
	}

	// if deleted, bail
	if wfl._deleted {
		return errors.New("update failed: marked for deletion")
	}

	// sql query
	const sqlstr = `UPDATE jlabo.wsub_finder_links SET ` +
		`url = ?, route = ?, title = ?, description = ?, indexdate = ?, md5sum = ?, published = ?, state = ?, access = ?, language = ?, publish_start_date = ?, publish_end_date = ?, start_date = ?, end_date = ?, list_price = ?, sale_price = ?, type_id = ?, object = ?` +
		` WHERE link_id = ?`

	// run query
	XOLog(sqlstr, wfl.URL, wfl.Route, wfl.Title, wfl.Description, wfl.Indexdate, wfl.Md5sum, wfl.Published, wfl.State, wfl.Access, wfl.Language, wfl.PublishStartDate, wfl.PublishEndDate, wfl.StartDate, wfl.EndDate, wfl.ListPrice, wfl.SalePrice, wfl.TypeID, wfl.Object, wfl.LinkID)
	_, err = db.Exec(sqlstr, wfl.URL, wfl.Route, wfl.Title, wfl.Description, wfl.Indexdate, wfl.Md5sum, wfl.Published, wfl.State, wfl.Access, wfl.Language, wfl.PublishStartDate, wfl.PublishEndDate, wfl.StartDate, wfl.EndDate, wfl.ListPrice, wfl.SalePrice, wfl.TypeID, wfl.Object, wfl.LinkID)
	return err
}

// Save saves the WsubFinderLink to the database.
func (wfl *WsubFinderLink) Save(db XODB) error {
	if wfl.Exists() {
		return wfl.Update(db)
	}

	return wfl.Insert(db)
}

// Delete deletes the WsubFinderLink from the database.
func (wfl *WsubFinderLink) Delete(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !wfl._exists {
		return nil
	}

	// if deleted, bail
	if wfl._deleted {
		return nil
	}

	// sql query
	const sqlstr = `DELETE FROM jlabo.wsub_finder_links WHERE link_id = ?`

	// run query
	XOLog(sqlstr, wfl.LinkID)
	_, err = db.Exec(sqlstr, wfl.LinkID)
	if err != nil {
		return err
	}

	// set deleted
	wfl._deleted = true

	return nil
}

// WsubFinderLinksByMd5sum retrieves a row from 'jlabo.wsub_finder_links' as a WsubFinderLink.
//
// Generated from index 'idx_md5'.
func WsubFinderLinksByMd5sum(db XODB, md5sum sql.NullString) ([]*WsubFinderLink, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`link_id, url, route, title, description, indexdate, md5sum, published, state, access, language, publish_start_date, publish_end_date, start_date, end_date, list_price, sale_price, type_id, object ` +
		`FROM jlabo.wsub_finder_links ` +
		`WHERE md5sum = ?`

	// run query
	XOLog(sqlstr, md5sum)
	q, err := db.Query(sqlstr, md5sum)
	if err != nil {
		return nil, err
	}
	defer q.Close()

	// load results
	res := []*WsubFinderLink{}
	for q.Next() {
		wfl := WsubFinderLink{
			_exists: true,
		}

		// scan
		err = q.Scan(&wfl.LinkID, &wfl.URL, &wfl.Route, &wfl.Title, &wfl.Description, &wfl.Indexdate, &wfl.Md5sum, &wfl.Published, &wfl.State, &wfl.Access, &wfl.Language, &wfl.PublishStartDate, &wfl.PublishEndDate, &wfl.StartDate, &wfl.EndDate, &wfl.ListPrice, &wfl.SalePrice, &wfl.TypeID, &wfl.Object)
		if err != nil {
			return nil, err
		}

		res = append(res, &wfl)
	}

	return res, nil
}

// WsubFinderLinksByPublishedStateAccessPublishStartDatePublishEndDateListPrice retrieves a row from 'jlabo.wsub_finder_links' as a WsubFinderLink.
//
// Generated from index 'idx_published_list'.
func WsubFinderLinksByPublishedStateAccessPublishStartDatePublishEndDateListPrice(db XODB, published bool, state sql.NullInt64, access sql.NullInt64, publishStartDate time.Time, publishEndDate time.Time, listPrice float64) ([]*WsubFinderLink, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`link_id, url, route, title, description, indexdate, md5sum, published, state, access, language, publish_start_date, publish_end_date, start_date, end_date, list_price, sale_price, type_id, object ` +
		`FROM jlabo.wsub_finder_links ` +
		`WHERE published = ? AND state = ? AND access = ? AND publish_start_date = ? AND publish_end_date = ? AND list_price = ?`

	// run query
	XOLog(sqlstr, published, state, access, publishStartDate, publishEndDate, listPrice)
	q, err := db.Query(sqlstr, published, state, access, publishStartDate, publishEndDate, listPrice)
	if err != nil {
		return nil, err
	}
	defer q.Close()

	// load results
	res := []*WsubFinderLink{}
	for q.Next() {
		wfl := WsubFinderLink{
			_exists: true,
		}

		// scan
		err = q.Scan(&wfl.LinkID, &wfl.URL, &wfl.Route, &wfl.Title, &wfl.Description, &wfl.Indexdate, &wfl.Md5sum, &wfl.Published, &wfl.State, &wfl.Access, &wfl.Language, &wfl.PublishStartDate, &wfl.PublishEndDate, &wfl.StartDate, &wfl.EndDate, &wfl.ListPrice, &wfl.SalePrice, &wfl.TypeID, &wfl.Object)
		if err != nil {
			return nil, err
		}

		res = append(res, &wfl)
	}

	return res, nil
}

// WsubFinderLinksByPublishedStateAccessPublishStartDatePublishEndDateSalePrice retrieves a row from 'jlabo.wsub_finder_links' as a WsubFinderLink.
//
// Generated from index 'idx_published_sale'.
func WsubFinderLinksByPublishedStateAccessPublishStartDatePublishEndDateSalePrice(db XODB, published bool, state sql.NullInt64, access sql.NullInt64, publishStartDate time.Time, publishEndDate time.Time, salePrice float64) ([]*WsubFinderLink, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`link_id, url, route, title, description, indexdate, md5sum, published, state, access, language, publish_start_date, publish_end_date, start_date, end_date, list_price, sale_price, type_id, object ` +
		`FROM jlabo.wsub_finder_links ` +
		`WHERE published = ? AND state = ? AND access = ? AND publish_start_date = ? AND publish_end_date = ? AND sale_price = ?`

	// run query
	XOLog(sqlstr, published, state, access, publishStartDate, publishEndDate, salePrice)
	q, err := db.Query(sqlstr, published, state, access, publishStartDate, publishEndDate, salePrice)
	if err != nil {
		return nil, err
	}
	defer q.Close()

	// load results
	res := []*WsubFinderLink{}
	for q.Next() {
		wfl := WsubFinderLink{
			_exists: true,
		}

		// scan
		err = q.Scan(&wfl.LinkID, &wfl.URL, &wfl.Route, &wfl.Title, &wfl.Description, &wfl.Indexdate, &wfl.Md5sum, &wfl.Published, &wfl.State, &wfl.Access, &wfl.Language, &wfl.PublishStartDate, &wfl.PublishEndDate, &wfl.StartDate, &wfl.EndDate, &wfl.ListPrice, &wfl.SalePrice, &wfl.TypeID, &wfl.Object)
		if err != nil {
			return nil, err
		}

		res = append(res, &wfl)
	}

	return res, nil
}

// WsubFinderLinksByTitle retrieves a row from 'jlabo.wsub_finder_links' as a WsubFinderLink.
//
// Generated from index 'idx_title'.
func WsubFinderLinksByTitle(db XODB, title sql.NullString) ([]*WsubFinderLink, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`link_id, url, route, title, description, indexdate, md5sum, published, state, access, language, publish_start_date, publish_end_date, start_date, end_date, list_price, sale_price, type_id, object ` +
		`FROM jlabo.wsub_finder_links ` +
		`WHERE title = ?`

	// run query
	XOLog(sqlstr, title)
	q, err := db.Query(sqlstr, title)
	if err != nil {
		return nil, err
	}
	defer q.Close()

	// load results
	res := []*WsubFinderLink{}
	for q.Next() {
		wfl := WsubFinderLink{
			_exists: true,
		}

		// scan
		err = q.Scan(&wfl.LinkID, &wfl.URL, &wfl.Route, &wfl.Title, &wfl.Description, &wfl.Indexdate, &wfl.Md5sum, &wfl.Published, &wfl.State, &wfl.Access, &wfl.Language, &wfl.PublishStartDate, &wfl.PublishEndDate, &wfl.StartDate, &wfl.EndDate, &wfl.ListPrice, &wfl.SalePrice, &wfl.TypeID, &wfl.Object)
		if err != nil {
			return nil, err
		}

		res = append(res, &wfl)
	}

	return res, nil
}

// WsubFinderLinksByTypeID retrieves a row from 'jlabo.wsub_finder_links' as a WsubFinderLink.
//
// Generated from index 'idx_type'.
func WsubFinderLinksByTypeID(db XODB, typeID int) ([]*WsubFinderLink, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`link_id, url, route, title, description, indexdate, md5sum, published, state, access, language, publish_start_date, publish_end_date, start_date, end_date, list_price, sale_price, type_id, object ` +
		`FROM jlabo.wsub_finder_links ` +
		`WHERE type_id = ?`

	// run query
	XOLog(sqlstr, typeID)
	q, err := db.Query(sqlstr, typeID)
	if err != nil {
		return nil, err
	}
	defer q.Close()

	// load results
	res := []*WsubFinderLink{}
	for q.Next() {
		wfl := WsubFinderLink{
			_exists: true,
		}

		// scan
		err = q.Scan(&wfl.LinkID, &wfl.URL, &wfl.Route, &wfl.Title, &wfl.Description, &wfl.Indexdate, &wfl.Md5sum, &wfl.Published, &wfl.State, &wfl.Access, &wfl.Language, &wfl.PublishStartDate, &wfl.PublishEndDate, &wfl.StartDate, &wfl.EndDate, &wfl.ListPrice, &wfl.SalePrice, &wfl.TypeID, &wfl.Object)
		if err != nil {
			return nil, err
		}

		res = append(res, &wfl)
	}

	return res, nil
}

// WsubFinderLinksByURL retrieves a row from 'jlabo.wsub_finder_links' as a WsubFinderLink.
//
// Generated from index 'idx_url'.
func WsubFinderLinksByURL(db XODB, url string) ([]*WsubFinderLink, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`link_id, url, route, title, description, indexdate, md5sum, published, state, access, language, publish_start_date, publish_end_date, start_date, end_date, list_price, sale_price, type_id, object ` +
		`FROM jlabo.wsub_finder_links ` +
		`WHERE url = ?`

	// run query
	XOLog(sqlstr, url)
	q, err := db.Query(sqlstr, url)
	if err != nil {
		return nil, err
	}
	defer q.Close()

	// load results
	res := []*WsubFinderLink{}
	for q.Next() {
		wfl := WsubFinderLink{
			_exists: true,
		}

		// scan
		err = q.Scan(&wfl.LinkID, &wfl.URL, &wfl.Route, &wfl.Title, &wfl.Description, &wfl.Indexdate, &wfl.Md5sum, &wfl.Published, &wfl.State, &wfl.Access, &wfl.Language, &wfl.PublishStartDate, &wfl.PublishEndDate, &wfl.StartDate, &wfl.EndDate, &wfl.ListPrice, &wfl.SalePrice, &wfl.TypeID, &wfl.Object)
		if err != nil {
			return nil, err
		}

		res = append(res, &wfl)
	}

	return res, nil
}

// WsubFinderLinkByLinkID retrieves a row from 'jlabo.wsub_finder_links' as a WsubFinderLink.
//
// Generated from index 'wsub_finder_links_link_id_pkey'.
func WsubFinderLinkByLinkID(db XODB, linkID uint) (*WsubFinderLink, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`link_id, url, route, title, description, indexdate, md5sum, published, state, access, language, publish_start_date, publish_end_date, start_date, end_date, list_price, sale_price, type_id, object ` +
		`FROM jlabo.wsub_finder_links ` +
		`WHERE link_id = ?`

	// run query
	XOLog(sqlstr, linkID)
	wfl := WsubFinderLink{
		_exists: true,
	}

	err = db.QueryRow(sqlstr, linkID).Scan(&wfl.LinkID, &wfl.URL, &wfl.Route, &wfl.Title, &wfl.Description, &wfl.Indexdate, &wfl.Md5sum, &wfl.Published, &wfl.State, &wfl.Access, &wfl.Language, &wfl.PublishStartDate, &wfl.PublishEndDate, &wfl.StartDate, &wfl.EndDate, &wfl.ListPrice, &wfl.SalePrice, &wfl.TypeID, &wfl.Object)
	if err != nil {
		return nil, err
	}

	return &wfl, nil
}
