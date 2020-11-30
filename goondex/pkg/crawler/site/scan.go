// Package site реализует сканер содержимого веб-сайтов.
// Пакет позволяет получить список ссылок и заголовков страниц внутри веб-сайта по его URL.
package site

import (
	"math/rand"
	"net/http"
	"strings"
	"time"

	"goondex/webpages"

	"golang.org/x/net/html"
)

// Scan осуществляет рекурсивный обход ссылок сайта, указанного в URL,
// с учётом глубины перехода по ссылкам, переданной в depth.
func Scan(url string, depth int) (pages []webpages.Page, err error) {
	data := make(map[string]string)

	err = parse(url, url, depth, data)
	if err != nil {
		return pages, err
	}

	rand.Seed(time.Now().UnixNano())
	pages = []webpages.Page{}
	for url, title := range data {
		page := webpages.New(title, url)
		pages = insertRandom(pages, *page)
	}
	return pages, nil
}

// добавляет страницу в произвольное позицию в слайсе
func insertRandom(pages []webpages.Page, page webpages.Page) []webpages.Page {
	pages = append(pages, webpages.Page{})
	pos := rand.Intn(len(pages))
	copy(pages[pos+1:], pages[pos:])
	pages[pos] = page
	return pages
}

// parse рекурсивно обходит ссылки на странице, переданной в url.
// Глубина рекурсии задаётся в depth.
// Каждая найденная ссылка записывается в ассоциативный массив
// data вместе с названием страницы.
func parse(url, baseurl string, depth int, data map[string]string) error {
	if depth == 0 {
		return nil
	}

	response, err := http.Get(url)
	if err != nil {
		return err
	}
	page, err := html.Parse(response.Body)
	if err != nil {
		return err
	}

	data[url] = pageTitle(page)

	links := pageLinks(nil, page)
	for _, link := range links {
		// ссылка уже отсканирована
		if data[link] != "" {
			continue
		}

		if isInternalLink(link, baseurl) {
			nextLink := fullUrl(link, baseurl)
			err := parse(nextLink, baseurl, depth-1, data)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func isInternalLink(link string, baseurl string) bool {
	return strings.HasPrefix(link, baseurl) || (strings.HasPrefix(link, "/") && len(link) > 1)
}

func fullUrl(link string, baseurl string) string {
	// if strings.HasPrefix(link, "/") && len(link) > 1 {
	// next := baseurl + link[1:]

	return baseurl + strings.ReplaceAll(link, baseurl, "")
}

// pageTitle осуществляет рекурсивный обход HTML-страницы и возвращает значение элемента <title>.
func pageTitle(n *html.Node) string {
	var title string
	if n.Type == html.ElementNode && n.Data == "title" && n.FirstChild != nil {
		return n.FirstChild.Data
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		title = pageTitle(c)
		if title != "" {
			break
		}
	}
	return title
}

// pageLinks рекурсивно сканирует узлы HTML-страницы и возвращает все найденные ссылки без дубликатов.
func pageLinks(links []string, n *html.Node) []string {
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
			if a.Key == "href" {
				if !sliceContains(links, a.Val) {
					links = append(links, a.Val)
				}
			}
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		links = pageLinks(links, c)
	}
	return links
}

// sliceContains возвращает true если массив содержит переданное значение
func sliceContains(slice []string, value string) bool {
	for _, v := range slice {
		if v == value {
			return true
		}
	}
	return false
}
