package main

import (
	"database/sql"
	"fmt"

	"github.com/aphecetche/joomla2hugo/wsub"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:juges3:sud@tcp(localhost:3306)/jlabo?charset=utf8&parseTime=true")
	checkErr(err)
	defer db.Close()

	res, err := wsub.WsubContents(db, "catid=17")
	checkErr(err)
	for _, r := range res {
		fmt.Println(r)
	}
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
