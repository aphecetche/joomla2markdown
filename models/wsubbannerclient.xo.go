// Package toto contains the types for schema 'jlabo'.
package toto

// Code generated by xo. DO NOT EDIT.

import (
	"errors"
	"time"
)

// WsubBannerClient represents a row from 'jlabo.wsub_banner_clients'.
type WsubBannerClient struct {
	ID               int       `json:"id"`                // id
	Name             string    `json:"name"`              // name
	Contact          string    `json:"contact"`           // contact
	Email            string    `json:"email"`             // email
	Extrainfo        string    `json:"extrainfo"`         // extrainfo
	State            int8      `json:"state"`             // state
	CheckedOut       uint      `json:"checked_out"`       // checked_out
	CheckedOutTime   time.Time `json:"checked_out_time"`  // checked_out_time
	Metakey          string    `json:"metakey"`           // metakey
	OwnPrefix        int8      `json:"own_prefix"`        // own_prefix
	MetakeyPrefix    string    `json:"metakey_prefix"`    // metakey_prefix
	PurchaseType     int8      `json:"purchase_type"`     // purchase_type
	TrackClicks      int8      `json:"track_clicks"`      // track_clicks
	TrackImpressions int8      `json:"track_impressions"` // track_impressions

	// xo fields
	_exists, _deleted bool
}

// Exists determines if the WsubBannerClient exists in the database.
func (wbc *WsubBannerClient) Exists() bool {
	return wbc._exists
}

// Deleted provides information if the WsubBannerClient has been deleted from the database.
func (wbc *WsubBannerClient) Deleted() bool {
	return wbc._deleted
}

// Insert inserts the WsubBannerClient to the database.
func (wbc *WsubBannerClient) Insert(db XODB) error {
	var err error

	// if already exist, bail
	if wbc._exists {
		return errors.New("insert failed: already exists")
	}

	// sql insert query, primary key provided by autoincrement
	const sqlstr = `INSERT INTO jlabo.wsub_banner_clients (` +
		`name, contact, email, extrainfo, state, checked_out, checked_out_time, metakey, own_prefix, metakey_prefix, purchase_type, track_clicks, track_impressions` +
		`) VALUES (` +
		`?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?` +
		`)`

	// run query
	XOLog(sqlstr, wbc.Name, wbc.Contact, wbc.Email, wbc.Extrainfo, wbc.State, wbc.CheckedOut, wbc.CheckedOutTime, wbc.Metakey, wbc.OwnPrefix, wbc.MetakeyPrefix, wbc.PurchaseType, wbc.TrackClicks, wbc.TrackImpressions)
	res, err := db.Exec(sqlstr, wbc.Name, wbc.Contact, wbc.Email, wbc.Extrainfo, wbc.State, wbc.CheckedOut, wbc.CheckedOutTime, wbc.Metakey, wbc.OwnPrefix, wbc.MetakeyPrefix, wbc.PurchaseType, wbc.TrackClicks, wbc.TrackImpressions)
	if err != nil {
		return err
	}

	// retrieve id
	id, err := res.LastInsertId()
	if err != nil {
		return err
	}

	// set primary key and existence
	wbc.ID = int(id)
	wbc._exists = true

	return nil
}

// Update updates the WsubBannerClient in the database.
func (wbc *WsubBannerClient) Update(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !wbc._exists {
		return errors.New("update failed: does not exist")
	}

	// if deleted, bail
	if wbc._deleted {
		return errors.New("update failed: marked for deletion")
	}

	// sql query
	const sqlstr = `UPDATE jlabo.wsub_banner_clients SET ` +
		`name = ?, contact = ?, email = ?, extrainfo = ?, state = ?, checked_out = ?, checked_out_time = ?, metakey = ?, own_prefix = ?, metakey_prefix = ?, purchase_type = ?, track_clicks = ?, track_impressions = ?` +
		` WHERE id = ?`

	// run query
	XOLog(sqlstr, wbc.Name, wbc.Contact, wbc.Email, wbc.Extrainfo, wbc.State, wbc.CheckedOut, wbc.CheckedOutTime, wbc.Metakey, wbc.OwnPrefix, wbc.MetakeyPrefix, wbc.PurchaseType, wbc.TrackClicks, wbc.TrackImpressions, wbc.ID)
	_, err = db.Exec(sqlstr, wbc.Name, wbc.Contact, wbc.Email, wbc.Extrainfo, wbc.State, wbc.CheckedOut, wbc.CheckedOutTime, wbc.Metakey, wbc.OwnPrefix, wbc.MetakeyPrefix, wbc.PurchaseType, wbc.TrackClicks, wbc.TrackImpressions, wbc.ID)
	return err
}

