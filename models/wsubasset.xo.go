// Package toto contains the types for schema 'jlabo'.
package toto

// Code generated by xo. DO NOT EDIT.

import (
	"errors"
)

// WsubAsset represents a row from 'jlabo.wsub_assets'.
type WsubAsset struct {
	ID       uint   `json:"id"`        // id
	ParentID int    `json:"parent_id"` // parent_id
	Lft      int    `json:"lft"`       // lft
	Rgt      int    `json:"rgt"`       // rgt
	Level    uint   `json:"level"`     // level
	Name     string `json:"name"`      // name
	Title    string `json:"title"`     // title
	Rules    string `json:"rules"`     // rules

	// xo fields
	_exists, _deleted bool
}

// Exists determines if the WsubAsset exists in the database.
func (wa *WsubAsset) Exists() bool {
	return wa._exists
}

// Deleted provides information if the WsubAsset has been deleted from the database.
func (wa *WsubAsset) Deleted() bool {
	return wa._deleted
}

// Insert inserts the WsubAsset to the database.
func (wa *WsubAsset) Insert(db XODB) error {
	var err error

	// if already exist, bail
	if wa._exists {
		return errors.New("insert failed: already exists")
	}

	// sql insert query, primary key provided by autoincrement
	const sqlstr = `INSERT INTO jlabo.wsub_assets (` +
		`parent_id, lft, rgt, level, name, title, rules` +
		`) VALUES (` +
		`?, ?, ?, ?, ?, ?, ?` +
		`)`

	// run query
	XOLog(sqlstr, wa.ParentID, wa.Lft, wa.Rgt, wa.Level, wa.Name, wa.Title, wa.Rules)
	res, err := db.Exec(sqlstr, wa.ParentID, wa.Lft, wa.Rgt, wa.Level, wa.Name, wa.Title, wa.Rules)
	if err != nil {
		return err
	}

	// retrieve id
	id, err := res.LastInsertId()
	if err != nil {
		return err
	}

	// set primary key and existence
	wa.ID = uint(id)
	wa._exists = true

	return nil
}

// Update updates the WsubAsset in the database.
func (wa *WsubAsset) Update(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !wa._exists {
		return errors.New("update failed: does not exist")
	}

	// if deleted, bail
	if wa._deleted {
		return errors.New("update failed: marked for deletion")
	}

	// sql query
	const sqlstr = `UPDATE jlabo.wsub_assets SET ` +
		`parent_id = ?, lft = ?, rgt = ?, level = ?, name = ?, title = ?, rules = ?` +
		` WHERE id = ?`

	// run query
	XOLog(sqlstr, wa.ParentID, wa.Lft, wa.Rgt, wa.Level, wa.Name, wa.Title, wa.Rules, wa.ID)
	_, err = db.Exec(sqlstr, wa.ParentID, wa.Lft, wa.Rgt, wa.Level, wa.Name, wa.Title, wa.Rules, wa.ID)
	return err
}

// Save saves the WsubAsset to the database.
func (wa *WsubAsset) Save(db XODB) error {
	if wa.Exists() {
		return wa.Update(db)
	}

	return wa.Insert(db)
}

// Delete deletes the WsubAsset from the database.
func (wa *WsubAsset) Delete(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !wa._exists {
		return nil
	}

	// if deleted, bail
	if wa._deleted {
		return nil
	}

	// sql query
	const sqlstr = `DELETE FROM jlabo.wsub_assets WHERE id = ?`

	// run query
	XOLog(sqlstr, wa.ID)
	_, err = db.Exec(sqlstr, wa.ID)
	if err != nil {
		return err
	}

	// set deleted
	wa._deleted = true

	return nil
}

// WsubAssetByName retrieves a row from 'jlabo.wsub_assets' as a WsubAsset.
//
// Generated from index 'idx_asset_name'.
func WsubAssetByName(db XODB, name string) (*WsubAsset, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`id, parent_id, lft, rgt, level, name, title, rules ` +
		`FROM jlabo.wsub_assets ` +
		`WHERE name = ?`

	// run query
	XOLog(sqlstr, name)
	wa := WsubAsset{
		_exists: true,
	}

	err = db.QueryRow(sqlstr, name).Scan(&wa.ID, &wa.ParentID, &wa.Lft, &wa.Rgt, &wa.Level, &wa.Name, &wa.Title, &wa.Rules)
	if err != nil {
		return nil, err
	}

	return &wa, nil
}

// WsubAssetsByLftRgt retrieves a row from 'jlabo.wsub_assets' as a WsubAsset.
//
// Generated from index 'idx_lft_rgt'.
func WsubAssetsByLftRgt(db XODB, lft int, rgt int) ([]*WsubAsset, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`id, parent_id, lft, rgt, level, name, title, rules ` +
		`FROM jlabo.wsub_assets ` +
		`WHERE lft = ? AND rgt = ?`

	// run query
	XOLog(sqlstr, lft, rgt)
	q, err := db.Query(sqlstr, lft, rgt)
	if err != nil {
		return nil, err
	}
	defer q.Close()

	// load results
	res := []*WsubAsset{}
	for q.Next() {
		wa := WsubAsset{
			_exists: true,
		}

		// scan
		err = q.Scan(&wa.ID, &wa.ParentID, &wa.Lft, &wa.Rgt, &wa.Level, &wa.Name, &wa.Title, &wa.Rules)
		if err != nil {
			return nil, err
		}

		res = append(res, &wa)
	}

	return res, nil
}

// WsubAssetsByParentID retrieves a row from 'jlabo.wsub_assets' as a WsubAsset.
//
// Generated from index 'idx_parent_id'.
func WsubAssetsByParentID(db XODB, parentID int) ([]*WsubAsset, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`id, parent_id, lft, rgt, level, name, title, rules ` +
		`FROM jlabo.wsub_assets ` +
		`WHERE parent_id = ?`

	// run query
	XOLog(sqlstr, parentID)
	q, err := db.Query(sqlstr, parentID)
	if err != nil {
		return nil, err
	}
	defer q.Close()

	// load results
	res := []*WsubAsset{}
	for q.Next() {
		wa := WsubAsset{
			_exists: true,
		}

		// scan
		err = q.Scan(&wa.ID, &wa.ParentID, &wa.Lft, &wa.Rgt, &wa.Level, &wa.Name, &wa.Title, &wa.Rules)
		if err != nil {
			return nil, err
		}

		res = append(res, &wa)
	}

	return res, nil
}

// WsubAssetByID retrieves a row from 'jlabo.wsub_assets' as a WsubAsset.
//
// Generated from index 'wsub_assets_id_pkey'.
func WsubAssetByID(db XODB, id uint) (*WsubAsset, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`id, parent_id, lft, rgt, level, name, title, rules ` +
		`FROM jlabo.wsub_assets ` +
		`WHERE id = ?`

	// run query
	XOLog(sqlstr, id)
	wa := WsubAsset{
		_exists: true,
	}

	err = db.QueryRow(sqlstr, id).Scan(&wa.ID, &wa.ParentID, &wa.Lft, &wa.Rgt, &wa.Level, &wa.Name, &wa.Title, &wa.Rules)
	if err != nil {
		return nil, err
	}

	return &wa, nil
}
