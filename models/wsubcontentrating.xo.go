// Package toto contains the types for schema 'jlabo'.
package toto

// Code generated by xo. DO NOT EDIT.

import (
	"errors"
)

// WsubContentRating represents a row from 'jlabo.wsub_content_rating'.
type WsubContentRating struct {
	ContentID   int    `json:"content_id"`   // content_id
	RatingSum   uint   `json:"rating_sum"`   // rating_sum
	RatingCount uint   `json:"rating_count"` // rating_count
	Lastip      string `json:"lastip"`       // lastip

	// xo fields
	_exists, _deleted bool
}

// Exists determines if the WsubContentRating exists in the database.
func (wcr *WsubContentRating) Exists() bool {
	return wcr._exists
}

// Deleted provides information if the WsubContentRating has been deleted from the database.
func (wcr *WsubContentRating) Deleted() bool {
	return wcr._deleted
}

// Insert inserts the WsubContentRating to the database.
func (wcr *WsubContentRating) Insert(db XODB) error {
	var err error

	// if already exist, bail
	if wcr._exists {
		return errors.New("insert failed: already exists")
	}

	// sql insert query, primary key must be provided
	const sqlstr = `INSERT INTO jlabo.wsub_content_rating (` +
		`content_id, rating_sum, rating_count, lastip` +
		`) VALUES (` +
		`?, ?, ?, ?` +
		`)`

	// run query
	XOLog(sqlstr, wcr.ContentID, wcr.RatingSum, wcr.RatingCount, wcr.Lastip)
	_, err = db.Exec(sqlstr, wcr.ContentID, wcr.RatingSum, wcr.RatingCount, wcr.Lastip)
	if err != nil {
		return err
	}

	// set existence
	wcr._exists = true

	return nil
}

// Update updates the WsubContentRating in the database.
func (wcr *WsubContentRating) Update(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !wcr._exists {
		return errors.New("update failed: does not exist")
	}

	// if deleted, bail
	if wcr._deleted {
		return errors.New("update failed: marked for deletion")
	}

	// sql query
	const sqlstr = `UPDATE jlabo.wsub_content_rating SET ` +
		`rating_sum = ?, rating_count = ?, lastip = ?` +
		` WHERE content_id = ?`

	// run query
	XOLog(sqlstr, wcr.RatingSum, wcr.RatingCount, wcr.Lastip, wcr.ContentID)
	_, err = db.Exec(sqlstr, wcr.RatingSum, wcr.RatingCount, wcr.Lastip, wcr.ContentID)
	return err
}

// Save saves the WsubContentRating to the database.
func (wcr *WsubContentRating) Save(db XODB) error {
	if wcr.Exists() {
		return wcr.Update(db)
	}

	return wcr.Insert(db)
}

// Delete deletes the WsubContentRating from the database.
func (wcr *WsubContentRating) Delete(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !wcr._exists {
		return nil
	}

	// if deleted, bail
	if wcr._deleted {
		return nil
	}

	// sql query
	const sqlstr = `DELETE FROM jlabo.wsub_content_rating WHERE content_id = ?`

	// run query
	XOLog(sqlstr, wcr.ContentID)
	_, err = db.Exec(sqlstr, wcr.ContentID)
	if err != nil {
		return err
	}

	// set deleted
	wcr._deleted = true

	return nil
}

// WsubContentRatingByContentID retrieves a row from 'jlabo.wsub_content_rating' as a WsubContentRating.
//
// Generated from index 'wsub_content_rating_content_id_pkey'.
func WsubContentRatingByContentID(db XODB, contentID int) (*WsubContentRating, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`content_id, rating_sum, rating_count, lastip ` +
		`FROM jlabo.wsub_content_rating ` +
		`WHERE content_id = ?`

	// run query
	XOLog(sqlstr, contentID)
	wcr := WsubContentRating{
		_exists: true,
	}

	err = db.QueryRow(sqlstr, contentID).Scan(&wcr.ContentID, &wcr.RatingSum, &wcr.RatingCount, &wcr.Lastip)
	if err != nil {
		return nil, err
	}

	return &wcr, nil
}
