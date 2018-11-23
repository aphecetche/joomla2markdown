package main

import (
	"database/sql"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/aphecetche/joomla2hugo/wsub"
	_ "github.com/go-sql-driver/mysql"
)

func discard(c wsub.Content) bool {
	// return c.Category.Title == "theoriehe" || c.Category.Title == "theoriebe" || c.Category.Title != "Administration"
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

// remaining returns list of menu paths that were not already used
func remaining(menus []*wsub.Menu, usedmenus map[string]struct{}) []string {
	rem := make(map[string]struct{})
	for _, m := range menus {
		print := false
		if strings.Contains(m.Path, "astro") {
			print = true
		}
		if print {
			fmt.Println("path=", m.Path)
		}
		sm := splitMenu(m.Path)
		for i, mpath := range sm {
			if print {
				fmt.Println("---", i, mpath)
			}
			_, ok := usedmenus[mpath]
			if !ok {
				rem[mpath] = struct{}{}
			}
		}
	}
	rv := make([]string, len(rem))
	n := 0
	for i := range rem {
		rv[n] = i
		n++
	}
	return rv
}

func splitMenu(mpath string) []string {
	mmap := make(map[string]struct{})
	s := strings.Split(mpath, "/")
	for i := 1; i < len(s); i++ {
		path := strings.Join(s[0:i], "/")
		mmap[path] = struct{}{}
	}
	rv := make([]string, len(mmap))

	n := 0
	for i := range mmap {
		rv[n] = i
		n++
	}
	return rv
}

func main() {
	db, err := sql.Open("mysql", "root:juges3:sud@tcp(localhost:3306)/jlabo?charset=utf8&parseTime=true")
	checkErr(err)
	defer db.Close()

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

	// Get menus from DB
	menus, err := wsub.Menus(db, "menutype='menu-principal-fr'")

	// should use articles map *and* the menu-principal-fr to
	// try to find orphan articles in order to remove them...
	checkErr(err)
	links := wsub.ListLinks(articles, menus)

	// filter out some undesired content, if any
	kept := filter(res)

	// update the MenuPath and link part of content
	for i := range kept {
		wsub.UpdateMenuField(&kept[i], menus)
		wsub.UpdateLinkField(&kept[i], links)
	}

	usedmenus := make(map[string]struct{})

	// now finally generate all the content files
	for _, r := range kept {
		dir := filepath.Join("content", r.DirName())
		os.MkdirAll(dir, os.ModePerm)
		filename := filepath.Join(dir, r.FileName())
		writeToFile(&r, filename)
		usedmenus[r.MenuPath] = struct{}{}
	}

	// finish by creating a menu configuration file for the remaining plumbing
	file, err := os.Create("menu.toml")
	defer file.Close()
	wsub.GenerateNonContentMenus(remaining(menus, usedmenus), file)
}

func writeToFile(c *wsub.Content, filename string) {
	file, err := os.Create(filename)
	defer file.Close()
	checkErr(err)
	c.Write(file)
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
