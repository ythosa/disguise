package main

import (
	"fmt"
	"io"
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
	name string
	dir string
}

type DirHref struct {
	href string
}

func (f *FileHref) GetMarkDown() string {
	return fmt.Sprintf("- [ ] [%s](%s) \n", f.name, f.href)
}

func CheckLink(n *html.Node, extension string) Element {
	var hasRightStyles = false
	var isTrackedFile = false
	var isDir = false

	var href string
	var fname string
	var dirname string

	matchFile := regexp.MustCompile(fmt.Sprintf(".+/blob/.+%s$", extension))
	matchDir := regexp.MustCompile(`.+/tree/.+`)

	for _, a := range n.Attr {
		switch a.Key {
		case "href":
			href = "https://github.com" + a.Val
			fname = strings.Split(href, "/")[len(strings.Split(href, "/")) - 1]

			if matchDir.Match([]byte(href)) {
				isDir = true
			} else if matchFile.Match([]byte(href)) {
				isTrackedFile = true
				pathArray := strings.Split(href, "/")
				dirname = strings.Join(pathArray[7:len(pathArray)-1], "/")
			}
		case "class":
			if a.Val == "js-navigation-open link-gray-dark" {
				hasRightStyles = true
			}
		}
	}

	if hasRightStyles && isTrackedFile {
		return FileHref{
			href: href,
			name: fname[:len(fname)-len(extension)],
			dir:  dirname,
		}
	}

	if hasRightStyles && isDir {
		return DirHref {
			href: href,
		}
	}

	return nil
}

func Extract(url, extension string) []Element {
	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		log.Fatalf("getting %s by HTML: %v", url, err)
	}

	doc, err := html.Parse(res.Body)
	if err != nil {
		log.Fatalf("analise %s by HTML: %v", url, err)
	}

	files := make([]Element, 0)
	visitNode := func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			files = append(files, CheckLink(n, extension))
		}
	}

	ForEachNode(doc, visitNode)
	return files
}

func ForEachNode(n *html.Node, f func(n *html.Node)) {
	f(n)
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		ForEachNode(c, f)
	}
}

func GroupByDir(files []FileHref) map[string][]FileHref {
	grouped := make(map[string][]FileHref)
	for _, f := range files {
		grouped[f.dir] = append(grouped[f.dir], f)
	}

	return grouped
}

func Crawl(url, extension string) []FileHref {
	worklist := make(chan []Element)
	results := make([]FileHref, 0)

	// Start with cmd arguments
	go func() {
		worklist <- Extract(url, extension)
	}()

	for n := 1; n > 0; n-- {
		list := <-worklist
		for _, f := range list {
			switch v := f.(type) {
			case DirHref:
				n++
				go func() {
					worklist <- Extract(v.href, extension)
				}()
			case FileHref:
				results = append(results, v)
			}
		}
	}

	return results
}

func PrintResults(out io.Writer, results []FileHref) {
	for dir, files := range GroupByDir(results) {
		_, err := fmt.Fprintf(out, "* ### %s\n", dir)
		if err != nil {
			log.Fatal(err)
		}
		for _, f := range files {
			_, err := fmt.Fprint(out, f.GetMarkDown())
			if err != nil {
				log.Fatal(err)
			}
		}
		_, err = fmt.Fprintf(out, "\n")
		if err != nil {
			log.Fatal(err)
		}
	}
}

func main() {
	url := os.Args[1]
	extension := os.Args[2]

	md := Crawl(url, extension)

	fname := strings.Split(url, "/")[len(strings.Split(url, "/")) - 1]
	f, err := os.Create(fmt.Sprintf("%s.txt", fname))
	if err != nil {
		log.Fatal(err)
	}

	PrintResults(f, md)
}
