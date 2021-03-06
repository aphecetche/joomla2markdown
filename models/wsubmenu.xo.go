// Package toto contains the types for schema 'jlabo'.
package toto

// Code generated by xo. DO NOT EDIT.

import (
	"errors"
	"time"
)

// WsubMenu represents a row from 'jlabo.wsub_menu'.
type WsubMenu struct {
	ID              int       `json:"id"`                // id
	Menutype        string    `json:"menutype"`          // menutype
	Title           string    `json:"title"`             // title
	Alias           string    `json:"alias"`             // alias
	Note            string    `json:"note"`              // note
	Path            string    `json:"path"`              // path
	Link            string    `json:"link"`              // link
	Type            string    `json:"type"`              // type
	Published       int8      `json:"published"`         // published
	ParentID        uint      `json:"parent_id"`         // parent_id
	Level           uint      `json:"level"`             // level
	ComponentID     uint      `json:"component_id"`      // component_id
	CheckedOut      uint      `json:"checked_out"`       // checked_out
	CheckedOutTime  time.Time `json:"checked_out_time"`  // checked_out_time
	Browsernav      int8      `json:"browserNav"`        // browserNav
	Access          uint      `json:"access"`            // access
	Img             string    `json:"img"`               // img
	TemplateStyleID uint      `json:"template_style_id"` // template_style_id
	Params          string    `json:"params"`            // params
	Lft             int       `json:"lft"`               // lft
	Rgt             int       `json:"rgt"`               // rgt
	Home            int8      `json:"home"`              // home
	Language        string    `json:"language"`          // language
	ClientID        int8      `json:"client_id"`         // client_id

	// xo fields
	_exists, _deleted bool
}

// Exists determines if the WsubMenu exists in the database.
func (wm *WsubMenu) Exists() bool {
	return wm._exists
}

// Deleted provides information if the WsubMenu has been deleted from the database.
func (wm *WsubMenu) Deleted() bool {
	return wm._deleted
}

// Insert inserts the WsubMenu to the database.
func (wm *WsubMenu) Insert(db XODB) error {
	var err error

	// if already exist, bail
	if wm._exists {
		return errors.New("insert failed: already exists")
	}

	// sql insert query, primary key provided by autoincrement
	const sqlstr = `INSERT INTO jlabo.wsub_menu (` +
		`menutype, title, alias, note, path, link, type, published, parent_id, level, component_id, checked_out, checked_out_time, browserNav, access, img, template_style_id, params, lft, rgt, home, language, client_id` +
		`) VALUES (` +
		`?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?` +
		`)`

	// run query
	XOLog(sqlstr, wm.Menutype, wm.Title, wm.Alias, wm.Note, wm.Path, wm.Link, wm.Type, wm.Published, wm.ParentID, wm.Level, wm.ComponentID, wm.CheckedOut, wm.CheckedOutTime, wm.Browsernav, wm.Access, wm.Img, wm.TemplateStyleID, wm.Params, wm.Lft, wm.Rgt, wm.Home, wm.Language, wm.ClientID)
	res, err := db.Exec(sqlstr, wm.Menutype, wm.Title, wm.Alias, wm.Note, wm.Path, wm.Link, wm.Type, wm.Published, wm.ParentID, wm.Level, wm.ComponentID, wm.CheckedOut, wm.CheckedOutTime, wm.Browsernav, wm.Access, wm.Img, wm.TemplateStyleID, wm.Params, wm.Lft, wm.Rgt, wm.Home, wm.Language, wm.ClientID)
	if err != nil {
		return err
	}

	// retrieve id
	id, err := res.LastInsertId()
	if err != nil {
		return err
	}

	// set primary key and existence
	wm.ID = int(id)
	wm._exists = true

	return nil
}

