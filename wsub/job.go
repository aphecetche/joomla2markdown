// Package toto contains the types for schema 'jlabo'.
package wsub

// Code generated by xo. DO NOT EDIT.

import (
	"database/sql"
	"fmt"
	"time"
	"io"

	"github.com/lunny/html2md"
)

// WsubJob represents a row from 'jlabo.wsub_jobs'.
type Job struct {
	ID              int64          `json:"id"`                // id
	Type            string         `json:"type"`              // type
	Alias           string         `json:"alias"`             // alias
	DateStart       time.Time      `json:"date_start"`        // date_start
	DateReply       time.Time      `json:"date_reply"`        // date_reply
	DescriptionFrFr string         `json:"description_fr_FR"` // description_fr_FR
	DescriptionEnGb string         `json:"description_en_GB"` // description_en_GB
	CheckedOut      int            `json:"checked_out"`       // checked_out
	CheckedOutTime  time.Time      `json:"checked_out_time"`  // checked_out_time
	TitleFrFr       string         `json:"title_fr_FR"`       // title_fr_FR
	TitleEnGb       string         `json:"title_en_GB"`       // title_en_GB
	State           int8           `json:"state"`             // state
	Created         time.Time      `json:"created"`           // created
	CreatedBy       int            `json:"created_by"`        // created_by
	Modified        time.Time      `json:"modified"`          // modified
	ModifiedBy      int            `json:"modified_by"`       // modified_by
	PublishUp       time.Time      `json:"publish_up"`        // publish_up
	PublishDown     time.Time      `json:"publish_down"`      // publish_down
	Attribs         string         `json:"attribs"`           // attribs
	MiscFrFr        string         `json:"misc_fr_FR"`        // misc_fr_FR
	MiscEnGb        string         `json:"misc_en_GB"`        // misc_en_GB
	Logo1           sql.NullString `json:"logo1"`             // logo1
	Logo2           sql.NullString `json:"logo2"`             // logo2
	Logo3           sql.NullString `json:"logo3"`             // logo3
	Logo4           sql.NullString `json:"logo4"`             // logo4
	Logo5           sql.NullString `json:"logo5"`             // logo5
	Group           sql.NullString `json:"group"`             // group
	Keywords        sql.NullString `json:"keywords"`          // keywords

	// xo fields
	_exists, _deleted bool
}

// Exists determines if the WsubJob exists in the database.
func (job *Job) Exists() bool {
	return job._exists
}

func (job Job) String() string {
	return fmt.Sprintf("[%s]:%s:%s", job.Alias, job.TitleFrFr, job.Type)
}

// JobsByAlias retrieves a row from 'jlabo.wsub_jobs' as a Job
//
// Generated from index 'alias'.
func Jobs(db* sql.DB, where string) ([]*Job, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`id, type, alias, date_start, date_reply, description_fr_FR, description_en_GB, checked_out, checked_out_time, title_fr_FR, title_en_GB, state, created, created_by, modified, modified_by, publish_up, publish_down, attribs, misc_fr_FR, misc_en_GB, logo1, logo2, logo3, logo4, logo5, jlabo.wsub_jobs.group, keywords ` +
		`FROM jlabo.wsub_jobs `

	// run query
	qs := sqlstr + `WHERE `+where
	q, err := db.Query(qs)
	if err != nil {
		return nil, err
	}
	defer q.Close()

	// load results
	res := []*Job{}
	for q.Next() {
		wj := Job{
			_exists: true,
		}

		// scan
		err = q.Scan(&wj.ID, &wj.Type, &wj.Alias, &wj.DateStart, &wj.DateReply, &wj.DescriptionFrFr, &wj.DescriptionEnGb, &wj.CheckedOut, &wj.CheckedOutTime, &wj.TitleFrFr, &wj.TitleEnGb, &wj.State, &wj.Created, &wj.CreatedBy, &wj.Modified, &wj.ModifiedBy, &wj.PublishUp, &wj.PublishDown, &wj.Attribs, &wj.MiscFrFr, &wj.MiscEnGb, &wj.Logo1, &wj.Logo2, &wj.Logo3, &wj.Logo4, &wj.Logo5, &wj.Group, &wj.Keywords)
		if err != nil {
			return nil, err
		}

		res = append(res, &wj)
	}

	return res, nil
}

func (job Job) Write(out io.Writer, french bool) {
	fmt.Fprintln(out, "---")
	fmt.Fprintf(out, "date: \"%s\"\n", job.DateStart)
	fmt.Fprintf(out, "type: \"%s\"\n", job.Type)
	//fmt.Fprintf(out, "alias: \"%s\"\n", job.Alias)
	if (job.Group.Valid && len(job.Group.String)>0) {
	fmt.Fprintf(out, "team: \"%s\"\n", job.Group.String)
	}
	if (job.Logo1.Valid && len(job.Logo1.String)>0) {
	fmt.Fprintf(out, "logo1: \"%s\"\n", job.Logo1.String)
	}
	if (job.Logo2.Valid && len(job.Logo2.String)>0) {
	fmt.Fprintf(out, "logo2: \"%s\"\n", job.Logo2.String)
	}
	if (job.Logo3.Valid && len(job.Logo3.String)>0) {
	fmt.Fprintf(out, "logo3: \"%s\"\n", job.Logo3.String)
	}
	if (job.Logo4.Valid && len(job.Logo4.String)>0) {
	fmt.Fprintf(out, "logo4: \"%s\"\n", job.Logo4.String)
	}
	if (job.Logo5.Valid && len(job.Logo5.String)>0) {
	fmt.Fprintf(out, "logo5: \"%s\"\n", job.Logo5.String)
	}
	if (french) {
	fmt.Fprintf(out, "title: \"%s\"\n", job.TitleFrFr)
	if (len(job.MiscFrFr)>0) {
	fmt.Fprintf(out, "misc: \"%s\"\n", job.MiscFrFr)
	}
	} else {
	fmt.Fprintf(out, "title: \"%s\"\n", job.TitleEnGb)
	if (len(job.MiscEnGb)>0) {
	fmt.Fprintf(out, "misc: \"%s\"\n", job.MiscEnGb)
	}
	}
	fmt.Fprintln(out, "---")
	if (french) {
	fmt.Fprintln(out, html2md.Convert(cleanupHTML(job.DescriptionFrFr)))
	} else {
	fmt.Fprintln(out, html2md.Convert(cleanupHTML(job.DescriptionEnGb)))
	}
}