// Save saves the WsubBannerClient to the database.
func (wbc *WsubBannerClient) Save(db XODB) error {
	if wbc.Exists() {
		return wbc.Update(db)
	}

	return wbc.Insert(db)
}

// Delete deletes the WsubBannerClient from the database.
func (wbc *WsubBannerClient) Delete(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !wbc._exists {
		return nil
	}

	// if deleted, bail
	if wbc._deleted {
		return nil
	}

	// sql query
	const sqlstr = `DELETE FROM jlabo.wsub_banner_clients WHERE id = ?`

	// run query
	XOLog(sqlstr, wbc.ID)
	_, err = db.Exec(sqlstr, wbc.ID)
	if err != nil {
		return err
	}

	// set deleted
	wbc._deleted = true

	return nil
}

// WsubBannerClientsByMetakeyPrefix retrieves a row from 'jlabo.wsub_banner_clients' as a WsubBannerClient.
//
// Generated from index 'idx_metakey_prefix'.
func WsubBannerClientsByMetakeyPrefix(db XODB, metakeyPrefix string) ([]*WsubBannerClient, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`id, name, contact, email, extrainfo, state, checked_out, checked_out_time, metakey, own_prefix, metakey_prefix, purchase_type, track_clicks, track_impressions ` +
		`FROM jlabo.wsub_banner_clients ` +
		`WHERE metakey_prefix = ?`

	// run query
	XOLog(sqlstr, metakeyPrefix)
	q, err := db.Query(sqlstr, metakeyPrefix)
	if err != nil {
		return nil, err
	}
	defer q.Close()

	// load results
	res := []*WsubBannerClient{}
	for q.Next() {
		wbc := WsubBannerClient{
			_exists: true,
		}

		// scan
		err = q.Scan(&wbc.ID, &wbc.Name, &wbc.Contact, &wbc.Email, &wbc.Extrainfo, &wbc.State, &wbc.CheckedOut, &wbc.CheckedOutTime, &wbc.Metakey, &wbc.OwnPrefix, &wbc.MetakeyPrefix, &wbc.PurchaseType, &wbc.TrackClicks, &wbc.TrackImpressions)
		if err != nil {
			return nil, err
		}

		res = append(res, &wbc)
	}

	return res, nil
}

// WsubBannerClientsByOwnPrefix retrieves a row from 'jlabo.wsub_banner_clients' as a WsubBannerClient.
//
// Generated from index 'idx_own_prefix'.
func WsubBannerClientsByOwnPrefix(db XODB, ownPrefix int8) ([]*WsubBannerClient, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`id, name, contact, email, extrainfo, state, checked_out, checked_out_time, metakey, own_prefix, metakey_prefix, purchase_type, track_clicks, track_impressions ` +
		`FROM jlabo.wsub_banner_clients ` +
		`WHERE own_prefix = ?`

	// run query
	XOLog(sqlstr, ownPrefix)
	q, err := db.Query(sqlstr, ownPrefix)
	if err != nil {
		return nil, err
	}
	defer q.Close()

	// load results
	res := []*WsubBannerClient{}
	for q.Next() {
		wbc := WsubBannerClient{
			_exists: true,
		}

		// scan
		err = q.Scan(&wbc.ID, &wbc.Name, &wbc.Contact, &wbc.Email, &wbc.Extrainfo, &wbc.State, &wbc.CheckedOut, &wbc.CheckedOutTime, &wbc.Metakey, &wbc.OwnPrefix, &wbc.MetakeyPrefix, &wbc.PurchaseType, &wbc.TrackClicks, &wbc.TrackImpressions)
		if err != nil {
			return nil, err
		}

		res = append(res, &wbc)
	}

	return res, nil
}

// WsubBannerClientByID retrieves a row from 'jlabo.wsub_banner_clients' as a WsubBannerClient.
//
// Generated from index 'wsub_banner_clients_id_pkey'.
func WsubBannerClientByID(db XODB, id int) (*WsubBannerClient, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`id, name, contact, email, extrainfo, state, checked_out, checked_out_time, metakey, own_prefix, metakey_prefix, purchase_type, track_clicks, track_impressions ` +
		`FROM jlabo.wsub_banner_clients ` +
		`WHERE id = ?`

	// run query
	XOLog(sqlstr, id)
	wbc := WsubBannerClient{
		_exists: true,
	}

	err = db.QueryRow(sqlstr, id).Scan(&wbc.ID, &wbc.Name, &wbc.Contact, &wbc.Email, &wbc.Extrainfo, &wbc.State, &wbc.CheckedOut, &wbc.CheckedOutTime, &wbc.Metakey, &wbc.OwnPrefix, &wbc.MetakeyPrefix, &wbc.PurchaseType, &wbc.TrackClicks, &wbc.TrackImpressions)
	if err != nil {
		return nil, err
	}

	return &wbc, nil
}