// Update updates the WsubMenu in the database.
func (wm *WsubMenu) Update(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !wm._exists {
		return errors.New("update failed: does not exist")
	}

	// if deleted, bail
	if wm._deleted {
		return errors.New("update failed: marked for deletion")
	}

	// sql query
	const sqlstr = `UPDATE jlabo.wsub_menu SET ` +
		`menutype = ?, title = ?, alias = ?, note = ?, path = ?, link = ?, type = ?, published = ?, parent_id = ?, level = ?, component_id = ?, checked_out = ?, checked_out_time = ?, browserNav = ?, access = ?, img = ?, template_style_id = ?, params = ?, lft = ?, rgt = ?, home = ?, language = ?, client_id = ?` +
		` WHERE id = ?`

	// run query
	XOLog(sqlstr, wm.Menutype, wm.Title, wm.Alias, wm.Note, wm.Path, wm.Link, wm.Type, wm.Published, wm.ParentID, wm.Level, wm.ComponentID, wm.CheckedOut, wm.CheckedOutTime, wm.Browsernav, wm.Access, wm.Img, wm.TemplateStyleID, wm.Params, wm.Lft, wm.Rgt, wm.Home, wm.Language, wm.ClientID, wm.ID)
	_, err = db.Exec(sqlstr, wm.Menutype, wm.Title, wm.Alias, wm.Note, wm.Path, wm.Link, wm.Type, wm.Published, wm.ParentID, wm.Level, wm.ComponentID, wm.CheckedOut, wm.CheckedOutTime, wm.Browsernav, wm.Access, wm.Img, wm.TemplateStyleID, wm.Params, wm.Lft, wm.Rgt, wm.Home, wm.Language, wm.ClientID, wm.ID)
	return err
}

// Save saves the WsubMenu to the database.
func (wm *WsubMenu) Save(db XODB) error {
	if wm.Exists() {
		return wm.Update(db)
	}

	return wm.Insert(db)
}

// Delete deletes the WsubMenu from the database.
func (wm *WsubMenu) Delete(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !wm._exists {
		return nil
	}

	// if deleted, bail
	if wm._deleted {
		return nil
	}

	// sql query
	const sqlstr = `DELETE FROM jlabo.wsub_menu WHERE id = ?`

	// run query
	XOLog(sqlstr, wm.ID)
	_, err = db.Exec(sqlstr, wm.ID)
	if err != nil {
		return err
	}

	// set deleted
	wm._deleted = true

	return nil
}

// WsubMenusByAlias retrieves a row from 'jlabo.wsub_menu' as a WsubMenu.
//
// Generated from index 'idx_alias'.
func WsubMenusByAlias(db XODB, alias string) ([]*WsubMenu, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`id, menutype, title, alias, note, path, link, type, published, parent_id, level, component_id, checked_out, checked_out_time, browserNav, access, img, template_style_id, params, lft, rgt, home, language, client_id ` +
		`FROM jlabo.wsub_menu ` +
		`WHERE alias = ?`

	// run query
	XOLog(sqlstr, alias)
	q, err := db.Query(sqlstr, alias)
	if err != nil {
		return nil, err
	}
	defer q.Close()

	// load results
	res := []*WsubMenu{}
	for q.Next() {
		wm := WsubMenu{
			_exists: true,
		}

		// scan
		err = q.Scan(&wm.ID, &wm.Menutype, &wm.Title, &wm.Alias, &wm.Note, &wm.Path, &wm.Link, &wm.Type, &wm.Published, &wm.ParentID, &wm.Level, &wm.ComponentID, &wm.CheckedOut, &wm.CheckedOutTime, &wm.Browsernav, &wm.Access, &wm.Img, &wm.TemplateStyleID, &wm.Params, &wm.Lft, &wm.Rgt, &wm.Home, &wm.Language, &wm.ClientID)
		if err != nil {
			return nil, err
		}

		res = append(res, &wm)
	}

	return res, nil
}

// WsubMenuByClientIDParentIDAliasLanguage retrieves a row from 'jlabo.wsub_menu' as a WsubMenu.
//
// Generated from index 'idx_client_id_parent_id_alias_language'.
func WsubMenuByClientIDParentIDAliasLanguage(db XODB, clientID int8, parentID uint, alias string, language string) (*WsubMenu, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`id, menutype, title, alias, note, path, link, type, published, parent_id, level, component_id, checked_out, checked_out_time, browserNav, access, img, template_style_id, params, lft, rgt, home, language, client_id ` +
		`FROM jlabo.wsub_menu ` +
		`WHERE client_id = ? AND parent_id = ? AND alias = ? AND language = ?`

	// run query
	XOLog(sqlstr, clientID, parentID, alias, language)
	wm := WsubMenu{
		_exists: true,
	}

	err = db.QueryRow(sqlstr, clientID, parentID, alias, language).Scan(&wm.ID, &wm.Menutype, &wm.Title, &wm.Alias, &wm.Note, &wm.Path, &wm.Link, &wm.Type, &wm.Published, &wm.ParentID, &wm.Level, &wm.ComponentID, &wm.CheckedOut, &wm.CheckedOutTime, &wm.Browsernav, &wm.Access, &wm.Img, &wm.TemplateStyleID, &wm.Params, &wm.Lft, &wm.Rgt, &wm.Home, &wm.Language, &wm.ClientID)
	if err != nil {
		return nil, err
	}

	return &wm, nil
}

