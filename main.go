package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"regexp"
	"strings"

	"golang.org/x/net/html"
)

type Element interface {}

type FileHref struct {
	href string
	fname string
}

type DirHref struct {
	href string
}

// BreadthFirst ...
func BreadthFirst(f func(item, extension string) []string, worklist []string, extension string) {
	seen := make(map[string]bool)
	for len(worklist) > 0 {
		items := worklist
		worklist = nil
		for _, item := range items {
			if !seen[item] {
				seen[item] = true
				worklist = append(worklist, f(item, extension)...)
			}
		}
	}
}

func Crawl(url, extension string) []string {
	list, err := Extract(url, extension)
	if err != nil {
		log.Print(err)
	}

	return list
}

func Extract(url, extension string) ([]string, error) {
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("getting %s by HTML: %v", url, err)
	}

	doc, err := html.Parse(res.Body)
	if err != nil {
		return nil, fmt.Errorf("analise %s by HTML: %v", url, err)
	}

	var dirs []string
	visitNode := func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			t := CheckLink(n, extension)
			switch v := t.(type) {
			case DirHref:
				dirs = append(dirs, v.href)
			case FileHref:
				v.PrintMarkDown()
			}
		}
	}

	ForEachNode(doc, visitNode, nil)
	return dirs, nil
}

func (f *FileHref) PrintMarkDown() {
	fmt.Printf("- [ ] (%s)[%s] \n", f.fname, f.href)
}

// ForEachNode ...
func ForEachNode(n *html.Node, pre, post func(n *html.Node)) {
	if pre != nil {
		pre(n)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		ForEachNode(c, pre, post)
	}
	if post != nil {
		post(n)
	}
}


func CheckLink(n *html.Node, extension string) Element {
	var hasRightStyles = false
	var isTrackedFile = false
	var isDir = false
	var href string
	var fname string

	matchFile := regexp.MustCompile(fmt.Sprintf(".*%s$", extension))
	matchDir := regexp.MustCompile(`.+/tree/.+`)

	for _, a := range n.Attr {
		switch a.Key {
		case "href":
			href = "https://github.com" + a.Val
			fname = strings.Split(href, "/")[len(strings.Split(href, "/")) - 1]

			if matchDir.Match([]byte(href)) {
				isDir = true
			} else if matchFile.Match([]byte(fname)) {
				isTrackedFile = true
			}
		case "class":
			if a.Val == "js-navigation-open link-gray-dark" {
				hasRightStyles = true
			}
		}
	}

	if hasRightStyles {
		if isTrackedFile {
			return FileHref{href, fname[:len(fname) - len(extension)]}
		}

		if isDir {
			return DirHref{href}
		}
	}

	return nil
}

func main() {
	url := os.Args[1]
	extension := os.Args[2]

	BreadthFirst(Crawl, []string{url}, extension)

	//fmt.Println(res)
}
