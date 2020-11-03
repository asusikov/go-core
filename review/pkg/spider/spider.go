// Package spider реализует сканер содержимого веб-сайтов.
// Пакет позволяет получить список ссылок и заголовков страниц внутри веб-сайта по его URL.
package spider

import (
	"net/http"
	"strings"

	"golang.org/x/net/html"
	"goondex/pkg/goondex"
)

type Spider struct{}

func New() *Spider {
	return &Spider{}
}

// Scan осуществляет рекурсивный обход ссылок сайта, указанного в URL,
// с учётом глубины перехода по ссылкам, переданной в depth.
func (s *Spider) Scan(url string, depth int) (pages []goondex.Page, err error) {
	data := make(map[string]string)

	err = parse(url, url, depth, data)
	if err != nil {
		return pages, err
	}

	pages = []goondex.Page{}
	index := 0
	for url, title := range data {
		page := goondex.Page{
			Id:    index,
			Title: title,
			Url:   url,
		}
		pages = append(pages, page)
		index++
	}
	return pages, nil
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