// WsubMenusByComponentIDMenutypePublishedAccess retrieves a row from 'jlabo.wsub_menu' as a WsubMenu.
//
// Generated from index 'idx_componentid'.
func WsubMenusByComponentIDMenutypePublishedAccess(db XODB, componentID uint, menutype string, published int8, access uint) ([]*WsubMenu, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`id, menutype, title, alias, note, path, link, type, published, parent_id, level, component_id, checked_out, checked_out_time, browserNav, access, img, template_style_id, params, lft, rgt, home, language, client_id ` +
		`FROM jlabo.wsub_menu ` +
		`WHERE component_id = ? AND menutype = ? AND published = ? AND access = ?`

	// run query
	XOLog(sqlstr, componentID, menutype, published, access)
	q, err := db.Query(sqlstr, componentID, menutype, published, access)
	if err != nil {
		return nil, err
	}
	defer q.Close()

	// load results
	res := []*WsubMenu{}
	for q.Next() {
		wm := WsubMenu{
			_exists: true,
		}

		// scan
		err = q.Scan(&wm.ID, &wm.Menutype, &wm.Title, &wm.Alias, &wm.Note, &wm.Path, &wm.Link, &wm.Type, &wm.Published, &wm.ParentID, &wm.Level, &wm.ComponentID, &wm.CheckedOut, &wm.CheckedOutTime, &wm.Browsernav, &wm.Access, &wm.Img, &wm.TemplateStyleID, &wm.Params, &wm.Lft, &wm.Rgt, &wm.Home, &wm.Language, &wm.ClientID)
		if err != nil {
			return nil, err
		}

		res = append(res, &wm)
	}

	return res, nil
}

// WsubMenusByLanguage retrieves a row from 'jlabo.wsub_menu' as a WsubMenu.
//
// Generated from index 'idx_language'.
func WsubMenusByLanguage(db XODB, language string) ([]*WsubMenu, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`id, menutype, title, alias, note, path, link, type, published, parent_id, level, component_id, checked_out, checked_out_time, browserNav, access, img, template_style_id, params, lft, rgt, home, language, client_id ` +
		`FROM jlabo.wsub_menu ` +
		`WHERE language = ?`

	// run query
	XOLog(sqlstr, language)
	q, err := db.Query(sqlstr, language)
	if err != nil {
		return nil, err
	}
	defer q.Close()

	// load results
	res := []*WsubMenu{}
	for q.Next() {
		wm := WsubMenu{
			_exists: true,
		}

		// scan
		err = q.Scan(&wm.ID, &wm.Menutype, &wm.Title, &wm.Alias, &wm.Note, &wm.Path, &wm.Link, &wm.Type, &wm.Published, &wm.ParentID, &wm.Level, &wm.ComponentID, &wm.CheckedOut, &wm.CheckedOutTime, &wm.Browsernav, &wm.Access, &wm.Img, &wm.TemplateStyleID, &wm.Params, &wm.Lft, &wm.Rgt, &wm.Home, &wm.Language, &wm.ClientID)
		if err != nil {
			return nil, err
		}

		res = append(res, &wm)
	}

	return res, nil
}

// WsubMenusByLftRgt retrieves a row from 'jlabo.wsub_menu' as a WsubMenu.
//
// Generated from index 'idx_left_right'.
func WsubMenusByLftRgt(db XODB, lft int, rgt int) ([]*WsubMenu, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`id, menutype, title, alias, note, path, link, type, published, parent_id, level, component_id, checked_out, checked_out_time, browserNav, access, img, template_style_id, params, lft, rgt, home, language, client_id ` +
		`FROM jlabo.wsub_menu ` +
		`WHERE lft = ? AND rgt = ?`

	// run query
	XOLog(sqlstr, lft, rgt)
	q, err := db.Query(sqlstr, lft, rgt)
	if err != nil {
		return nil, err
	}
	defer q.Close()

	// load results
	res := []*WsubMenu{}
	for q.Next() {
		wm := WsubMenu{
			_exists: true,
		}

		// scan
		err = q.Scan(&wm.ID, &wm.Menutype, &wm.Title, &wm.Alias, &wm.Note, &wm.Path, &wm.Link, &wm.Type, &wm.Published, &wm.ParentID, &wm.Level, &wm.ComponentID, &wm.CheckedOut, &wm.CheckedOutTime, &wm.Browsernav, &wm.Access, &wm.Img, &wm.TemplateStyleID, &wm.Params, &wm.Lft, &wm.Rgt, &wm.Home, &wm.Language, &wm.ClientID)
		if err != nil {
			return nil, err
		}

		res = append(res, &wm)
	}

	return res, nil
}

