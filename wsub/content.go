// Package toto contains the types for schema 'jlabo'.
package wsub

import (
	"bytes"
	"database/sql"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/url"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"time"

	"golang.org/x/net/html"
)

// Content represents a row from 'jlabo._content'.
type Content struct {
	ID             uint      `json:"id"`               // id
	AssetID        uint      `json:"asset_id"`         // asset_id
	Title          string    `json:"title"`            // title
	Alias          string    `json:"alias"`            // alias
	Introtext      string    `json:"introtext"`        // introtext
	Fulltext       string    `json:"fulltext"`         // fulltext
	State          int8      `json:"state"`            // state
	Catid          uint      `json:"catid"`            // catid
	Created        time.Time `json:"created"`          // created
	CreatedBy      uint      `json:"created_by"`       // created_by
	CreatedByAlias string    `json:"created_by_alias"` // created_by_alias
	Modified       time.Time `json:"modified"`         // modified
	ModifiedBy     uint      `json:"modified_by"`      // modified_by
	CheckedOut     uint      `json:"checked_out"`      // checked_out
	CheckedOutTime time.Time `json:"checked_out_time"` // checked_out_time
	PublishUp      time.Time `json:"publish_up"`       // publish_up
	PublishDown    time.Time `json:"publish_down"`     // publish_down
	Images         string    `json:"images"`           // images
	Urls           string    `json:"urls"`             // urls
	Attribs        string    `json:"attribs"`          // attribs
	Version        uint      `json:"version"`          // version
	Ordering       int       `json:"ordering"`         // ordering
	Metakey        string    `json:"metakey"`          // metakey
	Metadesc       string    `json:"metadesc"`         // metadesc
	Access         uint      `json:"access"`           // access
	Hits           uint      `json:"hits"`             // hits
	Metadata       string    `json:"metadata"`         // metadata
	Featured       int8      `json:"featured"`         // featured
	Language       string    `json:"language"`         // language
	Xreference     string    `json:"xreference"`       // xreference

	Category Category
}

func (w Content) String() string {
	return fmt.Sprintf("%3d:Title[%s]:Alias[%s]:Created[%v]:Modified[%v]:File[%v]", w.ID, w.Title,
		w.Alias, w.Created.Format("2 janvier 2006"), w.Modified.Format("2 janvier 2006"), w.FileName(),
	)
}

func stringReplace(s, from, dest string) string {
	return strings.Replace(s, from, dest, -1)
}

// changeAccent replaces &acute; and so on by their corresponding characters
func changeAccents(s string) string {
	s = stringReplace(s, "&eacute;", "é")
	s = stringReplace(s, "&agrave;", "à")
	s = stringReplace(s, "&egrave;", "è")
	s = stringReplace(s, "&#39;", "'")
	s = stringReplace(s, "&#34;", "\"")
	return s
}

func render(n *html.Node) string {
	out := bytes.NewBufferString("")
	html.Render(out, n)
	s := out.String()
	s = stringReplace(s, "<html><head></head>", "")
	s = stringReplace(s, "<body>", "")
	s = stringReplace(s, "</body></html>", "")
	return strings.TrimSpace(s)

}

