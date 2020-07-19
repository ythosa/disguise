package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"golang.org/x/net/html"
)

type Element interface {}

type FileHref struct {
	href string
}

type DirHref struct {
	href string
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

	var fileList []string
	visitNode := func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			t := CheckLink(n, extension)
			switch v := t.(type) {
			case DirHref:
				fmt.Println("Dir: ", v.href)
			case FileHref:
				fmt.Println("File: ", v.href)
			}
		}
	}

	ForEachNode(doc, visitNode, nil)
	return fileList, nil
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

	for _, a := range n.Attr {
		switch a.Key {
		case "href":
			href = a.Val
			fname := strings.Split(href, "/")[len(strings.Split(href, "/")) - 1]

			if len(strings.Split(fname, extension)) > 1 {
				isTrackedFile = true
			} else if len(strings.Split(href, "tree")) > 1 {
				isDir = true
			}

		case "class":
			if a.Val == "js-navigation-open link-gray-dark" {
				hasRightStyles = true
			}
		}
	}

	if hasRightStyles {
		if isTrackedFile {
			return FileHref{href}
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

	_, err := Extract(url, extension)
	if err != nil {
		fmt.Println(err)
	}

	//fmt.Println(res)
}
