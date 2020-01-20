package main

import (
	"database/sql"
	"os"
	"path/filepath"
	"strings"

	"github.com/aphecetche/joomla2markdown/wsub"
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

// remaining returns list of menus that were not already used
func remaining(menus []*wsub.Menu, usedmenus map[string]struct{}) []*wsub.Menu {
	rem := make(map[string]*wsub.Menu)
	for _, m := range menus {
		sm := splitMenu(m.Path)
		for _, mpath := range sm {
			_, ok := usedmenus[mpath]
			if !ok {
				rem[mpath] = m
			}
		}
	}
	rv := make([]*wsub.Menu, len(rem))
	n := 0
	for _, i := range rem {
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

	usedmenus := make(map[int]struct{})

	// update the MenuPath and link part of content
	for i := range kept {
		im := wsub.UpdateMenuField(&kept[i], menus)
		if im > 0 {
			usedmenus[im] = struct{}{}
		}
		wsub.UpdateLinkField(&kept[i], links)
	}

	// now finally generate all the content files
	for _, r := range kept {
		dir := filepath.Join("content", r.DirName())
		os.MkdirAll(dir, os.ModePerm)
		filename := filepath.Join(dir, r.FileName())
		writeToFile(&r, filename)
	}

	nused := 0
	var notused []*wsub.Menu

	for _, m := range menus {
		_, ok := usedmenus[m.ID]
		if ok {
			nused++
			// fmt.Printf("   USED %s : %s ACCESS: %d\n", m.Path, m.Title, m.Access)
		} else {
			// fmt.Printf("NOTUSED %s : %s ACCESS: %d\n", m.Path, m.Title, m.Access)
			if m.Access == 1 {
				notused = append(notused, m)
			}
		}
	}

	// fmt.Printf("nused %d total %d\n", nused, len(menus))

	//finish by creating a menu configuration file for the remaining plumbing
	file, err := os.Create("menu.toml")
	defer file.Close()
	wsub.GenerateNonContentMenus(notused, file)
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
