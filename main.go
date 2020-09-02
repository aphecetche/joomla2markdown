package main

import (
	"database/sql"
	"flag"
	"os"
	"path/filepath"

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
	_, err := wsub.Seminars(db, "true")
	checkErr(err)

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
		dir := filepath.Join("content", r.DirName())
		os.MkdirAll(dir, os.ModePerm)
		filename := filepath.Join(dir, r.FileName())
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
