// Package wsub contains the types for schema 'jlabo'.
package wsub

import (
	"database/sql"
	"fmt"
	"time"
)

// Seminar represents a row from 'jlabo.wsub_seminars'.
type Seminar struct {
	ID                 int64     `json:"id"`                   // id
	Author             string    `json:"author"`               // author
	AuthorURL          string    `json:"author_url"`           // author_url
	Date               time.Time `json:"date"`                 // date
	Location           string    `json:"location"`             // location
	Alias              string    `json:"alias"`                // alias
	Summary            string    `json:"summary"`              // summary
	CheckedOut         int       `json:"checked_out"`          // checked_out
	CheckedOutTime     time.Time `json:"checked_out_time"`     // checked_out_time
	Title              string    `json:"title"`                // title
	AuthorFiliation    string    `json:"author_filiation"`     // author_filiation
	AuthorFiliationURL string    `json:"author_filiation_url"` // author_filiation_url
	Type               string    `json:"type"`                 // type
	Comment            string    `json:"comment"`              // comment
	State              int8      `json:"state"`                // state
	Created            time.Time `json:"created"`              // created
	CreatedBy          int       `json:"created_by"`           // created_by
	Modified           time.Time `json:"modified"`             // modified
	ModifiedBy         int       `json:"modified_by"`          // modified_by
	PublishUp          time.Time `json:"publish_up"`           // publish_up
	PublishDown        time.Time `json:"publish_down"`         // publish_down
	Attribs            string    `json:"attribs"`              // attribs

	// xo fields
	_exists, _deleted bool
}

func (w Seminar) String() string {
	return fmt.Sprintf("%s:%s", w.Title, w.Type)
}

// SeminarsByAlias retrieves a row from 'jlabo.wsub_seminars' as a Seminar.
//
// Generated from index 'alias'.
func Seminars(db *sql.DB, where string) ([]*Seminar, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`id, author, author_url, date, location, alias, summary, checked_out, checked_out_time, title, author_filiation, author_filiation_url, type, comment, state, created, created_by, modified, modified_by, publish_up, publish_down, attribs ` +
		`FROM jlabo.wsub_seminars `

	qs := sqlstr + `WHERE ` + where

	// run query
	q, err := db.Query(qs)
	if err != nil {
		return nil, err
	}
	defer q.Close()

	// load results
	res := []*Seminar{}
	for q.Next() {
		ws := Seminar{
			_exists: true,
		}

		// scan
		err = q.Scan(&ws.ID, &ws.Author, &ws.AuthorURL, &ws.Date, &ws.Location, &ws.Alias, &ws.Summary, &ws.CheckedOut, &ws.CheckedOutTime, &ws.Title, &ws.AuthorFiliation, &ws.AuthorFiliationURL, &ws.Type, &ws.Comment, &ws.State, &ws.Created, &ws.CreatedBy, &ws.Modified, &ws.ModifiedBy, &ws.PublishUp, &ws.PublishDown, &ws.Attribs)
		if err != nil {
			return nil, err
		}

		res = append(res, &ws)
	}

	for _, s := range res {
		fmt.Println(s)
	}

	return res, nil
}
