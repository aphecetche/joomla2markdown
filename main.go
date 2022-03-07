package main

import (
	"database/sql"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"time"

	"github.com/aphecetche/joomla2markdown/wsub"
	_ "github.com/go-sql-driver/mysql"
)

func discard(c wsub.Content) bool {
	return c.Category.Title == "theoriehe" || c.Category.Title == "theoriebe"
}

func filter(c []*wsub.Content) []wsub.Content {

	keep := []wsub.Content{}
	for _, r := range c {
		if discard(*r) {
			continue
		}
		keep = append(keep, *r)
	}
	return keep
}

var seminars = flag.Bool("seminars", false, "convert seminars")
var content = flag.Bool("content", false, "convert main content (articles)")

func convertSeminars(db *sql.DB) {
	// Get seminars from DB
	//seminars, err := wsub.Seminars(db, "date > '2020-1-1'")
	seminars, err := wsub.Seminars(db, "true")
	checkErr(err)
	paris, err := time.LoadLocation("Europe/Paris")
	var i = 0
	for _, s := range seminars {
		i++
		filename := s.Date.In(paris).Format("2006-01-02-15-04")
		writeSeminarToFile(s, filename)
		// if i > 10 {
		// 	break
		// }
	}

}

func convertContent(db *sql.DB) {
	// Get categories from DB
	cats, err := wsub.Categories(db)
	checkErr(err)

	var categories = make(map[int]wsub.Category)
	for _, c := range cats {
		categories[c.ID] = *c
	}

	// Get content (articles) from DB
	res, err := wsub.Contents(db, "true", categories)
	checkErr(err)

	// Map ID->content
	var articles = make(map[int]wsub.Content)
	for _, r := range res {
		articles[int(r.ID)] = *r
	}

	// filter out some undesired content, if any
	kept := filter(res)

	// now finally generate all the content files
	for _, r := range kept {
		// files to exclude from the conversion, either because we've got new versions
		// (e.g. in MDX form, or just renamed) or because we already know they are obsolete
		exclude := []string{
			`recherche\/recherche.md`,
			`services-techniques-et-administration\/electronique\/actualit\303\251s`,
			`services-techniques-et-administration\/electronique\/contacts.md`,
			`services-techniques-et-administration\/electronique\/service-electronique.md`,
			`services-techniques-et-administration\/informatique\/service-informatique.md`}

		shouldWrite := true
		for _, t := range exclude {
			var test = regexp.MustCompile(t)
			if test.MatchString(r.FullPath()) {
				shouldWrite = false
				break
			}
		}

		if !shouldWrite {
			continue
		}

		dir := filepath.Join("content", r.DirName())
		os.MkdirAll(dir, os.ModePerm)
		filename := filepath.Join(dir, r.FileName())
		if len(filename) == 0 {
			fmt.Printf("WARNING: skipping a file : %s\n",
				r.String())
			continue
		}
		writeToFile(&r, filename, articles)
	}
}

func main() {
	flag.Parse()

	db, err := sql.Open("mysql", "root:juges3:sud@tcp(localhost:3306)/jlabo?charset=utf8&parseTime=true")
	checkErr(err)
	defer db.Close()

	if *content {
		convertContent(db)
	}
	if *seminars {
		convertSeminars(db)
	}
}

func writeSeminarToFile(s *wsub.Seminar, filename string) {
	path := "seminars/" + s.Date.Format("2006")
	os.MkdirAll(path, os.ModePerm)
	ext := ".fr.mdx"
	if !s.IsTwo() {
		file, err := os.Create(path + "/" + filename + ext)
		defer file.Close()
		checkErr(err)
		s.Write(file, false)
	} else {
		file1, err1 := os.Create(path + "/" + filename + "-1" + ext)
		defer file1.Close()
		checkErr(err1)
		file2, err2 := os.Create(path + "/" + filename + "-2" + ext)
		defer file2.Close()
		checkErr(err2)
		s.Write(file1, false)
		s.Write(file2, true)
	}
}

func writeToFile(c *wsub.Content, filename string, articles map[int]wsub.Content) {
	file, err := os.Create(filename)
	defer file.Close()
	checkErr(err)
	c.Write(file, articles)
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