// WsubMenusByMenutype retrieves a row from 'jlabo.wsub_menu' as a WsubMenu.
//
// Generated from index 'idx_menutype'.
func WsubMenusByMenutype(db XODB, menutype string) ([]*WsubMenu, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`id, menutype, title, alias, note, path, link, type, published, parent_id, level, component_id, checked_out, checked_out_time, browserNav, access, img, template_style_id, params, lft, rgt, home, language, client_id ` +
		`FROM jlabo.wsub_menu ` +
		`WHERE menutype = ?`

	// run query
	XOLog(sqlstr, menutype)
	q, err := db.Query(sqlstr, menutype)
	if err != nil {
		return nil, err
	}
	defer q.Close()

	// load results
	res := []*WsubMenu{}
	for q.Next() {
		wm := WsubMenu{
			_exists: true,
		}

		// scan
		err = q.Scan(&wm.ID, &wm.Menutype, &wm.Title, &wm.Alias, &wm.Note, &wm.Path, &wm.Link, &wm.Type, &wm.Published, &wm.ParentID, &wm.Level, &wm.ComponentID, &wm.CheckedOut, &wm.CheckedOutTime, &wm.Browsernav, &wm.Access, &wm.Img, &wm.TemplateStyleID, &wm.Params, &wm.Lft, &wm.Rgt, &wm.Home, &wm.Language, &wm.ClientID)
		if err != nil {
			return nil, err
		}

		res = append(res, &wm)
	}

	return res, nil
}

// WsubMenusByPath retrieves a row from 'jlabo.wsub_menu' as a WsubMenu.
//
// Generated from index 'idx_path'.
func WsubMenusByPath(db XODB, path string) ([]*WsubMenu, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`id, menutype, title, alias, note, path, link, type, published, parent_id, level, component_id, checked_out, checked_out_time, browserNav, access, img, template_style_id, params, lft, rgt, home, language, client_id ` +
		`FROM jlabo.wsub_menu ` +
		`WHERE path = ?`

	// run query
	XOLog(sqlstr, path)
	q, err := db.Query(sqlstr, path)
	if err != nil {
		return nil, err
	}
	defer q.Close()

	// load results
	res := []*WsubMenu{}
	for q.Next() {
		wm := WsubMenu{
			_exists: true,
		}

		// scan
		err = q.Scan(&wm.ID, &wm.Menutype, &wm.Title, &wm.Alias, &wm.Note, &wm.Path, &wm.Link, &wm.Type, &wm.Published, &wm.ParentID, &wm.Level, &wm.ComponentID, &wm.CheckedOut, &wm.CheckedOutTime, &wm.Browsernav, &wm.Access, &wm.Img, &wm.TemplateStyleID, &wm.Params, &wm.Lft, &wm.Rgt, &wm.Home, &wm.Language, &wm.ClientID)
		if err != nil {
			return nil, err
		}

		res = append(res, &wm)
	}

	return res, nil
}

// WsubMenuByID retrieves a row from 'jlabo.wsub_menu' as a WsubMenu.
//
// Generated from index 'wsub_menu_id_pkey'.
func WsubMenuByID(db XODB, id int) (*WsubMenu, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`id, menutype, title, alias, note, path, link, type, published, parent_id, level, component_id, checked_out, checked_out_time, browserNav, access, img, template_style_id, params, lft, rgt, home, language, client_id ` +
		`FROM jlabo.wsub_menu ` +
		`WHERE id = ?`

	// run query
	XOLog(sqlstr, id)
	wm := WsubMenu{
		_exists: true,
	}

	err = db.QueryRow(sqlstr, id).Scan(&wm.ID, &wm.Menutype, &wm.Title, &wm.Alias, &wm.Note, &wm.Path, &wm.Link, &wm.Type, &wm.Published, &wm.ParentID, &wm.Level, &wm.ComponentID, &wm.CheckedOut, &wm.CheckedOutTime, &wm.Browsernav, &wm.Access, &wm.Img, &wm.TemplateStyleID, &wm.Params, &wm.Lft, &wm.Rgt, &wm.Home, &wm.Language, &wm.ClientID)
	if err != nil {
		return nil, err
	}

	return &wm, nil
}
