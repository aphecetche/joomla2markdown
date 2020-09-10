package wsub

import (
	"bytes"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"strings"

	"golang.org/x/net/html"
)

func cleanupHTML(s string) string {
	s = stringReplace(s, "<p>&nbsp;</p>", "")
	s = removeStyle(s)
	s = stringReplace(s, "<br/>", "")
	s = changeAccents(s)
	s = stringReplace(s, "{emailcloak=off}", "")
	s = stringReplace(s, "{loadposition contact}", "")
	return s
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

// changeAccent replaces &acute; and so on by their corresponding characters
func changeAccents(s string) string {
	s = stringReplace(s, "&eacute;", "é")
	s = stringReplace(s, "&agrave;", "à")
	s = stringReplace(s, "&egrave;", "è")
	s = stringReplace(s, "&#39;", "'")
	s = stringReplace(s, "&#34;", "\"")
	return s
}

func html2mdnpx(s string) string {
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
	return out.String()
}
