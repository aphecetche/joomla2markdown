// Package toto contains the types for schema 'jlabo'.
package wsub

import (
	"database/sql"
	"fmt"
	"time"
)

// WsubCategory represents a row from 'jlabo.wsub_categories'.
type Category struct {
	ID             int       `json:"id"`               // id
	AssetID        uint      `json:"asset_id"`         // asset_id
	ParentID       uint      `json:"parent_id"`        // parent_id
	Lft            int       `json:"lft"`              // lft
	Rgt            int       `json:"rgt"`              // rgt
	Level          uint      `json:"level"`            // level
	Path           string    `json:"path"`             // path
	Extension      string    `json:"extension"`        // extension
	Title          string    `json:"title"`            // title
	Alias          string    `json:"alias"`            // alias
	Note           string    `json:"note"`             // note
	Description    string    `json:"description"`      // description
	Published      bool      `json:"published"`        // published
	CheckedOut     uint      `json:"checked_out"`      // checked_out
	CheckedOutTime time.Time `json:"checked_out_time"` // checked_out_time
	Access         uint      `json:"access"`           // access
	Params         string    `json:"params"`           // params
	Metadesc       string    `json:"metadesc"`         // metadesc
	Metakey        string    `json:"metakey"`          // metakey
	Metadata       string    `json:"metadata"`         // metadata
	CreatedUserID  uint      `json:"created_user_id"`  // created_user_id
	CreatedTime    time.Time `json:"created_time"`     // created_time
	ModifiedUserID uint      `json:"modified_user_id"` // modified_user_id
	ModifiedTime   time.Time `json:"modified_time"`    // modified_time
	Hits           uint      `json:"hits"`             // hits
	Language       string    `json:"language"`         // language
	Version        uint      `json:"version"`          // version

}

func (c Category) String() string {
	return fmt.Sprintf("id %d path %s title %s alias %s ",
		c.ID, c.Path, c.Title, c.Alias)
}

// Categories retrieves a row from 'jlabo.wsub_categories' as a WsubCategory.
func Categories(db *sql.DB) ([]*Category, error) {
	var err error

	qs := `SELECT ` +
		`id, path, title, alias ` +
		`FROM jlabo.wsub_categories `

	// run query
	q, err := db.Query(qs)
	if err != nil {
		return nil, err
	}
	defer q.Close()

	// load results
	res := []*Category{}
	for q.Next() {
		wc := Category{}
		// scan
		err = q.Scan(&wc.ID, &wc.Path, &wc.Title, &wc.Alias)
		if err != nil {
			return nil, err
		}

		res = append(res, &wc)
	}

	return res, nil
}