// removeStyle removes all the inline styles found in html elements
// (e.g. in p or img)
func removeStyle(s string) string {
	r := strings.NewReader(s)
	doc, err := html.Parse(r)
	if err != nil {
		return "error"
	}
	var f func(*html.Node)
	f = func(n *html.Node) {
		if n.Type == html.ElementNode {
			attrs := []html.Attribute{}
			for _, a := range n.Attr {
				if !strings.Contains(a.Key, "style") {
					attrs = append(attrs, a)
				}
			}
			n.Attr = attrs
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}
	f(doc)
	return render(doc)
}

func href(t html.Token) string {
	for _, a := range t.Attr {
		if a.Key == "href" {
			return a.Val
		}
	}
	return ""
}

func img(t html.Token) string {
	for _, a := range t.Attr {
		if a.Key == "img" {
			return a.Val
		}
	}
	return ""
}

func articleId(href string) int {
	u, err := url.Parse(href)
	if err != nil {
		panic(err)
	}
	q := u.Query()
	i := q["id"]
	if len(i) == 0 {
		fmt.Println("ERROR - could not decode link", href)
		return -1
	}
	said := strings.Split(i[0], ":")
	aid, err := strconv.Atoi(said[0])
	if err != nil {
		panic(err)
	}
	return aid
}

func massageImages(s string) string {
	r := strings.NewReplacer(
		"(images/Recherche/Erdre", "(images",
		"(images/Recherche/Astro", "(images",
		"(images/Recherche/Plasma", "(images",
		"(images/Recherche/Prisma", "(images",
		"(images/Recherche/Radiochimie", "(images",
		"(images/Recherche/TheorieBE", "(images",
		"(images/Recherche/TheorieHE", "(images",
		"(images/Recherche/Xenon", "(images",
		"(images/Recherche/neutrino", "(images",
		"(images/Enseignement", "(images",
		"(images/Recherche", "(images",
		"(images/banners", "(images")
	return r.Replace(s)
}

func massageDocumentLinks(s string) string {
	r := strings.NewReader(s)
	z := html.NewTokenizer(r)

	documentDir := "doc"

	imre := strings.NewReplacer(
		"images/Enseignement", documentDir,
		"images/Recherche/Plasma", documentDir,
		"images/Communication", documentDir,
		"images/Administration", documentDir)

	var pdfLink = regexp.MustCompile(`images\/.*\.pdf`)
	var externalLink = regexp.MustCompile(`^http`)

	rep := make(map[string]string)

out:
	for {
		tt := z.Next()
		switch {
		case tt == html.ErrorToken:
			// end of document
			break out
		case tt == html.StartTagToken:
			t := z.Token()
			if t.Data == "a" {
				h := href(t)
				if externalLink.MatchString(h) {
					break
				}
				fmt.Println(h, "match=", pdfLink.MatchString(h))
				if pdfLink.MatchString(h) {
					rep[h] = imre.Replace(h)
				}
			}
		}
	}
	s = html.UnescapeString(s)
	for key, element := range rep {
		s = strings.ReplaceAll(s, key, element)
	}
	return s
}

func massageJoomlaLinks(s string, articles map[int]Content) string {
	r := strings.NewReader(s)
	z := html.NewTokenizer(r)

	var joomlaLink = regexp.MustCompile(`(^index.php)|(component\/content)`)
	var externalLink = regexp.MustCompile(`^http`)

	rep := make(map[string]string)

out:
	for {
		tt := z.Next()
		switch {
		case tt == html.ErrorToken:
			// end of document
			break out
		case tt == html.StartTagToken:
			t := z.Token()
			if t.Data == "a" {
				h := href(t)
				if externalLink.MatchString(h) {
					break
				}
				if joomlaLink.MatchString(h) {
					aid := articleId(h)
					newpath := articles[aid].FullPath()
					newpath = strings.ReplaceAll(newpath, ".md", "")
					rep[h] = "/" + newpath
				}
			}
		}
	}
	s = html.UnescapeString(s)
	for key, element := range rep {
		s = strings.ReplaceAll(s, key, element)
	}
	return s
}

func changeMember(s string, cat string) string {
	comp := `
import Members from "gatsby-theme-ldap/src/components/Members"

<Members group="`

	comp += cat
	comp += "\" />"
	comp += "\n"

	return strings.ReplaceAll(s, "{loadposition membres}", comp)
}

// massage cleans up a bit the html string to
// remove inline style, empty paragraph, usage
// of <br/> etc...
func massage(s string, cat string, articles map[int]Content, verbose bool) string {
	s = stringReplace(s, "<p>&nbsp;</p>", "")
	s = removeStyle(s)
	s = stringReplace(s, "<br/>", "")
	s = changeAccents(s)

	s = massageJoomlaLinks(s, articles)
	s = massageDocumentLinks(s)

	if verbose {
		fmt.Printf(s)
		panic("this is the end")
	}

	file, err := ioutil.TempFile("", "prefix")
	if err != nil {
		log.Fatal(err)
	}
	defer os.Remove(file.Name())
	cmd := exec.Command("npx", "html2md")
	cmd.Stdin = strings.NewReader(s)
	var out bytes.Buffer
	cmd.Stdout = &out
	err = cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
	s = massageImages(out.String())
	s = changeMember(s, cat)

	return s
}

func (w Content) FullPath() string {
	return filepath.Join(w.DirName(), w.FileName())
}

func (w Content) DirName() string {
	dir, _ := filepath.Split(w.Category.Path)
	s := filepath.Join(dir, w.Category.Title)
	s = strings.TrimSpace(s)
	s = strings.Replace(s, " ", "-", -1)
	return strings.ToLower(s)
}
func (w Content) FileName() string {
	return fmt.Sprintf("%s.md", w.Alias)
}

func (w Content) Write(out io.Writer, articles map[int]Content) {
	fmt.Println("-----", w.FullPath())
	fmt.Fprintln(out, "---")
	title := stringReplace(w.Title, `\`, `\\`) // for mathjax syntax
	title = stringReplace(title, `"`, `\"`)
	fmt.Fprintf(out, "title: \"%s\"\n", title)
	fmt.Fprintf(out, "date: \"%s\"\n", w.Created)
	fmt.Fprintf(out, "lastmod: \"%s\"\n", w.Modified)
	fmt.Fprintf(out, "path: \"%s\"\n", w.FullPath())
	fmt.Fprintf(out, "joomlaid: %d\n", w.ID)
	fmt.Fprintf(out, "category: \"%s\"\n", w.Category.Title)
	fmt.Fprintf(out, "asides: [\"%s.+menu+\"]\n", w.Category.Title)

	base, _ := filepath.Split(w.DirName())
	if len(base) > 0 {
		fmt.Fprintf(out, "layout: \"%s\"\n", filepath.Clean(base))
	}

	fmt.Fprintln(out, "---")

	verbose := false
	//verbose = w.FullPath() == "enseignement/masters.md"

	outstring := massage(w.Introtext, w.Category.Title, articles, verbose)
	fmt.Fprintf(out, outstring)

}

func Contents(db *sql.DB, where string,
	categories map[int]Category) ([]*Content, error) {
	var err error

	sqlstr := `SELECT id, title, alias, introtext, state, ` +
		` catid, created, modified,language FROM jlabo.wsub_content `

	qs := sqlstr + "WHERE " + where + " and state=1 and language<>'en-GB'"

	// run query
	q, err := db.Query(qs)
	if err != nil {
		return nil, err
	}
	defer q.Close()

	// load results
	res := []*Content{}
	for q.Next() {
		wc := Content{}

		// scan
		err = q.Scan(&wc.ID, &wc.Title,
			&wc.Alias, &wc.Introtext, &wc.State, &wc.Catid,
			&wc.Created, &wc.Modified, &wc.Language)
		if err != nil {
			return nil, err
		}

		wc.Category = categories[int(wc.Catid)]
		res = append(res, &wc)
	}

	return res, nil
}
