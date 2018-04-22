package wsub

import (
	"fmt"
	"net/url"
	"regexp"
	"strconv"
	"strings"

	"golang.org/x/net/html"
)

func getLinks(c Content) []string {

	links := []string{}

	r := strings.NewReader(c.Introtext)
	doc, err := html.Parse(r)
	if err != nil {
		return nil
	}
	var f func(*html.Node)

	var externalLink = regexp.MustCompile(`^(http|mailto|www)`)

	f = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data[0] == 'a' {
			for _, a := range n.Attr {
				if strings.Contains(a.Key, "href") && !externalLink.MatchString(a.Val) {
					links = append(links, a.Val)
				}
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}
	f(doc)
	return links
}

func mapLinks(links []string, articles map[int]Content) map[string]string {

	olinks := map[string]string{}

	var arlink = regexp.MustCompile(`^index.php`)

	for _, l := range links {
		if arlink.MatchString(l) {
			u, err := url.Parse(l)
			if err != nil {
				panic(err)
			}
			q := u.Query()
			i := q["id"]
			if len(i) == 0 {
				fmt.Println("ERROR - could not decode link", l)
				continue
			}
			said := strings.Split(i[0], ":")
			aid, err := strconv.Atoi(said[0])
			if err != nil {
				panic(err)
			}
			olinks[l] = articles[aid].FullPath()
		}
	}
	return olinks
}

func menu2links(links []string, menus []*Menu) []string {
	olinks := []string{}
	var melink = regexp.MustCompile(`^fr/`)
	for _, l := range links {
		if melink.MatchString(l) {
			for _, m := range menus {
				if m.Path == l[3:] {
					l = m.Link
				}
			}
		}
		olinks = append(olinks, l)
	}
	return olinks
}

// ListLinks makes a map of all internal links in articles
// from the joomla way to hugo way
func ListLinks(articles map[int]Content, menus []*Menu) map[string]string {
	links := []string{}
	for _, a := range articles {
		links = append(links, getLinks(a)...)
	}
	links = menu2links(links, menus)
	return mapLinks(links, articles)
}

func UpdateLinks(c *Content, links map[string]string) {
	r := strings.NewReader(c.Introtext)
	doc, err := html.Parse(r)
	if err != nil {
		panic(err)
	}
	var f func(*html.Node)
	f = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data[0] == 'a' {
			for i := 0; i < len(n.Attr); i++ {
				if strings.Contains(n.Attr[i].Key, "href") {
					l, ok := links[n.Attr[i].Val]
					if ok {
						n.Attr[i].Val = l
					}
				}
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}
	f(doc)
	c.Introtext = render(doc)
}
