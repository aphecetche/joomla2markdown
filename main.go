package main

import (
	"database/sql"
	"os"
	"path/filepath"

	"github.com/aphecetche/joomla2hugo/wsub"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:juges3:sud@tcp(localhost:3306)/jlabo?charset=utf8&parseTime=true")
	checkErr(err)
	defer db.Close()

	cats, err := wsub.Categories(db)
	checkErr(err)
	var categories = make(map[int]wsub.Category)
	for _, c := range cats {
		categories[c.ID] = *c
	}

	//res, err := wsub.Contents(db, "(catid=17 or catid=16)", categories)
	res, err := wsub.Contents(db, "true", categories)
	checkErr(err)
	for _, r := range res {
		dir := filepath.Join("output", r.DirName())
		os.MkdirAll(dir, os.ModePerm)
		filename := filepath.Join(dir, r.FileName())
		file, err := os.Create(filename)
		checkErr(err)
		r.Write(file)
	}
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
