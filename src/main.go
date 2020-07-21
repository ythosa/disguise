package main

import (
	"fmt"
	"net/http"
	"os"
	"regexp"
	"strings"

	"golang.org/x/net/html"
)

type Element interface {}

type FileHref struct {
	href string
	name string
	dir string
}

type DirHref struct {
	href string
}

func Extract(url, extension string, files *map[string]chan FileHref) error {
	res, err := http.Get(url)
	if err != nil {
		return err
	}

	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("getting %s by HTML: %v", url, err)
	}

	doc, err := html.Parse(res.Body)
	if err != nil {
		return fmt.Errorf("analise %s by HTML: %v", url, err)
	}

	visitNode := func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			t := CheckLink(n, extension)
			switch v := t.(type) {
			case DirHref:
				//dirs = append(dirs, v.href)
				//fmt.Println(v.href)
				go Extract(v.href, extension, files)
			case FileHref:
				//v.PrintMarkDown()
				(*files)[v.dir] <- v
			}
		}
	}

	ForEachNode(doc, visitNode)
	return nil
}

// ForEachNode ...
func ForEachNode(n *html.Node, pre func(n *html.Node)) {
	if pre != nil {
		pre(n)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		ForEachNode(c, pre)
	}
}

func (f *FileHref) PrintMarkDown() {
	fmt.Printf("- [ ] [%s](%s) \n", f.name, f.href)
}

func CheckLink(n *html.Node, extension string) Element {
	var hasRightStyles = false
	var isTrackedFile = false
	var isDir = false

	var href string
	var fname string
	var dirname string

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
				dirname = strings.Join(strings.Split(href, "/")[4:], "")
			}
		case "class":
			if a.Val == "js-navigation-open link-gray-dark" {
				hasRightStyles = true
			}
		}
	}

	if hasRightStyles {
		if isTrackedFile {
			return FileHref {
				href: href,
				name: fname[:len(fname) - len(extension)],
				dir: dirname,
			}
		}

		if isDir {
			return DirHref {
				href: href,
			}
		}
	}

	return nil
}

//func CrawlLinks(url string, files map[string]chan string) {
//	Extract(url, files)
//}

func main() {
	url := os.Args[1]
	extension := os.Args[2]

	doccedfiles := make(map[string]chan FileHref)

	go Extract(url, extension, &doccedfiles)

	for {
		for _, d := range doccedfiles {
			select {
			case f := <- d:
				f.PrintMarkDown()
			}
		}
	}
}
