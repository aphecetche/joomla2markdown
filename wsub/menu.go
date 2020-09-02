// Package toto contains the types for schema 'jlabo'.
package wsub

// Code generated by xo. DO NOT EDIT.

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

// WsubMenu represents a row from 'jlabo.wsub_menu'.
type Menu struct {
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
}

func (m Menu) String() string {

	return fmt.Sprintf("id %4d type:%s title:%s path:%s link:%s type:%s published:%d",
		m.ID, m.Menutype, m.Title, m.Path, m.Link, m.Type, m.Published)
}

func findMenu(menus []*Menu, id int) *Menu {
	for _, m := range menus {
		if m.ID == id {
			return m
		}
	}
	return nil
}

func decodeAliases(menus []*Menu) []*Menu {
	var rv []*Menu
	for _, m := range menus {
		if m.Type == "alias" {
			var v interface{}
			json.Unmarshal([]byte(m.Params), &v)
			jm := v.(map[string]interface{})
			id, _ := strconv.Atoi(jm["aliasoptions"].(string))
			d := findMenu(menus, id)
			m.Params = d.Path
		}
		rv = append(rv, m)
	}
	return rv
}

// WsubMenusByAlias retrieves a row from 'jlabo.wsub_menu' as a WsubMenu.
//
// Generated from index 'idx_alias'.
func Menus(db *sql.DB, where string) ([]*Menu, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`id, menutype, title, alias, path, link, type, published, access, level, params ` +
		`FROM jlabo.wsub_menu `

	qs := sqlstr + " WHERE " + where
	// run query
	q, err := db.Query(qs)
	if err != nil {
		return nil, err
	}
	defer q.Close()

	// load results
	res := []*Menu{}
	for q.Next() {
		wm := Menu{}

		// scan
		err = q.Scan(&wm.ID, &wm.Menutype, &wm.Title, &wm.Alias, &wm.Path, &wm.Link,
			&wm.Type, &wm.Published, &wm.Access, &wm.Level, &wm.Params)
		if err != nil {
			return nil, err
		}

		res = append(res, &wm)
	}

	return decodeAliases(res), nil
}

func hasRightMenu(m *Menu) bool {
	if strings.Contains(m.Path, "/recherche/univers-a-haute-energie") {
		return true
	}
	return false
}

func GenerateNonContentMenus(menus []*Menu, out io.Writer) {
	for _, m := range menus {
		writeMenu(out, *m, " [[ menu.main ]]")
	}
}

func writeMenu(out io.Writer, menu Menu, menuName string) {
	if len(menu.Path) == 0 {
		// fmt.Printf("got no menupath")
		return
	}
	if menu.Path == "unknown" {
		return
	}

	fmt.Fprintf(out, "  %s:\n", menuName)
	fmt.Fprintf(out, "    identifier: \"%s\"\n", menu.Path)
	fmt.Fprintf(out, "    name: \"%s\"\n", menu.Title)
	if menu.Type == "alias" {
		fmt.Fprintf(out, "    url: \"%s\"\n", menu.Params)
	}
	dir := filepath.Dir(menu.Path)
	if dir != "." {
		fmt.Fprintf(out, "    parent: \"%s\"\n", dir)
	}
}
